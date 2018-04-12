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
		{name: "sed complex", cmd: `sed -i 's|-Xmx.*|-Xmx3G|' teste.sh`, wantOut: &Command{Command: "sed", Args: []string{"-i", "s|-Xmx.*|-Xmx3G|", "teste.sh"}}},
		{name: "exception double quotes", cmd: `/opt/wildfly-10.1.0.Final/bin/jboss-cli.sh -c --command="read-attribute server-state"`, wantOut: &Command{Command: "/opt/wildfly-10.1.0.Final/bin/jboss-cli.sh", Args: []string{`-c`, `--command="read-attribute server-state"`}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := Parse(tt.cmd); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("Parse() = %#v, want %#v", gotOut, tt.wantOut)
			}
		})
	}
}
