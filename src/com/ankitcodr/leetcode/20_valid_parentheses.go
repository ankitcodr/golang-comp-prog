package leetcode

import (
	"sync"
	"errors"
)

type stack struct {
	lock sync.Mutex
	s    []byte
}

func NewStack() *stack {
	return &stack{
		lock: sync.Mutex{},
		s:    []byte{},
	}
}
func (s *stack) Push(in byte) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.s = append(s.s, in)
}

func (s *stack) Pop() (byte, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	l := len(s.s)
	if l == 0 {
		return 0, errors.New("Empty Stack")
	}
	last := s.s[l-1]
	s.s = s.s[:l-1]
	return last, nil
}

func (s *stack) IsEmpty() bool {
	if len(s.s) == 0 {
		return true
	}
	return false
}

func isValid(s string) bool {
	st := NewStack()
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			st.Push(s[i])
		case ')':
			if p, err := st.Pop(); err == nil && p != '(' {
				return false
			}
		case ']':
			if p, err := st.Pop(); err == nil && p != '[' {
				return false
			}
		case '}':
			if p, err := st.Pop(); err == nil && p != '{' {
				return false
			}
		}
	}
	return st.IsEmpty()
}

func IsValidSolution(s string) bool {
	return isValid(s)
}