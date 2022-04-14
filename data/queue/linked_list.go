package queue

import (
	"sync/atomic"
	"time"
)

type Node struct {
	value *int64
	next  *Node
}

type LinkedListQueue struct {
	head *Node
	tail *Node
	Size int64
	lock int64 //atomic instead of sync.Mutex
	//lock sync.Mutex //sync.Mutex
}

func NewLinkedList() *LinkedListQueue {
	return &LinkedListQueue{}
}

func (q *LinkedListQueue) Enqueue(value *int64) {
	node := &Node{value: value}
	//q.lock.Lock()
	//defer q.lock.Unlock() //lock : close top:code
top:
	if !atomic.CompareAndSwapInt64(&q.lock, 0, 1) {
		time.Sleep(22500)
		goto top
	}
	goto end
end:
	if q.head == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.Size++
	if !atomic.CompareAndSwapInt64(&q.lock, 1, 0) {
		panic("unlock failed")
	}
}

func (q *LinkedListQueue) Dequeue() *int64 {
	if q.head == nil {
		return nil
	}
	value := q.head.value
	q.head = q.head.next
	q.Size--
	return value
}
