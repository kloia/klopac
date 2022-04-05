package flow

import (
	"entrypoint/pkg/shell"
	"fmt"
	"log"
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
	err, out, errout := p.shell.Run(command)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	log.Print("--- stdout ---")
	log.Print(out)
	log.Print("--- stderr ---")
	log.Print(errout)
}

// It basically take some sort of args like (provision, validate, healthCheck, logLevel) and depending to its value it execute relative yaml files.
func (p flowService) Run(provision, validate, healthCheck bool, logLevel string) {
	//
	if provision == true || validate == true || healthCheck == true {

		if healthCheck == true {
			p.ExecuteCommand(strings.Trim(fmt.Sprintf(`
					export LOGLEVEL=%v
					export HEALTHCHECK=%v
					cd finalizer;
					ansible-playbook finalizer.yaml;
				`, logLevel, healthCheck), " "))
		} else {
			p.ExecuteCommand(strings.Trim(fmt.Sprintf(`
				export LOGLEVEL=%v
				cd provisioner;
				ansible-playbook provisioner.yaml;
			`, logLevel), " "))

			if validate == true {
				p.ExecuteCommand(strings.Trim(fmt.Sprintf(`
					export LOGLEVEL=%v
					cd validator;
					ansible-playbook validator.yaml;
				`, logLevel), " "))
			}
		}

	} else {

		// execute provisioner
		p.ExecuteCommand(strings.Trim(fmt.Sprintf(`
				export LOGLEVEL=%v
				cd provisioner;
				ansible-playbook provisioner.yaml;
			`, logLevel), " "))

		// execute validator
		p.ExecuteCommand(strings.Trim(fmt.Sprintf(`
					export LOGLEVEL=%v
					cd validator;
					ansible-playbook validator.yaml;
				`, logLevel), " "))

		// execute controller
		p.ExecuteCommand(strings.Trim(fmt.Sprintf(`
		export LOGLEVEL=%v
		export HEALTHCHECK=%v
		cd controller;
		ansible-playbook controller.yaml;
		`, logLevel, healthCheck), " "))

		// execute healhtCheck
		p.ExecuteCommand(strings.Trim(fmt.Sprintf(`
				export LOGLEVEL=%v
				export HEALTHCHECK=%v
				cd finalizer;
				ansible-playbook finalizer.yaml;
			`, logLevel, healthCheck), " "))

	}
}

func NewFlowService(s shell.Shell) Flow {
	return &flowService{shell: s}
}
