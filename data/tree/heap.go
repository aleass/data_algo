package tree

type Heap struct {
	data []int
}

func NewHeap() *Heap {
	h := Heap{}
	h.data = make([]int, 1)
	return &h
}

func (h *Heap) Insert(val int) {
	if len(h.data) == 1 {
		h.data = append(h.data, val)
		return
	}
	h.data = append(h.data, val)
	var l = len(h.data) - 1
	var i = 0 //left 偶数  right 单数
	if l%2 == 1 {
		i = 1
	}
	var f = (l - i) / 2
	//
	for h.data[l] >= h.data[f] && l > 1 && f > 1 {
		h.data[l], h.data[f] = h.data[f], h.data[l]
		l = f
		f = (l - i) / 2
	}

}
