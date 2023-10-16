package main

import (
	"log"
	"os"
)

func main() {
	// Take arguments from command line
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatalln("Usage: generate_ast <output_directory")
	}
	outputDir := args[0]

	defineAst(outputDir, "Expr", []string{
		"Binary   : Expr left, Token operator, Expr right",
		"Grouping : Expr expression",
		"Literal  : Object value",
		"Unary    : Token operator, Expr right",
	})
}

func defineAst(outputDir, baseName string, types []string) {
	path := outputDir + "/" + baseName + ".go"
	content :=
		`package go-interpreter
	
	`
	os.WriteFile(path, []byte(content), 0644)
}
