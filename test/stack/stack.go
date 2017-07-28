package main

import "fmt"

type Sousa interface {
	Pop()
	Push()
}
type Stack struct {
	stk []string
}

func (s *Stack) Pop() string {
	pop := s.stk[len(s.stk)-1]
	s.stk = s.stk[:len(s.stk)-1]
	fmt.Println(pop)
	return pop
}
func (s *Stack) Push(str string) {
	s.stk = append(s.stk, str)
}

func main() {
	s := &Stack{}
	s.Push("dataA")
	s.Push("dataB")
	s.Push("dataC")
	s.Pop() // -> "dataC"
	s.Pop() // -> "dataB"
	s.Push("dataD")
	s.Pop() // -> "dataD"
	s.Pop() // -> "dataA"
	//s.Pop() // -> ""
}
