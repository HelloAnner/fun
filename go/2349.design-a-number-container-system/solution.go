// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/design-a-number-container-system/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type NumberContainers struct {
    
}

func Constructor() NumberContainers {

	return NumberContainers{}
}

func (n *NumberContainers) Change(index int, number int)  {
    
}

func (n *NumberContainers) Find(number int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	ops := Deserialize[[]string](ReadLine(stdin))
	params := MustSplitArray(ReadLine(stdin))
	output := make([]string, 0, len(ops))
	output = append(output, "null")

	obj := Constructor()

	for i := 1; i < len(ops); i++ {
		switch ops[i] {
		case "change":
			methodParams := MustSplitArray(params[i])
			index := Deserialize[int](methodParams[0])
			number := Deserialize[int](methodParams[1])
			obj.Change(index, number)
			output = append(output, "null")
		case "find":
			methodParams := MustSplitArray(params[i])
			number := Deserialize[int](methodParams[0])
			ans := Serialize(obj.Find(number))
			output = append(output, ans)
		}
	}
	fmt.Println("\noutput:", JoinArray(output))
}
