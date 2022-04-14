package queue

import (
	"sync"
	"testing"
)

//link list
/*
time.Sleep(22500)     	8740	    124582 ns/op
sync.mutex 				6550	    186217 ns/op
*/

func LinkedLIst() {
	var s sync.WaitGroup
	aq := NewLinkedList()
	var i int64 = 0
	s.Add(10)
	for ; i < 10; i++ {
		go func(i int64) {
			var _i = i * 100
			for ; _i < (i+1)*100; _i++ {
				var j = _i
				aq.Enqueue(&j)
			}
			s.Done()
		}(i)
	}

	s.Wait()
	//Close the following code when benchmarking
	//println(aq.Size)
	//var d = aq.Dequeue()
	//for d != nil {
	//	println(*d, aq.Size)
	//	d = aq.Dequeue()
	//}
}
func TestLinkedListQueue(t *testing.T) {
	LinkedLIst()
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LinkedLIst()
	}
}

// array list
/*
 18306	     66052 ns/op
*/
func ArrayLIst() {
	aq := NewArrayQueue()
	var s sync.WaitGroup

	var i int64 = 0
	s.Add(10)
	for ; i < 10; i++ {
		go func(i int64) {
			var _i = i * 100
			for ; _i < (i+1)*100; _i++ {
				var j = _i
				if err := aq.Enqueue(&j); err != nil {
					println(err.Error())
					break
				}
			}
			s.Done()
		}(i)
	}
	s.Wait()

	//var is = 0
	//data, err := aq.Dequeue()
	//for err == nil {
	//	println(*data)
	//	is++
	//	data, err = aq.Dequeue()
	//}
}

func TestArrayLIst(t *testing.T) {
	ArrayLIst()
}

func BenchmarkArrayLIst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ArrayLIst()
	}
}
