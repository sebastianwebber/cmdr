package cmdr

import (
	"testing"
)

var (
	okCmdList = []Command{
		{
			Command: "ls",
			Args:    []string{"-lh"},
		},
		{
			Command: "pwd",
		},
	}
	errCmdList = []Command{
		{
			Command: "ls",
			Args:    []string{"-lh"},
		},
		{
			Command: "agdsgsdgdsa 64323adgsgads  y42382842",
		},
		{
			Command: "pwd",
		},
	}
)

func Test_abortStrategy(t *testing.T) {
	type args struct {
		cmdList []Command
	}
	tests := []struct {
		name          string
		args          args
		wantExecCount int
		wantErr       bool
	}{
		{
			name: "2 commands without errors",
			args: args{
				cmdList: okCmdList,
			},
			wantExecCount: len(okCmdList),
			wantErr:       false,
		},
		{
			name: "3 commands 2nd will fail",
			args: args{
				cmdList: errCmdList,
			},
			wantExecCount: 1,
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotExecCount, err := abortStrategy(tt.args.cmdList)
			if (err != nil) != tt.wantErr {
				t.Errorf("abortGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotExecCount != tt.wantExecCount {
				t.Errorf("abortGroup() = %v, want %v", gotExecCount, tt.wantExecCount)
			}
		})
	}
}

func TestGroup(t *testing.T) {
	type args struct {
		strategy Strategy
		cmdList  []Command
	}
	tests := []struct {
		name          string
		args          args
		wantExecCount int
		wantErr       bool
	}{
		{
			name: "this should be ok",
			args: args{
				strategy: AbortOnError,
				cmdList:  okCmdList,
			},
			wantExecCount: len(okCmdList),
			wantErr:       false,
		},
		{
			name: "2nd command will abort",
			args: args{
				strategy: AbortOnError,
				cmdList:  errCmdList,
			},
			wantExecCount: 1,
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotExecCount, err := Group(tt.args.strategy, tt.args.cmdList...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Group() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotExecCount != tt.wantExecCount {
				t.Errorf("Group() = %v, want %v", gotExecCount, tt.wantExecCount)
			}
		})
	}
}
