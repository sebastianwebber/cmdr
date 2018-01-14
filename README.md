# cmdr

`cmdr` (pronounced like a _"commander"_) is a go package to abstract and simplify execution of commands on the operation system. 

## how to use it

Basically create a `Command` and call the `Run` function. Take a look:

```golang
package main

import (
	"fmt"

	"github.com/sebastianwebber/cmdr"
)

func main() {
	cmd := cmdr.New(true, "ls", "-lh", "~/tmp/*")
	if cmd.IsValid() {
		out, err := cmd.Run()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))
	}
}
```

