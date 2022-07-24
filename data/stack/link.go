package stack

type headStack struct {
	len  int
	head *stack
}

type stack struct {
	data interface{}
	pre  *stack
}

func NewLinkStack() *headStack {
	return &headStack{}
}

func (s *headStack) PushData(data interface{}) (int, error) {
	st := &stack{
		data: data,
	}
	if s.len == 0 {
		s.head = st
		s.len++
		return 1, nil
	}
	st.pre = s.head
	s.head = st
	s.len++
	return s.len, nil
}

func (s *headStack) PullData() (*stack, int, error) {
	if s.head == nil {
		return nil, 0, nil
	}
	st := s.head
	s.head = st.pre
	s.len--
	return st, s.len, nil
}
