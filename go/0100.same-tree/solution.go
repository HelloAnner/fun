// Created by anner at 2025/09/27 08:54
// leetgo: 1.4.15
// https://leetcode.cn/problems/same-tree/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isSameTree(p *TreeNode, q *TreeNode) bool {
    
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	p := Deserialize[*TreeNode](ReadLine(stdin))
	q := Deserialize[*TreeNode](ReadLine(stdin))
	ans := isSameTree(p, q)

	fmt.Println("\noutput:", Serialize(ans))
}
