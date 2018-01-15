# cmdr
[![Build Status](https://travis-ci.org/sebastianwebber/cmdr.svg?branch=master)](https://travis-ci.org/sebastianwebber/cmdr) [![Go Report Card](https://goreportcard.com/badge/github.com/sebastianwebber/cmdr)](https://goreportcard.com/report/github.com/sebastianwebber/cmdr) [![codecov](https://codecov.io/gh/sebastianwebber/cmdr/branch/master/graph/badge.svg)](https://codecov.io/gh/sebastianwebber/cmdr) 


`cmdr` (pronounced  _"commander"_) is a go package to abstract and simplify execution of commands on the operation system. 

## how to use it

Basically create a `Command` and call the `Run` function. Take a look:

```golang
package main

import (
    "fmt"

    "github.com/sebastianwebber/cmdr"
)

func main() {
    
    // New is a helper to create a Command
    // You can call it by a shell like bash if you want (useful to process expressions like *)
    cmd := cmdr.New(true, "ls", "-lh", "~/tmp/*")
    
    // You can declare the variable as well:
    // cmd := cmdr.Command{  }

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


## TODO List

- [x] Add option to timeout
- [ ] Enable way to group commands
- [ ] Pipe support
- [ ] add support por multiple commands

