// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/number-of-islands/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func numIslands(grid [][]byte) (ans int) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	count := 0

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= rows || j < 0 || j >= cols || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0' // 标记为已访问
		dfs(i+1, j)      // 下
		dfs(i-1, j)      // 上
		dfs(i, j+1)      // 右
		dfs(i, j-1)      // 左
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j)
			}
		}
	}

	return count
}

// @lc code=end

// Warning: this is a manual question, the generated test code may be incorrect.
func main() {
	stdin := bufio.NewReader(os.Stdin)
	grid := Deserialize[[][]byte](ReadLine(stdin))
	ans := numIslands(grid)

	fmt.Println("\noutput:", Serialize(ans))
}
