package algorithms

/**
 * 排序相关算法实现
 * Sorting related algorithms
 */

/**
 * 快速排序
 * Quick Sort
 */
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	var left, middle, right []int

	for _, x := range arr {
		if x < pivot {
			left = append(left, x)
		} else if x == pivot {
			middle = append(middle, x)
		} else {
			right = append(right, x)
		}
	}

	result := QuickSort(left)
	result = append(result, middle...)
	result = append(result, QuickSort(right)...)
	return result
}

/**
 * 归并排序
 * Merge Sort
 */
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return mergeSortHelper(left, right)
}

func mergeSortHelper(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

/**
 * 堆排序
 * Heap Sort
 */
func HeapSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	n := len(result)

	// 构建最大堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(result, n, i)
	}

	// 逐个提取元素
	for i := n - 1; i > 0; i-- {
		result[0], result[i] = result[i], result[0]
		heapify(result, i, 0)
	}

	return result
}

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

/**
 * 冒泡排序
 * Bubble Sort
 */
func BubbleSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	n := len(result)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}

	return result
}

/**
 * 插入排序
 * Insertion Sort
 */
func InsertionSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)

	for i := 1; i < len(result); i++ {
		key := result[i]
		j := i - 1

		for j >= 0 && result[j] > key {
			result[j+1] = result[j]
			j--
		}

		result[j+1] = key
	}

	return result
}

/**
 * 选择排序
 * Selection Sort
 */
func SelectionSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	n := len(result)

	for i := 0; i < n-1; i++ {
		minIdx := i

		for j := i + 1; j < n; j++ {
			if result[j] < result[minIdx] {
				minIdx = j
			}
		}

		if minIdx != i {
			result[i], result[minIdx] = result[minIdx], result[i]
		}
	}

	return result
}

/**
 * 计数排序
 * Counting Sort
 */
func CountingSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	max := arr[0]
	min := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	rangeSize := max - min + 1
	count := make([]int, rangeSize)
	result := make([]int, len(arr))

	// 计数
	for _, num := range arr {
		count[num-min]++
	}

	// 累积计数
	for i := 1; i < rangeSize; i++ {
		count[i] += count[i-1]
	}

	// 构建结果数组
	for i := len(arr) - 1; i >= 0; i-- {
		result[count[arr[i]-min]-1] = arr[i]
		count[arr[i]-min]--
	}

	return result
}
