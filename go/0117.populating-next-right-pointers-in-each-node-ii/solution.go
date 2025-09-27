// Created by anner at 2025/09/27 08:54
// leetgo: 1.4.15
// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func connect(root *Node) (ans *Node) {

	return
}

// @lc code=end

// Warning: this is a manual question, the generated test code may be incorrect.
func main() {
	stdin := bufio.NewReader(os.Stdin)
	root := Deserialize[*TreeNode](ReadLine(stdin))
	ans := connect(root)

	fmt.Println("\noutput:", Serialize(ans))
}
