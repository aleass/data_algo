package queue

import (
	"errors"
	"sync/atomic"
	"time"
)

const (
	NilErr  = "data is nil"
	FullErr = "data is full"
)

type ArrayQueue struct {
	data [1000]*int64
	head int64
	tail int64
}

func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{}
}

//Enqueue output data
func (q *ArrayQueue) Enqueue(val *int64) error {
	if q.tail-q.head >= int64(len(q.data)) {
		return errors.New(FullErr)
	}
top:
	t := q.tail
	for !atomic.CompareAndSwapInt64(&q.tail, t, t+1) {
		time.Sleep(22000)
		goto top
	}
	q.data[t%int64(len(q.data))] = val
	return nil
}

// Dequeue out data
func (q *ArrayQueue) Dequeue() (*int64, error) {
	if q.head == q.tail {
		return nil, errors.New(NilErr)
	}
	t := q.head
	atomic.AddInt64(&q.head, 1)
	return q.data[t%int64(len(q.data))], nil
}
