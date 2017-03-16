package parse

import (
	"errors"
	"fmt"
	"sync"
)

type Tree interface {
	Left() Tree
	Right() Tree
	Value() token
}

type tree struct {
	value token
	left  *tree
	right *tree
}

func (t *tree) Value() token {
	return t.value
}

func (t *tree) Left() Tree {
	return t.left
}

func (t *tree) Right() Tree {
	return t.right
}

func (t *tree) String() string {
	var left, right string
	if t.left != nil {
		left = t.left.String()
	}
	if t.right != nil {
		right = t.right.String()
	}
	return fmt.Sprintf("[%s %s %s]", left, t.value, right)
}

func newTree() *tree {
	return &tree{}
}

type stack struct {
	s []*tree
	sync.Mutex
}

func newStack() *stack {
	return &stack{
		s: make([]*tree, 0),
	}
}

func (s *stack) push(t *tree) {
	s.Lock()
	defer s.Unlock()

	s.s = append(s.s, t)
}

func (s *stack) pop() (*tree, error) {
	s.Lock()
	defer s.Unlock()

	l := len(s.s)
	if l == 0 {
		return nil, errors.New("empty stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]

	return res, nil
}
