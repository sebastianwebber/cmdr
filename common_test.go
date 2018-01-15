package cmdr

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		useShell bool
		cmd      string
		args     []string
	}
	tests := []struct {
		name string
		args args
		want Command
	}{
		{name: "compare", args: args{useShell: false, cmd: "ls", args: []string{"-lh"}}, want: Command{Options: Options{UseShell: false}, Command: "ls", Args: []string{"-lh"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.useShell, tt.args.cmd, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
