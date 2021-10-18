package instruction

type Instruction struct {
	Name  string
	Start uint8
	End   uint8
	Gas   int
}

func NewInstruction(name string, start uint8, end uint8, gas int) *Instruction {
	instruction := Instruction{
		Name:  name,
		Start: start,
		End:   end,
		Gas:   gas,
	}
	return &instruction
}
