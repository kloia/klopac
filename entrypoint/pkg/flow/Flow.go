package flow

import (
	"entrypoint/pkg/logger"
	"entrypoint/pkg/shell"
	"fmt"
	"go.uber.org/zap"
	"strings"
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
	err, out, errout := p.shell.Run(command)
	if err != nil {
		log.Error("Error while executing command", zap.String("command", command))
	}
	log.Debug("Streams",
		zap.String("command", command),
		zap.String("stdout", out),
		zap.String("stderr", errout))
}

// It basically take some sort of args like (provision, validate, healthCheck, logLevel) and depending to its value it execute relative yaml files.
func (p flowService) Run(provision, validate, healthCheck bool, logLevel string) {
	log := logger.GetLogger()
	switch {
	case !healthCheck || validate || provision:
		log.Info("START: PROVISIONER")
		p.ExecuteCommand(fmt.Sprintf(`
		export LOGLEVEL=%v
		cd provisioner;
		ansible-playbook provisioner.yaml;
	`, logLevel))
		log.Info("END: PROVISIONER")
	case !provision && !healthCheck || validate:
		log.Info("START: VALIDATOR")
		p.ExecuteCommand(fmt.Sprintf(`
		export LOGLEVEL=%v
		cd validator;
		ansible-playbook validator.yaml;
	`, logLevel))
		log.Info("END: VALIDATOR")
	case !provision && !validate && !healthCheck:
		log.Info("START: CONTROLLER")
		p.ExecuteCommand(fmt.Sprintf(`
		export LOGLEVEL=%v;
		export HEALTHCHECK=%v;
		cd controller;
		sh controller.sh
		`, logLevel, healthCheck))
		log.Info("END: CONTROLLER")
	case !provision && !validate || healthCheck:
		log.Info("START: FINALIZER")
		p.ExecuteCommand(fmt.Sprintf(`
		export LOGLEVEL=%v
		export HEALTHCHECK=%v
		cd finalizer;
		ansible-playbook finalizer.yaml;
	`, logLevel, healthCheck))
		log.Info("END: FINALIZER")
	default:
		log.Error("Options are wrong. Klopac Flow Failed.")
	}

}

func NewFlowService(s shell.Shell) Flow {
	return &flowService{shell: s}
}
