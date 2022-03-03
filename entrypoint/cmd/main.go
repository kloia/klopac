package main

import (
	"entrypoint/pkg/command"
	"entrypoint/pkg/flag"
	"entrypoint/pkg/flow"
	"entrypoint/pkg/option"
	"entrypoint/pkg/shell"
	"entrypoint/pkg/websocket"
	"fmt"
	"strings"
)

func main() {
	options := option.NewOptionService(flag.NewFlagService()).Get()
	webSocketEnabled := options["websocket"].(*bool)
	if *webSocketEnabled {

		websocket.Enable(*options["uri"].(*string), *options["username"].(*string), *options["password"].(*string))
	} else {
		shellService := shell.NewShellService(command.NewCommandService())
		if *options["provision"].(*bool) || *options["validate"].(*bool) {
			flow.NewProvisionService(shellService).Run(strings.Trim(fmt.Sprintf(`
				export LOGLEVEL=%v
				cd provisioner;
				ls;
			`, *options["loglevel"].(*string)), " "))
			//ansible-playbook provisioner.yaml
			if *options["validate"].(*bool) {
				flow.NewValidateService(shellService).Run(strings.Trim(fmt.Sprintf(`
					export LOGLEVEL=%v				
					cd validator;
					ls;
				`, *options["loglevel"].(*string)), " "))
				//ansible-playbook validator.yaml
			}
		}
		if *options["healthcheck"].(*bool) {
			flow.NewFinalizeService(shellService).Run(strings.Trim(fmt.Sprintf(`
				export LOGLEVEL=%v
				export HEALTHCHECK=%v
				cd finalizer;
				ls;
			`, *options["loglevel"].(*string), *options["healthcheck"].(*bool)), " "))
			//ansible-playbook finalizer.yaml
		}
	}
}
