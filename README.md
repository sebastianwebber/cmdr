# cmdr

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
    ///// func New(useShell bool, cmd string, args ...string) Command {}
    cmd := cmdr.New(true, "ls", "-lh", "~/tmp/*")
    
    // You can declare the variable as well:
    // cmd := cmdr.Command{  }


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

