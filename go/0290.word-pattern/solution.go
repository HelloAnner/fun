// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/word-pattern/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
	"strings"
)

// @lc code=begin

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	charToWord := make(map[rune]string)
	wordToChar := make(map[string]rune)

	patternRunes := []rune(pattern)

	for i := 0; i < len(patternRunes); i++ {
		char := patternRunes[i]
		word := words[i]

		if mapped, exists := charToWord[char]; exists {
			if mapped != word {
				return false
			}
		} else {
			charToWord[char] = word
		}

		if mapped, exists := wordToChar[word]; exists {
			if mapped != char {
				return false
			}
		} else {
			wordToChar[word] = char
		}
	}

	return true
}

@lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	pattern := Deserialize[string](ReadLine(stdin))
	s := Deserialize[string](ReadLine(stdin))
	ans := wordPattern(pattern, s)

	fmt.Println("\noutput:", Serialize(ans))
}
