package gointerpreter

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var hadError bool = false

func RunFile(path string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln("Error reading file")
	}
	run(string(bytes))
	if hadError {
		os.Exit(65) // 65 means there was an error related to data
	}
}

func RunPrompt() {
	for {
		fmt.Println("> ")
		var line string
		_, err := fmt.Scanln(&line)
		if err != nil {
			break
		}
		run(line)
		hadError = false
	}
}

func run(source string) {
	tokens := strings.Split(source, " ")
	for _, token := range tokens {
		fmt.Println(token)
	}
}

func Error(line int, message string) {
	report(line, "", message)
}

func report(lint int, where, message string) {
	log.Fatalf("[line %d] Error %s : %s\n", lint, where, message)
	hadError = true
}
