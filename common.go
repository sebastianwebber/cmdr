package cmdr

// Options contains Command configuration
type Options struct {
	UseShell bool
	Timeout  int
}

// Command defines how to call a program
type Command struct {
	Command     string
	Args        []string
	Options     Options
	Description string
}

// New creates a Command
func New(useShell bool, cmd string, args ...string) *Command {
	return &Command{
		Options: Options{UseShell: useShell},
		Command: cmd,
		Args:    args,
	}
}

// Run a Command
func (c *Command) Run() ([]byte, error) {
	return runCmd(*c)
}

// IsValid validates a Command
func (c *Command) IsValid() bool {
	return (validateCmd(*c) == nil)
}
