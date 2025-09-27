// Created by anner at 2025/09/27 08:54
// leetgo: 1.4.15
// https://leetcode.cn/problems/length-of-last-word/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func lengthOfLastWord(s string) (ans int) {
	i := len(s) - 1
	// 跳过尾部空格
	for i >= 0 && s[i] == ' ' {
		i--
	}

	j := i - 1
	// 找到最后一个单词的开始位置
	for j >= 0 && s[j] != ' ' {
		j--
	}

	return i - j
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := lengthOfLastWord(s)

	fmt.Println("\noutput:", Serialize(ans))
}
