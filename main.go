package main

import (
	"fmt"
	"os"
	user2 "os/user"

	"github.com/oriiyx/moonk-language/repl"
)

func main() {
	user, err := user2.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Moonk programming language!\n", user.Username)
	fmt.Printf("Type your commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
