package shell

import (
	"bytes"
	cmdService "entrypoint/pkg/command"
	"log"
)

type Shell interface {
	Run(command string) (error, string, string)
}

type shellService struct {
	command cmdService.Command
}

// It runs the command string and return its result as output
func (s shellService) Run(command string) (error, string, string) {
	log.Print(command)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := s.command.Exec(command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal("Error while running command: ", command, "\nError output: ", err.Error())
	}
	return err, stdout.String(), stderr.String()
}

func NewShellService(c cmdService.Command) Shell {
	return &shellService{command: c}
}
