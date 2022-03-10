package flow

import (
	"entrypoint/mock"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestProvision(t *testing.T) {
	t.Run("execute provision succesfully", func(t *testing.T) {
		tcs := []struct {
			given           string
			expected_err    error
			expected_stdout string
			expected_stderr string
		}{
			{given: "echo hi", expected_err: nil, expected_stdout: "hi\n", expected_stderr: ""},
		}
		mockShell := mock.NewMockShell(gomock.NewController(t))
		for _, tc := range tcs {
			mockShell.EXPECT().Run(tc.given).Return(tc.expected_err, tc.expected_stdout, tc.expected_stderr)
			NewProvisionService(mockShell).Run(tc.given)
		}

	})

}