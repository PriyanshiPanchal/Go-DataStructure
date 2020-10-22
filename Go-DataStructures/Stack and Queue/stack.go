package main

import "fmt"

type Stack struct {
	items []int
}


func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() int {
	lastIndex := len(s.items) - 1
	toRemove := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return toRemove
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(35)
	myStack.Push(23)
	myStack.Push(78)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
}