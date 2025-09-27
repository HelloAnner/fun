// Created by anner at 2025/09/27 08:54
// leetgo: 1.4.15
// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func buildTree(preorder []int, inorder []int) (ans *TreeNode) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	preorder := Deserialize[[]int](ReadLine(stdin))
	inorder := Deserialize[[]int](ReadLine(stdin))
	ans := buildTree(preorder, inorder)

	fmt.Println("\noutput:", Serialize(ans))
}
