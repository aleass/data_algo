package link_list

import "errors"

type DoubleLinkedList struct {
	Id   int
	Next *DoubleLinkedList
	Prev *DoubleLinkedList
	Data string
}

//NewDoubleLinkedList return a nil head of DoubleLinkedList
func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

func (d *DoubleLinkedList) Push(id int, v *DoubleLinkedList) error {
	if id == 0 {
		return errors.New("invalid id")
	}
	temp := d
	for temp.Id != id {
		temp = temp.Next
	}
	if temp.Id != id {
		return errors.New("invalid id")
	}

	temp.Data = v.Data
	return nil
}

func (d *DoubleLinkedList) Append(v *DoubleLinkedList) {
	temp := d
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = v
	v.Id = temp.Id + 1
	v.Prev = temp
	d.Id = v.Id
}

func (d *DoubleLinkedList) DelId(id int) error {
	if id == 0 {
		return errors.New("invalid id")
	}
	temp := d
	for temp.Id != id {
		temp = temp.Next
	}
	if temp.Id != id {
		return errors.New("invalid id")
	}

	temp.Prev.Next = temp.Next
	if temp.Next != nil {
		temp.Next.Prev = temp.Prev
	}
	return nil
}
