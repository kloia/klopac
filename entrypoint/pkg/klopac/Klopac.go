package klopac

import (
	"entrypoint/pkg/command"
	"entrypoint/pkg/flag"
	"entrypoint/pkg/flow"
	"entrypoint/pkg/model"
	"entrypoint/pkg/option"
	"entrypoint/pkg/shell"
	"entrypoint/pkg/websocket"
	"github.com/imdario/mergo"
	"os"
	"path/filepath"
)

var (
	OptionService = option.NewOptionService(flag.NewFlagService())
	FlowService   = flow.NewFlowService(shell.NewShellService(command.NewCommandService()))
)

func GetParam[V comparable](param string) V {
	return *OptionService.Params[param].(*V)
}

func Run() {
	webSocketEnabled := GetParam[bool]("websocket")

	if webSocketEnabled == true {
		uri := GetParam[string]("uri")
		username := GetParam[string]("username")
		password := GetParam[string]("password")
		websocket.Enable(uri, username, password)
	} else {
		bundleFile := GetParam[string]("bundleFile")
		if bundleFile != "bundle.tar.gz" {
			//TODO: UNTAR BUNDLE FILE AND OVERRIDE VARS FOLDER
		} else {
			valuesModel := model.ReadFile(GetParam[string]("valuesFile"))
			UpdateValuesFile(valuesModel, GetParam[string]("varsPath"))
		}
		provision := GetParam[bool]("provision")
		validate := GetParam[bool]("validate")
		logLevel := GetParam[string]("loglevel")
		healthCheck := GetParam[bool]("healthcheck")

		FlowService.Run(provision, validate, healthCheck, logLevel)
	}
}

func UpdateValuesFile(valuesModel map[string]interface{}, varsPath string) {
	filepath.Walk(varsPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && filepath.Ext(path) == ".yaml" {
				defaultModel := model.ReadFile(path)
				newMap := model.GetDifference(valuesModel, defaultModel)
				mergo.Merge(&defaultModel, newMap, mergo.WithOverride)
				model.WriteFile(path, defaultModel)
			}
			return nil
		})
}
