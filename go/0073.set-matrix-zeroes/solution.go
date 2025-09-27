// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/set-matrix-zeroes/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func setZeroes(matrix [][]int)  {
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

@lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	matrix := Deserialize[[][]int](ReadLine(stdin))
	setZeroes(matrix)
	ans := matrix

	fmt.Println("\noutput:", Serialize(ans))
}
