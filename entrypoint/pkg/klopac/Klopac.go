package klopac

import (
	"entrypoint/pkg/helper"
	"entrypoint/pkg/websocket"
)

func Run() {
	webSocketEnabled := helper.GetParam[bool]("websocket")

	if webSocketEnabled == true {
		uri := helper.GetParam[string]("uri")
		username := helper.GetParam[string]("username")
		password := helper.GetParam[string]("password")
		websocket.Enable(uri, username, password)
	} else {
		bundleFile := helper.GetParam[string]("bundleFile")
		if bundleFile != "bundle.tar.gz" {
			//TODO: UNTAR BUNDLE FILE AND OVERRIDE VARS FOLDER
		} else {
			valuesModel := helper.ReadFile(helper.GetParam[string]("valuesFile"))
			helper.UpdateValuesFile(valuesModel, helper.GetParam[string]("varsPath"))
		}
		provision := helper.GetParam[bool]("provision")
		validate := helper.GetParam[bool]("validate")
		logLevel := helper.GetParam[string]("loglevel")
		healthCheck := helper.GetParam[bool]("healthcheck")

		helper.GetFlowService().Run(provision, validate, healthCheck, logLevel)
	}
}
