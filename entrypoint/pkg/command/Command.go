package command

import "os/exec"

type Command interface {
	Exec(command string) *exec.Cmd
}

type commandService struct {
}

// Command function to make use of os.Exec() and here we create the template of the shell command.
func (c commandService) Exec(command string) *exec.Cmd {
	return exec.Command("bash", "-c", command)
}

func NewCommandService() Command {
	return &commandService{}
}
