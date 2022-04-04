package klopac

import (
	"entrypoint/pkg/helper"
	"entrypoint/pkg/websocket"
	"log"
)

// It might be considered as main function. It will execute some sort of code blocks depending to whether we are going to access klopac via a websocket or from command-line
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
			err := helper.UpdateValuesFile(valuesModel, helper.GetParam[string]("varsPath"))
			if err != nil {
				log.Println(err)
				return
			}
		}
		provision := helper.GetParam[bool]("provision")
		validate := helper.GetParam[bool]("validate")
		logLevel := helper.GetParam[string]("loglevel")
		healthCheck := helper.GetParam[bool]("healthcheck")

		helper.GetFlowService().Run(provision, validate, healthCheck, logLevel)
	}
}
