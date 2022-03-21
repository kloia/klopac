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

func (p flowService) Run(provision, validate, healthCheck bool, logLevel string) {
	if provision == true || validate == true {
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
	if healthCheck == true {
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
