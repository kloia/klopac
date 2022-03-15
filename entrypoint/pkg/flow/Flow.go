package flow

import (
	"entrypoint/pkg/shell"
	"log"
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
