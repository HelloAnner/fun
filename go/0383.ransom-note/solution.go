// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/ransom-note/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func canConstruct(ransomNote string, magazine string) bool {
    
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	ransomNote := Deserialize[string](ReadLine(stdin))
	magazine := Deserialize[string](ReadLine(stdin))
	ans := canConstruct(ransomNote, magazine)

	fmt.Println("\noutput:", Serialize(ans))
}
