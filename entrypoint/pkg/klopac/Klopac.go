package klopac

import (
	"encoding/json"
	"entrypoint/pkg/helper"
	"entrypoint/pkg/logger"
	"entrypoint/pkg/websocket"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/imdario/mergo"
	"go.uber.org/zap"
)

// It might be considered as main function. It will execute some sort of code blocks depending to whether we are going to access klopac via a websocket or from command-line
func Run() {
	logger.InitializeLogger(helper.GetParam[string]("logLevel"), helper.GetParam[string]("logFile"))
	log := logger.GetLogger()
	webSocketEnabled := helper.GetParam[bool]("websocket")
	if webSocketEnabled == true {
		log.Info("[WEBSOCKET CONNECTION - START]")
		uri := helper.GetParam[string]("uri")
		username := helper.GetParam[string]("username")
		password := helper.GetParam[string]("password")
		websocket.Enable(uri, username, password)
		log.Info("[WEBSOCKET CONNECTION - END]")
	} else {
		bundleFile := helper.GetParam[string]("bundleFile")
		bundleFileExists := false
		if _, err := os.Stat(bundleFile); !errors.Is(err, os.ErrNotExist) {
			bundleFileExists = true
			log.Info("[BUNDLE FILE UNTAR - START]")
			err := helper.Untar(bundleFile, helper.GetParam[string]("dataPath"))
			log.Debug(bundleFile)
			if err != nil {
				log.Panic("Error while untarring bundle file, please check whether you have correct named bundlefile ")
			}
			log.Debug("SUCCESS")
			log.Info("[BUNDLE FILE UNTAR - END]")
		} else {
			valuesModel := helper.ReadFile(helper.GetParam[string]("valuesFile"))
			varsPath := helper.GetParam[string]("varsPath")
			manifestsPath := helper.GetParam[string]("manifestsPath")

			for rootKey, rootValue := range valuesModel {
				//Merge Intersection Objects - Default and Domain Objects
				mergeVariables(varsPath, rootKey, rootValue)
				mergeVariables(manifestsPath, rootKey, rootValue)

				//Create New Objects
				createNewObjects(rootKey, varsPath, manifestsPath, rootValue, valuesModel)
			}
		}

		provision := helper.GetParam[bool]("provision")
		validate := helper.GetParam[bool]("validate")
		logLevel := helper.GetParam[string]("logLevel")
		healthCheck := helper.GetParam[bool]("healthcheck")
		log.Info("[KLOPAC FLOW - START]",
			zap.Bool("provision", provision),
			zap.Bool("validate", validate),
			zap.Bool("healthcheck", healthCheck))
		helper.GetFlowService().Run(provision, validate, healthCheck, logLevel, bundleFileExists)
		log.Info("[KLOPAC FLOW - END]")
	}
}

func mergeVariables(varsPath, rootKey string, rootValue interface{}) {
	temp := make(map[string]interface{})
	temp[rootKey] = rootValue
	err := helper.UpdateValuesFile(temp, varsPath)
	if err != nil {
		log.Panic(fmt.Sprintf("Error while patching default values for %v", varsPath), zap.Error(err))
	}
}

func createNewObjects(rootKey, varsPath, manifestsPath string, rootValue interface{}, valuesModel map[string]interface{}) {
	//Create New Domain Defaults, Domain Objects and Executors
	if rootKey == "app" || rootKey == "int" {

		for innerKey, innerValue := range rootValue.(map[string]interface{}) {
			if _, ok := innerValue.(map[string]interface{}); ok {
				createDefaultFile(varsPath, rootKey, innerKey, innerValue)
				createRunnerFile(innerKey, manifestsPath, innerValue, valuesModel)
			}
		}

	}
}

func createDefaultFile(varsPath, rootKey, innerKey string, innerValue interface{}) {
	//Check Default File Exists
	newDefaultObjectPath := fmt.Sprintf("%v/%v-%v.yaml", varsPath+"/defaults", rootKey, innerKey)
	if _, err := os.Stat(newDefaultObjectPath); err != nil {
		version := innerValue.(map[string]interface{})["version"]
		//Create New Files
		if version != nil {
			//Create Default File with Version Variable
			jsonString := fmt.Sprintf(`{"%v":{"%v": {"version": "%v"}}}`, rootKey, innerKey, innerValue.(map[string]interface{})["version"])
			jsonMap, _ := convertStringToJsonMap(jsonString)
			helper.WriteFile(newDefaultObjectPath, jsonMap)

			//In App or Int Layer, Create New Domain Object
			deleteKeyFromDomainObject(innerValue.(map[string]interface{}), []string{"version"})
			jsonMap = createDomainObject(innerValue, fmt.Sprintf(`{"%v":{}}`, rootKey), rootKey, innerKey)
			mergeFiles(varsPath, rootKey, jsonMap)
		}
	}
}

func deleteKeyFromDomainObject(innerValue map[string]interface{}, exclude []string) {
	for _, i := range exclude {
		delete(innerValue, i)
	}
}

func createDomainObject(innerValue interface{}, jsonString, rootKey, innerKey string) map[string]interface{} {
	jsonMap, _ := convertStringToJsonMap(jsonString)
	jsonMap[rootKey].(map[string]interface{})[innerKey] = innerValue
	return jsonMap
}

func mergeFiles(varsPath, rootKey string, jsonMap map[string]interface{}) {
	filePath := fmt.Sprintf("%v/%v", varsPath, getFileNameByRootKey(rootKey))
	defaultModel := helper.ReadFile(filePath)
	mergo.Merge(&defaultModel, jsonMap)
	helper.WriteFile(filePath, defaultModel)
}

func getFileNameByRootKey(rootKey string) string {
	var v string
	if rootKey == "app" {
		v = "applications.yaml"
	} else {
		v = "integrations.yaml"
	}
	return v
}

func createRunnerFile(innerKey, manifestsPath string, innerValue interface{}, valuesModel map[string]interface{}) {
	//Create Runner File with Domain Object
	if runner, ok := innerValue.(map[string]interface{})["runner"]; ok {
		if manifestType, ok := runner.(map[string]interface{})["type"]; ok {
			manifestType := manifestType.(string)
			if platformRunner, ok := valuesModel["platform"].(map[string]interface{})[manifestType]; ok {
				if newPlatformObject, ok := platformRunner.(map[string]interface{})[innerKey]; ok {
					newManifestObjectPath := fmt.Sprintf("%v/%v/%v-%v.yaml", manifestsPath, manifestType, innerKey, newPlatformObject.(map[string]interface{})["version"])
					jsonMap, _ := convertStringToJsonMap(fmt.Sprintf(`{"platform":{"%v": {"%v": {}}}}`, manifestType, innerKey))
					jsonMap["platform"].(map[string]interface{})[manifestType].(map[string]interface{})[innerKey] = newPlatformObject
					helper.WriteFile(newManifestObjectPath, jsonMap)
				}
			}
		}
	}
}

func convertStringToJsonMap(jsonString string) (map[string]interface{}, error) {
	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &jsonMap)
	return jsonMap, err
}
