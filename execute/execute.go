package execute

import (
	"encoding/hex"
	"fmt"
	"long-voyage/instruction"
	"long-voyage/stack"
)

const (
	PUSH = "push"
	ADD  = "add"
	MUL  = "mul"
	SDIV = "sdiv"
	STOP = "stop"
	SHA3 = "sha3"
	Null = "null"
)

var (
	push, add, mul, mstore, sdiv, exp, stop, sha3, nullInstruction instruction.Instruction
)

func init() {
	nullInstruction = *instruction.NewInstruction(Null, 0, 0, 0)
	stop = *instruction.NewInstruction(STOP, 0, 0, 0)
	push = *instruction.NewInstruction(PUSH, 96, 127, 3)
	add = *instruction.NewInstruction(ADD, 1, 1, 3)
	mul = *instruction.NewInstruction(MUL, 2, 2, 5)
	sdiv = *instruction.NewInstruction(SDIV, 5, 5, 5)
	sha3 = *instruction.NewInstruction(SHA3, 32, 32, 0)
	//TODO: initialize for exo and mstore for dynamic gas pricing
}

// Check if input hex string is valid
func IsValidHexStr(str string) bool {
	if str != "" {
		return len(str)%2 == 0
	}

	return false
}

// decode string to byte slice
func StrToByte(str string) []byte {
	data, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return data
}

func GetInstruction(bc uint8) instruction.Instruction {
	if bc >= push.Start && bc <= push.End {
		return push
	} else if bc == add.Start {
		return add
	} else if bc == mul.Start {
		return mul
	} else if bc == sdiv.Start {
		return sdiv
	} else if bc == stop.Start {
		return stop
	}

	return nullInstruction
}

func Execute(data []uint8) {
	if len(data)%2 == 0 {

		var gasFee int
		memory := stack.Stack{}

		for i := 0; i < len(data); i += 2 {
			switch inst := GetInstruction(data[i]); inst.Name {
			case PUSH:
				memory.Push(data[i+1])
				gasFee += inst.Gas
			}
		}
		fmt.Println(memory, gasFee)
	}
}
