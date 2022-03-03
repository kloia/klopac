package flow

import (
	"awesomeProject/pkg/shell"
	"log"
)

type Validate interface {
	Run(command string)
}

type validateService struct {
	shell shell.Shell
}

func (v validateService) Run(command string) {
	err, out, errout := v.shell.Run(command)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	log.Print("--- stdout ---")
	log.Print(out)
	log.Print("--- stderr ---")
	log.Print(errout)
}

func NewValidateService(s shell.Shell) Validate {
	return &validateService{shell: s}
}
