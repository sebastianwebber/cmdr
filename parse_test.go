package cmdr

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		cmd     string
		wantOut *Command
	}{
		{name: "simple spaces", cmd: "echo hello world", wantOut: &Command{Command: "echo", Args: []string{"hello", "world"}}},
		{name: "simple single quotes", cmd: `echo 'hello world'`, wantOut: &Command{Command: "echo", Args: []string{"hello world"}}},
		{name: "simple double quotes", cmd: `echo "hello world"`, wantOut: &Command{Command: "echo", Args: []string{"hello world"}}},
		{name: "complex single quotes", cmd: `psql -U postgres -c 'select now() as "agora"'`, wantOut: &Command{Command: "psql", Args: []string{"-U", "postgres", "-c", `select now() as "agora"`}}},
		{name: "complex double quotes", cmd: `psql -U postgres -c "select 'seba' as nome"`, wantOut: &Command{Command: "psql", Args: []string{"-U", "postgres", "-c", `select 'seba' as nome`}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := Parse(tt.cmd); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("Parse() = %#v, want %#v", gotOut, tt.wantOut)
			}
		})
	}
}
