package model

import (
	"fmt"
	"github.com/imdario/mergo"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Engine struct {
	Hits int64 `yaml:"hits,omitempty"`
	Time int64 `yaml:"time,omitempty"`
}

type Model struct {
	Engine Engine `yaml:"engine,omitempty"`
}

func (c *Model) Set(filename string) *Model {

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func (c *Model) Merge(other *Model) {
	_ = mergo.Merge(c, other, mergo.WithOverride)
}

func (c *Model) WriteFile(filename string) error {
	yamlData, err := yaml.Marshal(c)

	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
	}

	err = ioutil.WriteFile(filename, yamlData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func CreateByFileName(filename string) *Model {
	var model Model
	model.Set(filename)
	return &model
}
