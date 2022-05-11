package helper

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"entrypoint/pkg/command"
	"entrypoint/pkg/flag"
	"entrypoint/pkg/flow"
	"entrypoint/pkg/logger"
	"entrypoint/pkg/option"
	"entrypoint/pkg/shell"
	"fmt"

	"github.com/imdario/mergo"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Reads content of the yaml file and returns it
func ReadFile(filename string) map[string]interface{} {
	log := logger.GetLogger()
	log.Debug("START: READ FILE", zap.String("filename", filename))
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Debug("Could not open",
			zap.String("filename", filename))
	}

	m := make(map[string]interface{})
	err = yaml.Unmarshal(yamlFile, m)
	if err != nil {
		log.Debug("Could not decode map",
			zap.Any("map", m))
	}
	log.Debug("END: READ FILE", zap.String("filename", filename))
	return m
}

// Writes content to a yaml file

func WriteFile(filename string, data map[string]interface{}) error {
	log := logger.GetLogger()
	log.Debug("START: WRITE FILE", zap.String("filename", filename))
	var b bytes.Buffer
	yamlEncoder := yaml.NewEncoder(&b)
	yamlEncoder.SetIndent(2)

	yamlEncoder.Encode(&data)
	err := ioutil.WriteFile(filename, b.Bytes(), 0644)
	if err != nil {
		log.Error("Error while writing file ", zap.Error(err), zap.String("filename", filename))
	}
	log.Debug("END: WRITE FILE", zap.String("filename", filename))
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
	log := logger.GetLogger()
	return filepath.Walk(varsPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && filepath.Ext(path) == ".yaml" || filepath.Ext(path) == ".yml" {
				defaultModel := ReadFile(path)
				intersectionMap := Intersection(valuesModel, defaultModel)
				if len(intersectionMap) > 0 {
					log.Debug("INTERSECTION", zap.Any("map", intersectionMap))
					mergo.Merge(&defaultModel, intersectionMap, mergo.WithOverride)
					WriteFile(path, defaultModel)
				}
			}
			return nil
		})
}

func Untar(tarball, target string) error {

	reader, err := os.Open(tarball)
	log := logger.GetLogger()
	if err != nil {
		log.Debug("Bundle file could not be opened.")
		return err
	}
	log.Debug("Bundle file exists.")
	defer reader.Close()

	gzf, err := gzip.NewReader(reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Debug("Gunzip done.")

	tarReader := tar.NewReader(gzf)

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
	log.Debug("Untar done.")
	return nil
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
