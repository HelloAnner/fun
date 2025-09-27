// Created by anner at 2025/09/27 08:27
// leetgo: 1.4.15
// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func lowestCommonAncestor(root, p, q *TreeNode) (ans *TreeNode) {

	return
}

// @lc code=end

// Warning: this is a manual question, the generated test code may be incorrect.
func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	p := Deserialize[int](ReadLine(stdin))
	q := Deserialize[int](ReadLine(stdin))
	ans := lowestCommonAncestor(root, p, q)

	fmt.Println("\noutput:", Serialize(ans))
}
