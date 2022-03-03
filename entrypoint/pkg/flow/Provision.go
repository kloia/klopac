package flow

import (
	"awesomeProject/pkg/shell"
	"log"
)

type Provision interface {
	Run(command string)
}

type provisionService struct {
	shell shell.Shell
}

func (p provisionService) Run(command string) {
	err, out, errout := p.shell.Run(command)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	log.Print("--- stdout ---")
	log.Print(out)
	log.Print("--- stderr ---")
	log.Print(errout)
}

func NewProvisionService(s shell.Shell) Provision {
	return &provisionService{shell: s}
}
