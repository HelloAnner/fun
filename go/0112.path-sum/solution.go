// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/path-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	return HasPathSum(root.Left, targetSum-root.Val) ||
		HasPathSum(root.Right, targetSum-root.Val)
}

@lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	targetSum := Deserialize[int](ReadLine(stdin))
	ans := hasPathSum(root, targetSum)

	fmt.Println("\noutput:", Serialize(ans))
}
