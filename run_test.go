package cmdr

import (
	"strings"
	"testing"
)

func TestRunCmd(t *testing.T) {
	tests := []struct {
		name          string
		cmd           string
		args          []string
		expectedError string
	}{
		{name: "run empty command", cmd: "", expectedError: "Missing command name"},
		{name: "run fake command", cmd: "xl8t23hj6__68235896253gf3", expectedError: "Command not found in PATH"},
		{name: "run valid command with wrong args", cmd: "ls", args: []string{"-lh52313252362336", "2133266324"}, expectedError: "Error running a command"},
		{name: "this should be ok", cmd: "ls", args: []string{"-lh"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			_, err := RunCmd(Command{Command: test.cmd, Args: test.args})

			if test.expectedError != "" {

				if err == nil {
					t.Errorf("Expected '%v' error", test.expectedError)
				}

				if err != nil {
					// first test if not starts with the message
					if !strings.HasPrefix(err.Error(), test.expectedError) {
						// if not, then check if its different
						if err.Error() != test.expectedError {
							t.Errorf("Expected '%v' but returned '%v", test.expectedError, err)
						}
					}
				}

			}
		})
	}
}
