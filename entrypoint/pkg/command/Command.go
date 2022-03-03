package command

import "os/exec"

type Command interface {
	Exec(command string) *exec.Cmd
}

type commandService struct {
}

func (c commandService) Exec(command string) *exec.Cmd {
	return exec.Command("bash", "-c", command)
}

func NewCommandService() Command {
	return &commandService{}
}
