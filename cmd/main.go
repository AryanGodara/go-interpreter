package main

import (
	"fmt"
	"os"

	glox "github.com/aryangodara/go-interpreter/lox"
)

func main() {
	// Take arguments from command line
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("Usage: glox [script]")
	} else if len(args) == 1 {
		glox.RunFile(args[0])
	} else {
		glox.RunPrompt()
	}
}
