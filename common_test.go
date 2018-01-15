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
		{
			name: "compare",
			args: args{
				useShell: false,
				cmd:      "ls",
				args:     []string{"-lh"}},
			want: Command{
				Options: Options{UseShell: false},
				Command: "ls",
				Args:    []string{"-lh"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.useShell, tt.args.cmd, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommand_IsValid(t *testing.T) {
	type fields struct {
		Command string
		Args    []string
		Options Options
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "invalid empty",
			fields: fields{
				Command: "",
			},
			want: false,
		},
		{
			name: "invalid notfound",
			fields: fields{
				Command: "jt23g6emdsbxzgmvksdyg7089 v3g4069tygkahmxzbvuweg5 t",
			},
			want: false,
		},
		{
			name: "valid",
			fields: fields{
				Command: "ls",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Command{
				Command: tt.fields.Command,
				Args:    tt.fields.Args,
				Options: tt.fields.Options,
			}
			if got := c.IsValid(); got != tt.want {
				t.Errorf("Command.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommand_Run(t *testing.T) {
	type fields struct {
		Command string
		Args    []string
		Options Options
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "simple echo",
			fields: fields{
				Command: "echo",
				Args:    []string{"hello"},
			},
			want:    []byte("hello\n"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Command{
				Command: tt.fields.Command,
				Args:    tt.fields.Args,
				Options: tt.fields.Options,
			}
			got, err := c.Run()
			if (err != nil) != tt.wantErr {
				t.Errorf("Command.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Command.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
