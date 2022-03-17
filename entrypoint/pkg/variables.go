package pkg

import (
	"entrypoint/pkg/command"
	"entrypoint/pkg/flag"
	"entrypoint/pkg/flow"
	"entrypoint/pkg/option"
	"entrypoint/pkg/shell"
)

var (
	OptionService = option.NewOptionService(flag.NewFlagService())
	FlowService   = flow.NewFlowService(shell.NewShellService(command.NewCommandService()))
)
