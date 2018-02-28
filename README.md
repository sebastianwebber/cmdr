# cmdr
[![Build Status](https://travis-ci.org/sebastianwebber/cmdr.svg?branch=master)](https://travis-ci.org/sebastianwebber/cmdr) [![Go Report Card](https://goreportcard.com/badge/github.com/sebastianwebber/cmdr)](https://goreportcard.com/report/github.com/sebastianwebber/cmdr) [![codecov](https://codecov.io/gh/sebastianwebber/cmdr/branch/master/graph/badge.svg)](https://codecov.io/gh/sebastianwebber/cmdr)


`cmdr` (pronounced  _"commander"_) is a go package to abstract and simplify execution of commands on the operation system.

## how to use it

First things first:
```
go get -u -v go get github.com/sebastianwebber/cmdr
```


Basically create a `Command` and call the `Run` function. Take a look:

```golang
package main

import (
	"fmt"

	"github.com/sebastianwebber/cmdr"
)

func main() {
    // *** short version ***********
	out, err := cmdr.New(true, "ls", "-lh", "~/tmp2/*").Run()
	fmt.Println("Output:", string(out))
	if err != nil {
		fmt.Println("OOPS:", err.Error())
    }

    // *** verbose version ***********
	// New is a helper to create a Command
	// You can call it by a shell like bash if you want (useful to process expressions like *)
	cmd := cmdr.New(true, "ls", "-lh", "~/tmp/*")

	// You can declare the variable as well:
	// cmd := cmdr.Command{  }

	// You can also parse a command into a Command:
	// cmd := cmdr.Parse(`psql -At -c 'select now();'`)

	// Enable timeout if you want (5s by example)
	cmd.Options.Timeout = 5

	// To check if the inputed command is valid, use the IsValid function.
	// It checks if the command exists in PATH
	if cmd.IsValid() {

		// To execute the command, just call the Run function
		out, err := cmd.Run()
		if err != nil {
			panic(err)
		}

		// here comes the output
		fmt.Println(string(out))
	}
}
```

## Grouping commands

Its possible to group a list of commands:

```golang
package main

import (
    "fmt"

    "github.com/sebastianwebber/cmdr"
)

func main() {
    // Group options (experimental)
    total, err := cmdr.Group(
        cmdr.AbortOnError,
        cmdr.New(false, "ls", "-lh"),
        cmdr.New(false, "pwd 123q6236"),
        cmdr.New(false, "cat", "/etc/hosts"),
    )
    fmt.Printf("%d commands executed without error. \n", total)

    if err != nil {
        fmt.Printf("Houston, we have a problem! %v\n", err)
    }
}
```
> **This is a work in progress.**


## TODO List

- [x] Add option to timeout
- [x] Enable way to group commands
- [ ] Print output of each command in the group (perhaps adding a `name` option?)
- [ ] Pipe support
- [ ] add support por multiple commands

