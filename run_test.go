package cmdr

import (
	"os/exec"
	"reflect"
	"strings"
	"testing"
)

func TestRunCmd(t *testing.T) {
	tests := []struct {
		name          string
		cmd           string
		args          []string
		expectedError string
		Options
	}{
		{name: "run empty command", cmd: "", expectedError: "Missing command name", Options: Options{CheckPath: true}},
		{name: "run fake command", cmd: "xl8t23hj6__68235896253gf3", expectedError: "Command not found in PATH", Options: Options{CheckPath: true}},
		{name: "run valid command with wrong args", cmd: "ls", args: []string{"-lh52313252362336", "2133266324"}, expectedError: "Error running a command"},
		{name: "this should be ok", cmd: "ls", args: []string{"-lh"}},
		{name: "this should be ok on bash too", cmd: "ls", args: []string{"-lh"}, Options: Options{UseShell: true}},
		{name: "this should abort by timemout", cmd: "sleep", args: []string{"2"}, expectedError: "Error running a command", Options: Options{Timeout: 1}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			cmd := Command{Command: test.cmd, Args: test.args}

			cmd.Options = test.Options

			_, err := RunCmd(cmd)

			if test.expectedError != "" {

				if err == nil {
					t.Errorf("Expected '%v' error", test.expectedError)
				}

				if err != nil {
					// first test if not starts with the message
					if !strings.Contains(err.Error(), test.expectedError) {
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

func Test_makeCmd(t *testing.T) {
	type args struct {
		c Command
	}
	tests := []struct {
		name    string
		args    args
		wantCmd *exec.Cmd
	}{
		{
			name: "with shell",
			args: args{
				c: Command{
					Command: "ls",
					Args:    []string{"-lh"},
					Options: Options{
						UseShell: true,
					},
				},
			},
			wantCmd: exec.Command("bash", "-c", "ls -lh"),
		},
		{
			name: "without shell",
			args: args{
				c: Command{
					Command: "ls",
					Args:    []string{"-lh"},
				},
			},
			wantCmd: exec.Command("ls", "-lh"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCmd := makeCmd(tt.args.c); !reflect.DeepEqual(gotCmd, tt.wantCmd) {
				t.Errorf("makeCmd() = %v, want %v", gotCmd, tt.wantCmd)
			}
		})
	}
}

func Test_findInPath(t *testing.T) {

	tests := []struct {
		name      string
		args      string
		wantFound bool
	}{
		{"partial current dir", "./superscript.sh", true},
		{"partial previous path", "../superscript.sh", true},
		{"partial home dir", "~/superscript.sh", true},
		{"full prefix", "/app/test/superscript.sh", true},
		{"random non existing script", "shfjasfhkjsafgasjajsmvxbsjty jaghh", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFound := findInPath(tt.args); gotFound != tt.wantFound {
				t.Errorf("findInPath() = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}

func Test_runCmd(t *testing.T) {
	type args struct {
		c Command
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []byte
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput, err := runCmd(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("runCmd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("runCmd() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
