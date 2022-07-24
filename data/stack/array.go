package stack

import "errors"

type arrayStack struct {
	len  int
	curr int
	head []interface{}
}

func NewArrayStack(l int) *arrayStack {
	return &arrayStack{
		len:  l,
		head: make([]interface{}, l),
	}
}

func (s *arrayStack) PushData(data interface{}) (int, error) {
	if s.curr == s.len {
		return 0, errors.New("stack is full")
	}

	s.head[s.curr] = data
	s.curr++
	return s.curr, nil
}

func (s *arrayStack) PullData() (interface{}, int, error) {
	if s.curr == 0 {
		return nil, 0, errors.New("stack is empty")
	}
	s.curr--
	st := s.head[s.curr]
	return st, s.curr, nil
}
