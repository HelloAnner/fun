// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/valid-sudoku/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val := board[i][j]
			if val == '.' {
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

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	board := Deserialize[[][]byte](ReadLine(stdin))
	ans := isValidSudoku(board)

	fmt.Println("\noutput:", Serialize(ans))
}
