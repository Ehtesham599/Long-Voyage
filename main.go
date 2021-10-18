package main

import (
	"fmt"
	"long-voyage/execute"
	"os"
	"reflect"
)

func main() {

	if len(os.Args) <= 1 {
		panic("Please provide a bytecode as an argument")
	}

	// Get bytecode from cmd argument
	bytecode := os.Args[1]

	// check for valid bytecode
	if !execute.IsValidHexStr(bytecode) {
		panic("Invalid bytecode!")
	}

	//decode retrieved string to byte slice
	decodedBytes := execute.StrToByte(bytecode)

	execute.Execute(decodedBytes)

	fmt.Println(decodedBytes, reflect.TypeOf(decodedBytes))
}
