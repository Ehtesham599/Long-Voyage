package execute

import (
	"encoding/hex"
	"errors"
	"fmt"
	"long-voyage/instruction"
	"long-voyage/operation"
)

const (
	PUSH   = "push"
	ADD    = "add"
	MUL    = "mul"
	SDIV   = "sdiv"
	STOP   = "stop"
	SHA3   = "sha3"
	EXP    = "exp"
	MSTORE = "mstore"
	NULL   = "null"
)

var (
	push, add, mul, mstore, sdiv, exp, stop, sha3, nullInstruction instruction.Instruction
)

type Result struct {
	Bytecode      string
	Keccak256Hash string
	GasFee        int
}

func init() {
	nullInstruction = *instruction.NewInstruction(NULL, 0, 0, 0)
	stop = *instruction.NewInstruction(STOP, 0, 0, 0)
	push = *instruction.NewInstruction(PUSH, 96, 127, 3)
	add = *instruction.NewInstruction(ADD, 1, 1, 3)
	mul = *instruction.NewInstruction(MUL, 2, 2, 5)
	sdiv = *instruction.NewInstruction(SDIV, 5, 5, 5)
	sha3 = *instruction.NewInstruction(SHA3, 32, 32, 0)
	exp = *instruction.NewInstruction(EXP, 10, 10, 50)
	mstore = *instruction.NewInstruction(MSTORE, 82, 83, 3)
}

// Check if input hex string is valid
func IsValidHexStr(str string) bool {
	if str != "" {
		return len(str)%2 == 0
	}

	return false
}

// decode string to byte slice
func StrToByte(str string) ([]byte, error) {
	data, err := hex.DecodeString(str)
	if err != nil {
		return []byte{}, errors.New("cannot convert bytecode to byte slice")
	}
	return data, nil
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
	} else if bc == mstore.Start {
		return mstore
	} else if bc == exp.Start {
		return exp
	}

	return nullInstruction
}

func Execute(data string) {
	if len(data) <= 0 {
		panic("Empty bytecode")
	}

	if !IsValidHexStr(data) {
		panic("Invalid bytecode")
	}

	decodedBytes, err := StrToByte(data)

	if err != nil {
		panic(err)
	}

	var gasFee int
	memoryStack := operation.Stack{}
	memory := operation.Stack{}

	for i := 0; i < len(decodedBytes); i++ {
		switch inst := GetInstruction(decodedBytes[i]); inst.Name {

		case PUSH:
			memoryStack.Push(decodedBytes[i+1])
			gasFee += inst.Gas
			i += 1
			fmt.Printf("Pushed %v to Stack for %v\n", decodedBytes[i], gasFee)

		case ADD:
			memoryStack.Add()
			gasFee += inst.Gas
			fmt.Println("Added stack values for ", gasFee)

		case MUL:
			memoryStack.Mul()
			gasFee += inst.Gas
			fmt.Println("Mul stack values for ", gasFee)

		case SDIV:
			memoryStack.SDiv()
			gasFee += inst.Gas
			fmt.Println("Div stack values for ", gasFee)

		case EXP:
			_, exponent, _ := memoryStack.Exp()
			gasFee += inst.Gas * int(exponent)
			fmt.Println("Exp stack values for ", gasFee)

		case MSTORE:
			temp, offset := memory.MStore(&memoryStack)
			fmt.Println(temp, offset)

		case NULL:
			panic("unknown instruction")
		}
	}
	fmt.Println("FINAL: ", memoryStack, gasFee)

}
