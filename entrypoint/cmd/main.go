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
		flowService := flow.NewFlowService(shellService)
		if *options["provision"].(*bool) || *options["validate"].(*bool) {
			flowService.Run(strings.Trim(fmt.Sprintf(`
				export LOGLEVEL=%v
				cd provisioner;
				ansible-playbook provisioner.yaml;
			`, *options["loglevel"].(*string)), " "))
			if *options["validate"].(*bool) {
				flow.NewFlowService(shellService).Run(strings.Trim(fmt.Sprintf(`
					export LOGLEVEL=%v				
					cd validator;
					ansible-playbook validator.yaml;
				`, *options["loglevel"].(*string)), " "))
			}
		}
		if *options["healthcheck"].(*bool) {
			flowService.Run(strings.Trim(fmt.Sprintf(`
				export LOGLEVEL=%v
				export HEALTHCHECK=%v
				cd finalizer;
				ansible-playbook finalizer.yaml;
			`, *options["loglevel"].(*string), *options["healthcheck"].(*bool)), " "))
		}
	}
}
