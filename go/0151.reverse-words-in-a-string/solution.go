// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/reverse-words-in-a-string/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
	"strings"
)

// @lc code=begin

func reverseWords(s string) string {
	// 去除首尾空格并按空格分割
	words := strings.Fields(s)

	// 反转单词数组
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

@lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := reverseWords(s)

	fmt.Println("\noutput:", Serialize(ans))
}
