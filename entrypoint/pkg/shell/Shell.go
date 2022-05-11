package shell

import (
	"bytes"
	cmdService "entrypoint/pkg/command"
	"io"
	"os"
)

type Shell interface {
	Run(command string) (error, string, string)
}

type shellService struct {
	command cmdService.Command
}

// It runs the command string and return its result as output
func (s shellService) Run(command string) (error, string, string) {
	cmd := s.command.Exec(command)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	err := cmd.Run()

	return err, string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
}

func NewShellService(c cmdService.Command) Shell {
	return &shellService{command: c}
}
