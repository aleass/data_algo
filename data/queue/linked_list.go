package queue

import "errors"

type Node struct {
	value interface{}
	next  *Node //next node
	pre   *Node //pre node
}

type LinkedListQueue struct {
	head  *Node
	size  int64
	count int64
}

func NewLinkedList(l int64) *LinkedListQueue {
	return &LinkedListQueue{
		size: l,
	}
}

func (q *LinkedListQueue) Enqueue(value interface{}) {
	if q.count == q.size {
		return
	}
	n := &Node{
		value: value,
	}
	q.count++
	if q.head == nil {
		q.head = n
		q.head.pre = n
		return
	}
	n.pre = q.head.pre
	q.head.pre.next = n
	q.head.pre = n
	return
}

func (q *LinkedListQueue) Dequeue() (interface{}, error) {
	if q.count == 0 {
		return nil, errors.New(NilErr)
	}
	val := q.head
	q.head = val.next
	q.count--
	return val, nil
}
