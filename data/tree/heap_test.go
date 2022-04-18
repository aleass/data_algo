package tree

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	h := NewHeap()
	h.Insert(33)
	h.Insert(16)
	h.Insert(13)
	h.Insert(17)
	h.Insert(21)
	h.Insert(15)
	h.Insert(9)
	h.Insert(6)
	h.Insert(8)
	h.Insert(1)
	h.Insert(7)
	h.Insert(2)
	h.Insert(5)
	fmt.Println(h.data)
	h.Insert(22)

	fmt.Println(h.data)
}
