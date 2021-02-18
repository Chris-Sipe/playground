package main

import "testing"

func TestIsEmpty(t *testing.T) {
	s := Stack{
		arr: []int{1, 2, 3},
	}
	if s.isEmpty() {
		t.Errorf("Stack should not be empty. Stack: %d", s.arr)
	}
	s = Stack{
		arr: []int{},
	}
	if !s.isEmpty() {
		t.Errorf("Stack should be empty. Stack: %d", s.arr)
	}
}
func TestPop(t *testing.T) {
	s := Stack{
		arr: []int{1, 2, 3},
	}
	p, err := s.pop()
	if err != nil {
		t.Errorf("Pop error should be nil.")
	}
	if p != 3 {
		t.Errorf("Pop value is incorrect, got: %d, want: %d.", p, 3)
	}
	l := len(s.arr)
	if l != 2 {
		t.Errorf("Array is incorrect length after Pop, got: %d, want: %d.", l, 2)
	}
}

func TestPush(t *testing.T) {
	s := Stack{
		arr: []int{1, 2, 3},
	}
	s.push(4)
	l := len(s.arr)
	lastValue := s.arr[l-1]
	if l != 4 {
		t.Errorf("Array is incorrect length after Push, got: %d, want: %d.", l, 4)
	}
	if lastValue != 4 {
		t.Errorf("Last value in array is incorrect after Push, got: %d, want: %d.", lastValue, 4)
	}
}

func TestPeek(t *testing.T) {
	s := Stack{
		arr: []int{1, 2, 3},
	}
	p, err := s.peek()
	if err != nil {
		t.Errorf("Peek error should be nil.")
	}
	if p != 3 {
		t.Errorf("Peek value is incorrect, got: %d, want: %d.", p, 3)
	}
	l := len(s.arr)
	if l != 3 {
		t.Errorf("Array is incorrect length after Peek, got: %d, want: %d.", l, 3)
	}
}
