package algorithms

import "fmt"

// MajorityElement 找到数组中的多数元素（出现次数大于 n/2 的元素）
// 使用排序方法：多数元素排序后一定在中间位置
func MajorityElement(nums []int) int {
	// 创建副本进行排序
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	
	// 简单排序
	for i := 0; i < len(sorted); i++ {
		for j := i + 1; j < len(sorted); j++ {
			if sorted[i] > sorted[j] {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}
	
	return sorted[len(sorted)/2]
}

// MajorityElementByCount 使用计数方法找到多数元素
func MajorityElementByCount(nums []int) int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}
	
	maxCount := 0
	result := nums[0]
	for num, count := range counts {
		if count > maxCount {
			maxCount = count
			result = num
		}
	}
	
	return result
}

// RotateArray 将数组向右轮转 k 个位置
func RotateArray(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return nums
	}
	
	k = k % n // 处理 k 大于数组长度的情况
	result := make([]int, n)
	
	// 将后 k 个元素移到前面
	for i := 0; i < k; i++ {
		result[i] = nums[n-k+i]
	}
	
	// 将前 n-k 个元素移到后面
	for i := 0; i < n-k; i++ {
		result[k+i] = nums[i]
	}
	
	return result
}

// RotateArrayInPlace 原地轮转数组
func RotateArrayInPlace(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	
	k = k % n
	if k == 0 {
		return
	}
	
	// 使用三次反转的方法
	reverse(nums, 0, n-1)     // 反转整个数组
	reverse(nums, 0, k-1)     // 反转前 k 个元素
	reverse(nums, k, n-1)     // 反转后 n-k 个元素
}

// reverse 反转数组的指定部分
func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// DemonstrateArrayAlgorithms 演示数组算法
func DemonstrateArrayAlgorithms() {
	fmt.Println("=== 数组算法演示 ===")
	
	// 多数元素
	nums := []int{3, 2, 3}
	majority := MajorityElement(nums)
	fmt.Printf("数组: %v, 多数元素: %d\n", nums, majority)
	
	// 轮转数组
	nums2 := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	fmt.Printf("原数组: %v\n", nums2)
	RotateArray(nums2, k)
	fmt.Printf("轮转 %d 位后: %v\n", k, nums2)
	
	// 旋转矩阵
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("原矩阵:")
	printMatrix(matrix)
	RotateMatrix(matrix)
	fmt.Println("旋转90度后:")
	printMatrix(matrix)
	fmt.Println()
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Printf("%v\n", row)
	}
}

// RotateMatrix 顺时针旋转 90 度矩阵
func RotateMatrix(matrix [][]int) {
	n := len(matrix)
	if n == 0 {
		return
	}
	
	// 创建临时矩阵
	tmp := make([][]int, n)
	for i := range tmp {
		tmp[i] = make([]int, n)
		copy(tmp[i], matrix[i])
	}
	
	// 根据旋转公式：matrix[j][n-1-i] = tmp[i][j]
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[j][n-1-i] = tmp[i][j]
		}
	}
}

// RotateMatrixInPlace 原地旋转矩阵
func RotateMatrixInPlace(matrix [][]int) {
	n := len(matrix)
	if n == 0 {
		return
	}
	
	// 先转置矩阵
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	
	// 再水平翻转每一行
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}