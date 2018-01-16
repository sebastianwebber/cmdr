package cmdr

import (
	"fmt"
)

// Strategy contais details about action to take when a
// command fail inside a group execution
type Strategy string

const (
	// AbortOnError is a strategy to run grouped commands
	// This stategy end the group execution resulting a error and output.
	// It is the default option.
	AbortOnError Strategy = "ABORT"

	// // IgnoreOnError is a strategy to run grouped commands
	// // This strategy ignores the problem
	// IgnoreOnError Strategy = "IGNORE"
)

// Group allows to run lots of commands in sequence
// The strategy defines the behavior when a error occours
func Group(strategy Strategy, cmdList ...Command) (execCount int, err error) {

	var rFunc (func([]Command) (int, error))

	// if strategy == AbortOnError {
	rFunc = abortGroup
	// }

	return rFunc(cmdList)
}

func abortGroup(cmdList []Command) (execCount int, err error) {
	for _, cmd := range cmdList {
		out, cmdErr := cmd.Run()

		if cmdErr != nil {
			err = fmt.Errorf("Error running a command: %v", cmdErr)
			message := string(out)
			if len(out) > 0 {
				err = fmt.Errorf("%v Output: %v", err, message)
			}
			break
		}

		execCount++
	}
	return
}
