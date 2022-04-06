package helper

import (
	"archive/tar"
	"bytes"
	"entrypoint/pkg/command"
	"entrypoint/pkg/flag"
	"entrypoint/pkg/flow"
	"entrypoint/pkg/option"
	"entrypoint/pkg/shell"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"

	"github.com/imdario/mergo"
	"gopkg.in/yaml.v3"
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
func Intersection(first, second map[string]interface{}) map[string]interface{} {
	newMap := make(map[string]interface{})
	IntersectionHelper(first, second, newMap)
	return newMap
}

// Code blocks of the intersection logic. It compares two different map and runs the provided conditions
func IntersectionHelper(inputMap, defaultMap, newMap interface{}) {
	for inputKey, inputVal := range inputMap.(map[string]interface{}) {
		if defaultVal, ok := defaultMap.(map[string]interface{})[inputKey]; ok {
			if reflect.TypeOf(defaultVal).Kind() == reflect.Bool || reflect.TypeOf(defaultVal).Kind() == reflect.String {
				newMap.(map[string]interface{})[inputKey] = inputVal
			} else if reflect.TypeOf(defaultVal).Kind() == reflect.Slice {
				for _, defaultValElem := range defaultVal.([]interface{}) {
					if reflect.TypeOf(defaultValElem).Kind() == reflect.Bool || reflect.TypeOf(defaultValElem).Kind() == reflect.String {
						if ok, i := contains(inputVal.([]interface{}), defaultValElem); ok {
							if newMap.(map[string]interface{})[inputKey] == nil {
								newMap.(map[string]interface{})[inputKey] = make([]interface{}, 0)
							}
							newMap.(map[string]interface{})[inputKey] = append(newMap.(map[string]interface{})[inputKey].([]interface{}), inputVal.([]interface{})[i])
						}
					} else {
						if newMap.(map[string]interface{})[inputKey] == nil {
							newMap.(map[string]interface{})[inputKey] = make([]interface{}, 0)
						}
						for _, inputValElem := range inputVal.([]interface{}) {
							tempMap := make(map[string]interface{})
							IntersectionHelper(inputValElem, defaultValElem, tempMap)
							if len(tempMap) > 0 {
								newMap.(map[string]interface{})[inputKey] = append(newMap.(map[string]interface{})[inputKey].([]interface{}), tempMap)
							}
						}
					}
				}
			} else {
				newMap.(map[string]interface{})[inputKey] = make(map[string]interface{})
				IntersectionHelper(inputVal, defaultVal, newMap.(map[string]interface{})[inputKey])
			}
		}
	}
}

// To check whether its content have the searched struct or not
func contains(s []interface{}, e interface{}) (bool, int) {
	for i, a := range s {
		if a == e {
			return true, i
		}
	}
	return false, -1
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
