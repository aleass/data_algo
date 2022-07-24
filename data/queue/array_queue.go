package queue

import (
	"errors"
)

const (
	NilErr  = "data is nil"
	FullErr = "data is full"
)

type ArrayQueue struct {
	data  []interface{}
	head  int64
	tail  int64
	len   int64 //  size of the circular queue
	count int64 // total data in the queue
}

func NewArrayQueue(l int64) *ArrayQueue {
	return &ArrayQueue{
		data: make([]interface{}, l),
		len:  l,
	}
}

//Enqueue input data
func (q *ArrayQueue) Enqueue(val interface{}) error {
	if q.count == q.len {
		return errors.New(FullErr)
	}
	q.data[q.tail] = val
	q.count++
	q.tail++
	if q.tail == q.len {
		q.tail = 0
	}
	return nil
}

// Dequeue out data
func (q *ArrayQueue) Dequeue() (interface{}, error) {
	if q.count == 0 {
		return nil, errors.New(NilErr)
	}
	var val = q.data[q.head]
	q.head++
	q.count--
	if q.head == q.len {
		q.head = 0
	}
	return val, nil
}
