package helper

import (
	"archive/tar"
	"bytes"
	"entrypoint/pkg/command"
	"entrypoint/pkg/flag"
	"entrypoint/pkg/flow"
	"entrypoint/pkg/option"
	"entrypoint/pkg/shell"
	"github.com/imdario/mergo"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"

	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Reads content of the yaml file and returns it
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

// Writes content to a yaml file

func WriteFile(filename string, data map[string]interface{}) error {
	var b bytes.Buffer
	yamlEncoder := yaml.NewEncoder(&b)
	yamlEncoder.SetIndent(2)

	yamlEncoder.Encode(&data)

	err := ioutil.WriteFile(filename, b.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}

// Basically we have two map and we compare them if there are some sort of values that should be changed according to its logic
func Intersection(inputMap, defaultMap map[string]interface{}) (newMap map[string]interface{}) {
	newMap = make(map[string]interface{})
	for inputKey, inputVal := range inputMap {
		if inputInnerMap, ok := inputVal.(map[string]interface{}); ok {
			defaultInnerMap, ok := defaultMap[inputKey].(map[string]interface{})
			if ok {
				newMap[inputKey] = Intersection(inputInnerMap, defaultInnerMap)
			}
		} else {
			_, ok := defaultMap[inputKey]
			if ok {
				newMap[inputKey] = inputVal
			}
		}
	}
	return newMap
}

// Basically takes a interface and varsPath(which is path of the variable files) then it starts to override or leaves unchanged depending to intersection logic
func UpdateValuesFile(valuesModel map[string]interface{}, varsPath string) error {
	return filepath.Walk(varsPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && filepath.Ext(path) == ".yaml" || filepath.Ext(path) == ".yml" {
				defaultModel := ReadFile(path)
				intersectionMap := Intersection(valuesModel, defaultModel)
				if len(intersectionMap) > 0 {
					mergo.Merge(&defaultModel, intersectionMap, mergo.WithOverride)
					WriteFile(path, defaultModel)
				}
			}
			return nil
		})
}

func Untar(tarball, target string) error {
	reader, err := os.Open(tarball)
	if err != nil {
		return err
	}
	defer reader.Close()
	tarReader := tar.NewReader(reader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		path := filepath.Join(target, header.Name)
		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(path, info.Mode()); err != nil {
				return err
			}
			continue
		}

		file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(file, tarReader)
		if err != nil {
			return err
		}
	}
	return nil
}

func InitializeLogger() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	logFile, _ := os.OpenFile("deneme.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)
	defaultLogLevel, err := zapcore.ParseLevel(GetParam[string]("loglevel"))
	if err != nil {
		defaultLogLevel = zapcore.InfoLevel
	}
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)
	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

var (
	optionService = option.NewOptionService(flag.NewFlagService())
	flowService   = flow.NewFlowService(shell.NewShellService(command.NewCommandService()))
	logger        *zap.Logger
)

func GetOptionService() *option.OptionService {
	return optionService
}

func GetFlowService() flow.Flow {
	return flowService
}

func GetLogger() *zap.Logger {
	return logger
}

func GetParam[V comparable](param string) V {
	return *GetOptionService().Params[param].(*V)
}
