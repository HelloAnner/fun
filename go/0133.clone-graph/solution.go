// Created by anner at 2025/09/27 08:27
// leetgo: 1.4.15
// https://leetcode.cn/problems/clone-graph/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func cloneGraph(node *Node) (ans *Node) {

	return
}

// @lc code=end

// Warning: this is a manual question, the generated test code may be incorrect.
func main() {
	stdin := bufio.NewReader(os.Stdin)
	edges := Deserialize[[][]int](ReadLine(stdin))
	ans := cloneGraph(edges)

	fmt.Println("\noutput:", Serialize(ans))
}
