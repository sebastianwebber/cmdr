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
		timeout       int
		useShell      bool
	}{
		{name: "run empty command", cmd: "", expectedError: "Missing command name"},
		{name: "run fake command", cmd: "xl8t23hj6__68235896253gf3", expectedError: "Command not found in PATH"},
		{name: "run valid command with wrong args", cmd: "ls", args: []string{"-lh52313252362336", "2133266324"}, expectedError: "Error running a command"},
		{name: "this should be ok", cmd: "ls", args: []string{"-lh"}},
		{name: "this should be ok on bash too", cmd: "ls", args: []string{"-lh"}, useShell: true},
		{name: "this should abort by timemout", cmd: "sleep", args: []string{"5"}, timeout: 2, expectedError: "Error running a command"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			cmd := Command{Command: test.cmd, Args: test.args}

			if test.timeout > 0 {
				cmd.Options.Timeout = test.timeout
			}

			if test.useShell {
				cmd.Options.UseShell = test.useShell
			}

			_, err := RunCmd(cmd)

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
