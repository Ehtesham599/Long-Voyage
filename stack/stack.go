package stack

import "math"

type Stack []uint8

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(x uint8) (uint8, bool) {
	*s = append(*s, x)
	return x, true
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

func (m *Stack) MStore(s *Stack) (*Stack, uint8) {
	*m = append(*m, (*s)[:len(*s)-1]...)
	offset, _ := s.Peek()
	return m, offset
}

func (s *Stack) Add() (uint8, bool) {
	if s.IsEmpty() || len(*s) < 2 {
		return 0, false
	}
	i := len(*s) - 1
	j := len(*s) - 2

	x := (*s)[i] + (*s)[j]

	return x, true
}

func (s *Stack) Mul() (uint8, bool) {
	if s.IsEmpty() || len(*s) < 2 {
		return 0, false
	}
	i := len(*s) - 1
	j := len(*s) - 2

	x := (*s)[i] * (*s)[j]

	return x, true
}

func (s *Stack) SDiv() (uint8, bool) {
	if s.IsEmpty() || len(*s) < 2 {
		return 0, false
	}
	i := len(*s) - 1
	j := len(*s) - 2

	x := (*s)[i] / (*s)[j]

	return x, true
}

func (s *Stack) Exp() (uint8, uint8, bool) {
	if s.IsEmpty() || len(*s) < 2 {
		return 0, 0, false
	}
	i := len(*s) - 1
	j := len(*s) - 2

	x := uint8(math.Pow(float64((*s)[i]), float64((*s)[j])))

	return x, (*s)[j], true
}
