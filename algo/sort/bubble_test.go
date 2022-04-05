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
	quickAscendingSort(data, 0, len(data)-1)
	fmt.Println(data)
}
