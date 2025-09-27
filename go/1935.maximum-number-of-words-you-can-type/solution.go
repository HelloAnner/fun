// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/maximum-number-of-words-you-can-type/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func canBeTypedWords(text string, brokenLetters string) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	text := Deserialize[string](ReadLine(stdin))
	brokenLetters := Deserialize[string](ReadLine(stdin))
	ans := canBeTypedWords(text, brokenLetters)

	fmt.Println("\noutput:", Serialize(ans))
}
