// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/min-stack/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type MinStack struct {
    
}

func Constructor() MinStack {

	return MinStack{}
}

func (m *MinStack) Push(val int)  {
    
}

func (m *MinStack) Pop()  {
    
}

func (m *MinStack) Top() (ans int) {

	return
}

func (m *MinStack) GetMin() (ans int) {

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
		case "push":
			methodParams := MustSplitArray(params[i])
			val := Deserialize[int](methodParams[0])
			obj.Push(val)
			output = append(output, "null")
		case "pop":
			obj.Pop()
			output = append(output, "null")
		case "top":
			ans := Serialize(obj.Top())
			output = append(output, ans)
		case "getMin":
			ans := Serialize(obj.GetMin())
			output = append(output, ans)
		}
	}
	fmt.Println("\noutput:", JoinArray(output))
}
