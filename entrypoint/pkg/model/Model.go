package model

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
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
