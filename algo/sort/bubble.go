package sort

// BubbleSort 泡沫排序
func bubbleSort(data []int) []int {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[j] > data[i] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	return data
}

// insertSort 插入排序
func insertSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0 && arr[j] > arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

// selectionSort 插入排序
func selectionSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}
	for i := 0; i < n; i++ {
		var k = i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[k] {
				k = j
			}
		}
		if k != i {
			arr[i], arr[k] = arr[k], arr[i]
		}
	}
}

func shellSort(arr []int) {
	n := len(arr)
	h := 1
	//寻找合适的间隔h
	for h < n/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h && arr[j] > arr[j-1]; j -= h {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
		h /= 3
	}
}

func quickAscendingSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		if start < j {
			quickAscendingSort(arr, start, j)
		}
		if end > i {
			quickAscendingSort(arr, i, end)
		}
	}
}
