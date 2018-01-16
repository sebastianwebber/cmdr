package cmdr

import (
	"fmt"
)

func abortStrategy(cmdList []Command) (execCount int, err error) {
	for _, cmd := range cmdList {
		_, cmdErr := cmd.Run()

		if cmdErr != nil {
			err = fmt.Errorf("Error running a command: %v", cmdErr)
			break
		}

		execCount++
	}
	return
}

// Strategy contains details about action to take when a
// command fail inside a group execution
type Strategy (func([]Command) (int, error))

var (
	// AbortOnError is a strategy to run grouped commands
	// This stategy end the group execution resulting a error and output.
	// It is the default option.
	AbortOnError Strategy = abortStrategy

	// // IgnoreOnError is a strategy to run grouped commands
	// // This strategy ignores the problem
	// IgnoreOnError Strategy = "IGNORE"
)

// Group allows to run lots of commands in sequence
// The strategy defines the behavior when a error occours
func Group(strategy Strategy, cmdList ...Command) (execCount int, err error) {
	return strategy(cmdList)
}
