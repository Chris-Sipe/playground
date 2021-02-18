package main

import "fmt"

func Sum(x int, y int) int {
	return x + y
}

func main() {
	Sum(5, 5)
}

type Stack struct {
	arr []int
}

func (s *Stack) isEmpty() bool {
	return len(s.arr) == 0
}

func (s *Stack) pop() (int, error) {
	if s.isEmpty() {
		return 0, fmt.Errorf("tried to pop, but stack is empty")
	}
	l := len(s.arr)
	top := s.arr[l-1]
	s.arr = s.arr[:l-1]
	return top, nil
}

func (s *Stack) peek() (int, error) {
	l := len(s.arr)
	if l <= 0 {
		return 0, fmt.Errorf("tried to peek, but stack is empty")
	}
	top := s.arr[l-1]
	return top, nil
}

func (s *Stack) push(n int) {
	s.arr = append(s.arr, n)
}
