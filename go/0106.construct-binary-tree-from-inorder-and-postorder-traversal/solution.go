// Created by anner at 2025/09/27 08:54
// leetgo: 1.4.15
// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func buildTree(inorder []int, postorder []int) (ans *TreeNode) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	inorder := Deserialize[[]int](ReadLine(stdin))
	postorder := Deserialize[[]int](ReadLine(stdin))
	ans := buildTree(inorder, postorder)

	fmt.Println("\noutput:", Serialize(ans))
}
