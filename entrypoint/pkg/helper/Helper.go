package helper

import (
	"entrypoint/pkg/command"
	"entrypoint/pkg/flag"
	"entrypoint/pkg/flow"
	"entrypoint/pkg/option"
	"entrypoint/pkg/shell"
	"fmt"
	"github.com/imdario/mergo"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func ReadFile(filename string) map[string]interface{} {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	m := make(map[string]interface{})
	err = yaml.Unmarshal(yamlFile, m)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return m
}

func WriteFile(filename string, data map[string]interface{}) error {
	yamlData, err := yaml.Marshal(data)

	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
	}

	err = ioutil.WriteFile(filename, yamlData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func GetDifference(first map[string]interface{}, second map[string]interface{}) map[string]interface{} {
	newMap := make(map[string]interface{})
	for k, v := range first {
		if _, ok := second[k]; ok {
			newMap[k] = v
		}
	}
	return newMap
}

func UpdateValuesFile(valuesModel map[string]interface{}, varsPath string) {
	filepath.Walk(varsPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && filepath.Ext(path) == ".yaml" {
				defaultModel := ReadFile(path)
				newMap := GetDifference(valuesModel, defaultModel)
				mergo.Merge(&defaultModel, newMap, mergo.WithOverride)
				WriteFile(path, defaultModel)
			}
			return nil
		})
}

var (
	optionService = option.NewOptionService(flag.NewFlagService())
	flowService   = flow.NewFlowService(shell.NewShellService(command.NewCommandService()))
)

func GetOptionService() *option.OptionService {
	return optionService
}

func GetFlowService() flow.Flow {
	return flowService
}

func GetParam[V comparable](param string) V {
	return *GetOptionService().Params[param].(*V)
}
