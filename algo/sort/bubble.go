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

//CountingSort 计数排序
func CountingSort(arr []int, MaxVal int) {
	var assistArr = make([]int, MaxVal+1)
	for _, v := range arr {
		assistArr[v] += 1
	}
	var key int
	for k, v := range assistArr {
		for i := 0; i < v; i++ {
			arr[key] = k
			key++
		}
	}
}
func MergeSort(array []int) []int {
	n := len(array)
	if n < 2 {
		return array
	}
	key := n / 2
	left := MergeSort(array[0:key])
	right := MergeSort(array[key:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	tmp := make([]int, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			tmp = append(tmp, left[i])
			i++
		} else {
			tmp = append(tmp, right[j])
			j++
		}
	}
	tmp = append(tmp, left[i:]...)
	tmp = append(tmp, right[j:]...)
	return tmp
}
