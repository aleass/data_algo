package sort

import (
	"fmt"
	"testing"
)

func TestBubble(t *testing.T) {
	var data = []int{15, 2, 4, 8, 9, 5, 10, 7, 12, 36}
	//fmt.Println(bubbleSort([]int{15, 2, 4, 8, 9, 5, 10, 7, 12, 36}))
	//insertSort(data)
	//selectionSort(data)
	//shellSort(data)
	//quickAscendingSort(data, 0, len(data)-1)
	//CountingSort(data, 36)
	//merge(data, []int{1, 8, 6, 2, 58, 2, 8, 2, 5})
	data = MergeSort(data)
	fmt.Println(data)
}
