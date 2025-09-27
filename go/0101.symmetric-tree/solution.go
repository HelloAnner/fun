// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/symmetric-tree/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSameTree(root.Left, root.Right)
}

@lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := isSymmetric(root)

	fmt.Println("\noutput:", Serialize(ans))
}
