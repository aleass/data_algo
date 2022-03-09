package linked_list

type SinglyLinkedList struct {
	Id      int
	NextPor *SinglyLinkedList
	Data    string
}

// NewSingly return a nil head of SinglyLinkedList
func NewSingly() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

// ReadSingly return index equal id of value
func (s *SinglyLinkedList) ReadSingly(id int) (*SinglyLinkedList, error) {
	if s.NextPor == nil {
		return nil, nil
	}
	var p = s.NextPor
	for p.NextPor != nil && p.Id != id {
		p = p.NextPor
	}
	if p.Id != id {
		p = nil
	}
	return p, nil
}

// AddSingly add a SinglyLinkedList from index id of next
func (s *SinglyLinkedList) AddSingly(id int, val *SinglyLinkedList) error {
	if s.NextPor == nil {
		s.NextPor = val
		return nil
	}
	var p = s.NextPor
	for p.NextPor != nil && p.Id != id {
		p = p.NextPor
	}
	if p.Id == id {
		val.NextPor = p.NextPor
	}
	p.NextPor = val
	return nil
}

func (s *SinglyLinkedList) DeleteOrChangeSingly(id int, val *SinglyLinkedList) error {
	if s.NextPor == nil {
		return nil
	}
	var p = s.NextPor
	var pre = s
	for p.NextPor != nil && p.Id != id {
		pre = p
		p = p.NextPor
	}
	if p.Id == id {
		if val != nil {
			val.NextPor = p.NextPor
			p.NextPor = val
		}
		pre.NextPor = p.NextPor
	}

	return nil
}
