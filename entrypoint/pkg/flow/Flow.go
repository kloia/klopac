package flow

import (
	"entrypoint/pkg"
	"entrypoint/pkg/option"
	"entrypoint/pkg/shell"
	"entrypoint/pkg/websocket"
	"fmt"
	"log"
	"strings"
)

type Flow interface {
	Run(command string)
}

type flowService struct {
	shell shell.Shell
}

func (p flowService) Run(command string) {
	err, out, errout := p.shell.Run(command)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	log.Print("--- stdout ---")
	log.Print(out)
	log.Print("--- stderr ---")
	log.Print(errout)
}

func NewFlowService(s shell.Shell) Flow {
	return &flowService{shell: s}
}

func Run() {
	webSocketEnabled := option.GetParam[bool]("websocket")
	if webSocketEnabled == true {
		uri := option.GetParam[string]("uri")
		username := option.GetParam[string]("username")
		password := option.GetParam[string]("password")
		websocket.Enable(uri, username, password)
	} else {
		provision := option.GetParam[bool]("provision")
		validate := option.GetParam[bool]("validate")
		logLevel := option.GetParam[string]("loglevel")
		healthCheck := option.GetParam[bool]("healthcheck")
		OpenSource(provision, validate, healthCheck, logLevel)
	}

}

func OpenSource(provision, validate, healthCheck bool, logLevel string) {
	if provision == true || validate == true {
		pkg.FlowService.Run(strings.Trim(fmt.Sprintf(`
				export LOGLEVEL=%v
				cd provisioner;
				ansible-playbook provisioner.yaml;
			`, logLevel), " "))
		if validate == true {
			pkg.FlowService.Run(strings.Trim(fmt.Sprintf(`
					export LOGLEVEL=%v
					cd validator;
					ansible-playbook validator.yaml;
				`, logLevel), " "))
		}
	}
	if healthCheck == true {
		pkg.FlowService.Run(strings.Trim(fmt.Sprintf(`
				export LOGLEVEL=%v
				export HEALTHCHECK=%v
				cd finalizer;
				ansible-playbook finalizer.yaml;
			`, logLevel, healthCheck), " "))
	}
}
