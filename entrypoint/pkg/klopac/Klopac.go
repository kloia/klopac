package klopac

import (
	"entrypoint/pkg/helper"
	"entrypoint/pkg/logger"
	"entrypoint/pkg/websocket"
	"errors"
	"os"

	"go.uber.org/zap"
)

// It might be considered as main function. It will execute some sort of code blocks depending to whether we are going to access klopac via a websocket or from command-line
func Run() {
	logger.InitializeLogger(helper.GetParam[string]("logLevel"), helper.GetParam[string]("logFile"))
	log := logger.GetLogger()
	webSocketEnabled := helper.GetParam[bool]("websocket")
	if webSocketEnabled == true {
		log.Info("[WEBSOCKET CONNECTION - START]")
		uri := helper.GetParam[string]("uri")
		username := helper.GetParam[string]("username")
		password := helper.GetParam[string]("password")
		websocket.Enable(uri, username, password)
		log.Info("[WEBSOCKET CONNECTION - END]")
	} else {
		bundleFile := helper.GetParam[string]("bundleFile")
		if _, err := os.Stat(bundleFile); !errors.Is(err, os.ErrNotExist) {
			log.Info("[BUNDLE FILE UNTAR - START]")
			err := helper.Untar(bundleFile, helper.GetParam[string]("dataPath"))
			if err != nil {
				log.Panic("Error while untarring bundle file, please check whether you have correct named bundlefile ")
			}
			log.Info("[BUNDLE FILE UNTAR - END]")
		} else {
			valuesModel := helper.ReadFile(helper.GetParam[string]("valuesFile"))
			err := helper.UpdateValuesFile(valuesModel, helper.GetParam[string]("varsPath"))
			if err != nil {
				log.Panic("Error while patching default values", zap.Error(err))
			}
		}
		provision := helper.GetParam[bool]("provision")
		validate := helper.GetParam[bool]("validate")
		logLevel := helper.GetParam[string]("logLevel")
		healthCheck := helper.GetParam[bool]("healthcheck")
		log.Info("[KLOPAC FLOW - START]",
			zap.Bool("provision", provision),
			zap.Bool("validate", validate),
			zap.Bool("healthcheck", healthCheck))
		helper.GetFlowService().Run(provision, validate, healthCheck, logLevel)
		log.Info("[KLOPAC FLOW - END]")
	}
}
