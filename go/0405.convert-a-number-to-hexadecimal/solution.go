// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/convert-a-number-to-hexadecimal/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	hex := "0123456789abcdef"
	result := ""

	// 处理负数
	var n uint32
	if num < 0 {
		n = uint32(num)
	} else {
		n = uint32(num)
	}

	for n > 0 {
		result = string(hex[n&15]) + result
		n >>= 4
	}

	return result
}

@lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	num := Deserialize[int](ReadLine(stdin))
	ans := toHex(num)

	fmt.Println("\noutput:", Serialize(ans))
}
