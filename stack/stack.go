package stack

type Stack []uint8

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(x uint8) {
	*s = append(*s, x)
}

// remove and return top element of stack, return false if stack is empty
func (s *Stack) Pop() (uint8, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	i := len(*s) - 1
	x := (*s)[i]
	*s = (*s)[:i]

	return x, true
}

// return top element of stack, return false if stack is empty
func (s *Stack) Peek() (uint8, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	i := len(*s) - 1
	x := (*s)[i]

	return x, true
}

func (s *Stack) MStore() {

}

func (s *Stack) Add() {

}

func (s *Stack) Mul() {

}

func (s *Stack) SDiv() {

}

func (s *Stack) Exp() {

}
