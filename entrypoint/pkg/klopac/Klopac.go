package klopac

import (
	"entrypoint/pkg/command"
	"entrypoint/pkg/flag"
	"entrypoint/pkg/flow"
	"entrypoint/pkg/model"
	"entrypoint/pkg/option"
	"entrypoint/pkg/shell"
	"entrypoint/pkg/websocket"
	"fmt"
	"io/ioutil"
	"log"
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
		updateValuesFile(GetParam[string]("varsPath"))
		/*provision := GetParam[bool]("provision")
		validate := GetParam[bool]("validate")
		logLevel := GetParam[string]("loglevel")
		healthCheck := GetParam[bool]("healthcheck")
		FlowService.Run(provision, validate, healthCheck, logLevel)*/
	}
}

func updateValuesFile(varsPath string) {
	valuesModel := model.CreateByFileName(GetParam[string]("valuesFile"))
	files, err := ioutil.ReadDir(varsPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println()
		mainModel := model.CreateByFileName(fmt.Sprintf("%v/%v", varsPath, f.Name()))
		mainModel.Merge(valuesModel)
		mainModel.WriteFile(fmt.Sprintf("%v/%v", varsPath, f.Name()))
	}
}
