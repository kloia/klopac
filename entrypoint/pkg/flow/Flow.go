package flow

import (
	"entrypoint/pkg/logger"
	"entrypoint/pkg/shell"
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
)

type Flow interface {
	ExecuteCommand(command string)
	Run(provision, validate, healthCheck bool, logLevel string, file bool)
}

type flowService struct {
	shell shell.Shell
}

// It uses the shell type to execute commands

func (p flowService) ExecuteCommand(command string) {
	log := logger.GetLogger()
	replacer := strings.NewReplacer("\n", "", "\t", " ")
	command = replacer.Replace(command)
	log.Info("Running command", zap.String("command", command))
	err, _, _ := p.shell.Run(command)
	if err != nil {
		log.Panic(fmt.Sprint(err))
	}
}

// It basically take some sort of args like (provision, validate, healthCheck, logLevel) and depending to its value it execute relative yaml files.
func (p flowService) Run(provision, validate, healthCheck bool, logLevel string, file bool) {
	log := logger.GetLogger()
	if !healthCheck || validate || provision {
		log.Info("[PROVISIONER - START]")
		p.ExecuteCommand(fmt.Sprintf(`
		export LOGLEVEL=%v;
		cd provisioner;
		ansible-playbook provisioner.yml;
	`, logLevel))
		log.Info("[PROVISIONER - END]")
	}
	if !provision && !healthCheck || validate {
		log.Info("[VALIDATOR - START]")
		p.ExecuteCommand(fmt.Sprintf(`
		export LOGLEVEL=%v;
		cd validator;
		ansible-playbook validator.yml;
	`, logLevel))
		log.Info("[VALIDATOR - END]")
	}
	if !provision && !validate && !healthCheck {
		log.Info("[CONTROLLER - START]")
		p.ExecuteCommand(fmt.Sprintf(`
		export LOGLEVEL=%v;
		export HEALTHCHECK=%v;
		cd controller;
		bash controller.sh
		`, logLevel, healthCheck))
		log.Info("[CONTROLLER - END]")
	}
	if !provision && !validate || healthCheck {
		log.Info("[FINALIZER - START]")
		p.ExecuteCommand(fmt.Sprintf(`
		export LOGLEVEL=%v;
		export HEALTHCHECK=%v;
		cd finalizer;
		ansible-playbook finalizer.yml -e COMPRESS_BUNDLE=%v;
	`, logLevel, healthCheck, !file))
		log.Info("[FINALIZER - OUTPUT]")
		dat, err := os.ReadFile("/data/bundle/output.md")
		if err != nil {
			log.Panic("Could not read file")
		}
		log.Info(string(dat))
		log.Info("[FINALIZER - END]")
	}
}

func NewFlowService(s shell.Shell) Flow {
	return &flowService{shell: s}
}
