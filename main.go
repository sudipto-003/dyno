package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/sudipto-003/dyno/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the DYNO programming language!\n", user.Username)
	fmt.Printf("Thank you for testing DYNO, feel free to type commands.\n")

	repl.Start(os.Stdin, os.Stdout)
}
