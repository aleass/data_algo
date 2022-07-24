package queue

import (
	"fmt"
	"testing"
)

func TestLIstQueue(t *testing.T) {
	q := NewLinkedList(5)
	for i := 0; i < 100; i++ {
		q.Enqueue(i)
	}
	for data, err := q.Dequeue(); err == nil; data, err = q.Dequeue() {
		fmt.Println(data)
	}
	for i := 30; i < 100; i++ {
		q.Enqueue(i)
	}
	for data, err := q.Dequeue(); err == nil; data, err = q.Dequeue() {
		fmt.Println(data)
	}
	for i := 10; i < 100; i++ {
		q.Enqueue(i)
	}
	for data, err := q.Dequeue(); err == nil; data, err = q.Dequeue() {
		fmt.Println(data)
	}
}

func TestArrayQueue(t *testing.T) {
	q := NewArrayQueue(5)
	for i := 0; i < 100; i++ {
		q.Enqueue(i)
	}
	for data, err := q.Dequeue(); err == nil; data, err = q.Dequeue() {
		fmt.Println(data)
	}
	for i := 30; i < 100; i++ {
		q.Enqueue(i)
	}
	for data, err := q.Dequeue(); err == nil; data, err = q.Dequeue() {
		fmt.Println(data)
	}
	for i := 10; i < 100; i++ {
		q.Enqueue(i)
	}
	for data, err := q.Dequeue(); err == nil; data, err = q.Dequeue() {
		fmt.Println(data)
	}
}
