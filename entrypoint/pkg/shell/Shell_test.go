package shell

import (
	"awesomeProject/pkg/command"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShell(t *testing.T) {
	t.Run("execute command succesfully", func(t *testing.T) {
		tcs := []struct {
			given           string
			expected_err    error
			expected_stdout string
			expected_stderr string
		}{
			{given: "echo hi", expected_err: nil, expected_stdout: "hi\n", expected_stderr: ""},
		}
		for _, tc := range tcs {
			err, stdout, stderr := NewShellService(command.NewCommandService()).Run(tc.given)
			assert.Equal(t, tc.expected_err, err)
			assert.Equal(t, tc.expected_stdout, stdout)
			assert.Equal(t, tc.expected_stderr, stderr)
		}
	})
}
