package main

import (
	"fmt"
	"long-voyage/decode"
	"os"
)

func main() {

	if len(os.Args) <= 1 {
		panic("Please provide a bytecode as an argument")
	}

	// Get bytecode from cmd argument
	bytecode := os.Args[1]

	// check for valid bytecode
	if !decode.IsValidHexString(bytecode) {
		panic("Invalid bytecode!")
	}

	//decode retrieved string to byte slice
	decodedBytes := decode.StrToByteSlice(bytecode)

	fmt.Println(decodedBytes)
}
