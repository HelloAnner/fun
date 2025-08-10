package algorithms

import (
	"math"
	"sort"
)

/**
 * 二分查找相关算法实现
 * Binary Search related algorithms
 */

/**
 * 704. 二分查找
 * Binary Search
 */
func Search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

/**
 * 69. x 的平方根
 * Sqrt(x)
 */
func MySqrt(x int) int {
	if x == 0 {
		return 0
	}

	left, right := 1, x

	for left <= right {
		mid := (left + right) / 2
		square := mid * mid

		if square == x {
			return mid
		} else if square < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}

/**
 * 35. 搜索插入位置
 * Search Insert Position
 */
func SearchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

/**
 * 34. 在排序数组中查找元素的第一个和最后一个位置
 * Find First and Last Position of Element in Sorted Array
 */
func SearchRange(nums []int, target int) []int {
	findFirst := func(nums []int, target int) int {
		left, right := 0, len(nums)-1
		result := -1

		for left <= right {
			mid := (left + right) / 2

			if nums[mid] == target {
				result = mid
				right = mid - 1 // 继续向左查找
			} else if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		return result
	}

	findLast := func(nums []int, target int) int {
		left, right := 0, len(nums)-1
		result := -1

		for left <= right {
			mid := (left + right) / 2

			if nums[mid] == target {
				result = mid
				left = mid + 1 // 继续向右查找
			} else if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		return result
	}

	return []int{findFirst(nums, target), findLast(nums, target)}
}

/**
 * 3143. 正方形中的最多点数
 * Maximum Points Inside the Square
 */
func MaxPointsInsideSquare(points [][]int, s string) int {
	n := len(points)
	distances := make([]DistChar, n)
	for i := 0; i < n; i++ {
		x, y := points[i][0], points[i][1]
		dist := int(math.Max(math.Abs(float64(x)), math.Abs(float64(y))))
		distances[i] = DistChar{Dist: dist, Char: s[i]}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Dist < distances[j].Dist
	})

	left, right := 0, distances[n-1].Dist
	result := 0

	for left <= right {
		mid := (left + right) / 2

		if isValidSquareSize(distances, mid) {
			result = countPointsInSquare(distances, mid)
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

type DistChar struct {
	Dist int
	Char byte
}

func isValidSquareSize(distances []DistChar, size int) bool {
	seen := make(map[byte]bool)

	for _, dc := range distances {
		if dc.Dist <= size {
			if seen[dc.Char] {
				return false
			}
			seen[dc.Char] = true
		}
	}

	return true
}

func countPointsInSquare(distances []DistChar, size int) int {
	seen := make(map[byte]bool)

	for _, dc := range distances {
		if dc.Dist <= size {
			seen[dc.Char] = true
		}
	}

	return len(seen)
}