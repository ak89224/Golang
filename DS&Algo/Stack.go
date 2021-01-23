package main

import (
	"errors"
	"fmt"
	"sync"
)

type stack struct {
	lock sync.Mutex // you don't have to do this if you don't want thread safety
	s []int
}

func NewStack() *stack {
	return &stack {sync.Mutex{}, make([]int,0) }
}

func (s *stack) Push(v int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

func (s *stack) Pop() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()


	l := len(s.s)
	if l == 0 {
		return 0, errors.New("Empty Stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}

func (s *stack) Print() {
	for i := range s.s {
		fmt.Print(s.s[len(s.s)-1-i], " ")
	}
	fmt.Println()
}


func main(){
	s := NewStack()
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Print()
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.s)
	s.Print()

}
