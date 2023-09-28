package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Take arguments from command line
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("Usage: glox [script]")
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}

func runFile(path string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln("Error reading file")
	}
	run(string(bytes))
}

func runPrompt() {
	for {
		fmt.Println("> ")
		var line string
		_, err := fmt.Scanln(line)
		if err != nil {
			break
		}
		run(line)
	}
}

func run(source string) {
	tokens := strings.Split(source, " ")
	for _, token := range tokens {
		fmt.Println(token)
	}
}
