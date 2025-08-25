package algorithms

import (
	"math"
	"sort"
)

/**
 * 模拟算法相关实现
 * Simulation algorithms
 */

/**
 * 36. 有效的数独
 * Valid Sudoku
 */
func IsValidSudoku(board [][]string) bool {
	rows := make([]map[string]bool, 9)
	cols := make([]map[string]bool, 9)
	boxes := make([]map[string]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[string]bool)
		cols[i] = make(map[string]bool)
		boxes[i] = make(map[string]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val := board[i][j]
			if val == "." {
				continue
			}

			boxIndex := (i/3)*3 + j/3

			if rows[i][val] || cols[j][val] || boxes[boxIndex][val] {
				return false
			}

			rows[i][val] = true
			cols[j][val] = true
			boxes[boxIndex][val] = true
		}
	}

	return true
}

/**
 * 48. 旋转图像
 * Rotate Image
 */
func Rotate(matrix [][]int) {
	n := len(matrix)

	// 转置矩阵
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 反转每一行
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}

/**
 * 54. 螺旋矩阵
 * Spiral Matrix
 */
func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	result := []int{}
	top := 0
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	for top <= bottom && left <= right {
		// 从左到右
		for j := left; j <= right; j++ {
			result = append(result, matrix[top][j])
		}
		top++

		// 从上到下
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// 从右到左
		if top <= bottom {
			for j := right; j >= left; j-- {
				result = append(result, matrix[bottom][j])
			}
			bottom--
		}

		// 从下到上
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}

/**
 * 73. 矩阵置零
 * Set Matrix Zeroes
 */
func SetZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	firstRowZero := false
	firstColZero := false

	// 检查第一行是否有零
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			firstRowZero = true
			break
		}
	}

	// 检查第一列是否有零
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstColZero = true
			break
		}
	}

	// 使用第一行和第一列作为标记
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 根据标记置零
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 处理第一行
	if firstRowZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	// 处理第一列
	if firstColZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

/**
 * 274. H 指数
 * H-Index
 */
func HIndex(citations []int) int {
	citationsCopy := make([]int, len(citations))
	copy(citationsCopy, citations)
	sort.Slice(citationsCopy, func(i, j int) bool {
		return citationsCopy[i] > citationsCopy[j]
	})

	h := 0
	for i := 0; i < len(citationsCopy); i++ {
		if citationsCopy[i] >= i+1 {
			h = i + 1
		} else {
			break
		}
	}

	return h
}

/**
 * 1103. 分糖果 II
 * Distribute Candies to People
 */
func DistributeCandies(candies, numPeople int) []int {
	result := make([]int, numPeople)
	give := 1
	person := 0

	for candies > 0 {
		toGive := give
		if toGive > candies {
			toGive = candies
		}
		result[person] += toGive
		candies -= toGive
		give++
		person = (person + 1) % numPeople
	}

	return result
}

/**
 * 3195 找到所有 Ones 的最小矩形面积 I
 * https://leetcode.cn/problems/find-the-minimum-area-to-cover-all-ones-i/?envType=daily-question&envId=2025-08-22
 */
func minimumArea(grid [][]int) int {
	// 寻找四个方向的边界位置
	left, right := math.MaxInt, 0
	top, bottom := math.MaxInt, 0
	for i, row := range grid {
		for j, x := range row {
			if x == 1 {
				left = min(left, j)
				right = max(right, j)
				top = min(top, i)
				bottom = max(bottom, i)
			}
		}
	}
	return (right - left + 1) * (bottom - top + 1)
}

// 498. 对角线遍历
// 给你一个大小为 m x n 的矩阵 mat ，请以对角线遍历的顺序，用一个数组返回这个矩阵中的所有元素。
// https://leetcode.cn/problems/diagonal-traverse/description/?envType=daily-question&envId=2025-08-25
func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	ans := make([]int, 0, m*n) // 预分配空间
	for k := range m + n - 1 {
		minJ := max(k-m+1, 0)
		maxJ := min(k, n-1)
		if k%2 == 0 { // 偶数从左到右
			for j := minJ; j <= maxJ; j++ {
				ans = append(ans, mat[k-j][j])
			}
		} else { // 奇数从右到左
			for j := maxJ; j >= minJ; j-- {
				ans = append(ans, mat[k-j][j])
			}
		}
	}
	return ans
}
