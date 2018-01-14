package cmdr

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func getPATH() []string {
	return strings.Split(os.Getenv("PATH"), ":")
}

func findInPath(cmd string) (found bool) {

	for _, dir := range getPATH() {

		fullPath := fmt.Sprintf("%s/%s", dir, cmd)

		if fileExist(fullPath) {
			found = true
			break
		}
	}

	return
}

func fileExist(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

// RunCmd runs a command in the operating system
func RunCmd(c Command) ([]byte, error) {
	return runCmd(c)
}

func printCommand(cmd *exec.Cmd) {
	log.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		log.Printf("==> Output: %s\n", string(outs))
	}
}

func runCmd(c Command) (output []byte, err error) {

	err = validateCmd(c)
	if err != nil {
		return
	}

	var cmd *exec.Cmd

	if c.Options.UseShell {
		cmd = exec.Command("bash", "-c", fmt.Sprintf("%s %s", c.Command, strings.Join(c.Args, " ")))
	} else {
		cmd = exec.Command(c.Command, c.Args...)
	}

	output, err = cmd.CombinedOutput()

	if err != nil {
		err = fmt.Errorf("Error running a command: %v", err)
	}

	return
}

func validateCmd(c Command) (err error) {

	if c.Command == "" {
		err = fmt.Errorf("Missing command name")
		return
	}

	if !findInPath(c.Command) {
		err = fmt.Errorf("Command not found in PATH")
		return
	}

	return
}
