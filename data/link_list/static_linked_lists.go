package link_list

const maxArr int = 10

type staticLinkedLists struct {
	data  [maxArr]*node
	index int
}

type node struct {
	data string
	next int
}

func NewStaticLinked() staticLinkedLists {
	return staticLinkedLists{index: 1}
}

func (s *staticLinkedLists) Append(n *node) {
	if s.data[s.index] == nil {
		s.data[s.index] = n
		s.data[s.index].next = s.index + 1
		return
	}
	k := s.data[s.index].next
	for s.data[k] != nil {
		k = s.data[k].next
	}

}
