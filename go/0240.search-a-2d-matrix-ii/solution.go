// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/search-a-2d-matrix-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func searchMatrix(matrix [][]int, target int) bool {
    
}

// @lc code=end

// Warning: this is a manual question, the generated test code may be incorrect.
func main() {
	stdin := bufio.NewReader(os.Stdin)
	matrix := Deserialize[[][]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := searchMatrix(matrix, target)

	fmt.Println("\noutput:", Serialize(ans))
}
