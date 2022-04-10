package link_list

import (
	"fmt"
	"testing"
)

// singly linked list
func TestSinglyLinkedList(t *testing.T) {
	var link = NewSingly()
	link.AddSingly(0, &SinglyLinkedList{
		Id:   1,
		Data: "i am singly 1",
	})
	link.AddSingly(20, &SinglyLinkedList{
		Id:   21,
		Data: "i am singly 50",
	})
	link.AddSingly(1, &SinglyLinkedList{
		Id:   31,
		Data: "i am singly 31",
	})
	res, _ := link.ReadSingly(31)
	fmt.Println(res)
	link.DeleteOrChangeSingly(31, &SinglyLinkedList{
		Id:   99,
		Data: "i am singly 99",
	})
}

// singly linked list
func TestDoubleLinkedList(t *testing.T) {
	var link = NewDoubleLinkedList()
	link.Append(&DoubleLinkedList{
		Data: "data 1",
	})
	link.Append(&DoubleLinkedList{
		Data: "data 50",
	})
	link.Append(&DoubleLinkedList{
		Data: "data 31",
	})
	link.Push(1, &DoubleLinkedList{
		Data: "data 311111",
	})

	link.DelId(1)

	fmt.Println(link)
}
