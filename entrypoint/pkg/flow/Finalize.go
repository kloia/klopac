package flow

import (
	"awesomeProject/pkg/shell"
	"log"
)

type Finalize interface {
	Run(command string)
}

type finalizeService struct {
	shell shell.Shell
}

func (f finalizeService) Run(command string) {
	err, out, errout := f.shell.Run(command)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	log.Print("--- stdout ---")
	log.Print(out)
	log.Print("--- stderr ---")
	log.Print(errout)
}

func NewFinalizeService(s shell.Shell) Finalize {
	return &finalizeService{shell: s}
}
