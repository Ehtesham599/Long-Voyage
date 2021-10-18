package main

import (
	"long-voyage/execute"
	"os"
)

func main() {

	if len(os.Args) <= 1 {
		panic("Please provide a bytecode as an argument")
	}

	// Get bytecode from cmd argument
	bytecode := os.Args[1]

	execute.Execute(bytecode)
}
