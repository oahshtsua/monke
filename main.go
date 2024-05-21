package main

import (
	"fmt"
	"os"
	"os/user"

	"codeberg.org/oahshtsua/monke/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Welcome %s to the Monkey repl.\n", user.Username)
	fmt.Println("Type in Monkey commands")
	repl.Start(os.Stdin, os.Stdout)
}
