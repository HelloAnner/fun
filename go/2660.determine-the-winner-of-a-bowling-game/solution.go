// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/determine-the-winner-of-a-bowling-game/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isWinner(player1 []int, player2 []int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	player1 := Deserialize[[]int](ReadLine(stdin))
	player2 := Deserialize[[]int](ReadLine(stdin))
	ans := isWinner(player1, player2)

	fmt.Println("\noutput:", Serialize(ans))
}
