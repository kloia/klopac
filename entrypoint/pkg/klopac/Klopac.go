package klopac

import (
	"entrypoint/pkg/helper"
	"entrypoint/pkg/logger"
	"entrypoint/pkg/websocket"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"os"
)

// It might be considered as main function. It will execute some sort of code blocks depending to whether we are going to access klopac via a websocket or from command-line
func Run() {
	logger.InitializeLogger(helper.GetParam[string]("logLevel"), helper.GetParam[string]("logFile"))
	log := logger.GetLogger()
	webSocketEnabled := helper.GetParam[bool]("websocket")
	if webSocketEnabled == true {
		log.Info("START: WEBSOCKET ENABLING")
		uri := helper.GetParam[string]("uri")
		username := helper.GetParam[string]("username")
		password := helper.GetParam[string]("password")
		websocket.Enable(uri, username, password)
		log.Info("END: WEBSOCKET ENABLING")
	} else {
		bundleFile := helper.GetParam[string]("bundleFile")
		if _, err := os.Stat(bundleFile); !errors.Is(err, os.ErrNotExist) {
			log.Info("START: BUNDLE FILE UNTAR")
			err := helper.Untar(bundleFile, helper.GetParam[string]("dataPath"))
			if err != nil {
				log.Error("Error while untarring bundle file, please check whether you have correct named bundlefile ")
			}
			log.Info("END: BUNDLE FILE UNTAR")
		} else {
			valuesModel := helper.ReadFile(helper.GetParam[string]("valuesFile"))
			err := helper.UpdateValuesFile(valuesModel, helper.GetParam[string]("varsPath"))
			fmt.Println("hiiii")
			if err != nil {
				log.Error("Error while patching default values", zap.Error(err))
			}
		}
		fmt.Println("asd23")
		provision := helper.GetParam[bool]("provision")
		validate := helper.GetParam[bool]("validate")
		logLevel := helper.GetParam[string]("logLevel")
		healthCheck := helper.GetParam[bool]("healthcheck")
		fmt.Println("asd")
		log.Info("START: KLOPAC FLOW",
			zap.Bool("provision", provision),
			zap.Bool("validate", validate),
			zap.Bool("healthcheck", healthCheck))
		helper.GetFlowService().Run(provision, validate, healthCheck, logLevel)
		log.Info("END: KLOPAC FLOW")
	}
}
