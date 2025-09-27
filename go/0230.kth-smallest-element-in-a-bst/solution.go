// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/kth-smallest-element-in-a-bst/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func kthSmallest(root *TreeNode, k int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := kthSmallest(root, k)

	fmt.Println("\noutput:", Serialize(ans))
}
