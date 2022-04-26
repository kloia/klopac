package flow

import (
	"entrypoint/pkg/logger"
	"entrypoint/pkg/shell"
	"fmt"
	"strings"

	"go.uber.org/zap"
)

type Flow interface {
	ExecuteCommand(command string)
	Run(provision, validate, healthCheck bool, logLevel string)
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
	err, out, _ := p.shell.Run(command)
	if err != nil {
		log.Info(out)
		log.Panic(fmt.Sprint(err))
	}
	log.Info(out)
}

// It basically take some sort of args like (provision, validate, healthCheck, logLevel) and depending to its value it execute relative yaml files.
func (p flowService) Run(provision, validate, healthCheck bool, logLevel string) {
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
		sh controller.sh
		`, logLevel, healthCheck))
		log.Info("[CONTROLLER - END]")
	}
	if !provision && !validate || healthCheck {
		log.Info("[FINALIZER - START]")
		p.ExecuteCommand(fmt.Sprintf(`
		export LOGLEVEL=%v;
		export HEALTHCHECK=%v
		cd finalizer;
		ansible-playbook finalizer.yml;
	`, logLevel, healthCheck))
		log.Info("[FINALIZER - END]")
	}
}

func NewFlowService(s shell.Shell) Flow {
	return &flowService{shell: s}
}
