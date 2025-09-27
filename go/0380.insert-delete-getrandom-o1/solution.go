// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/insert-delete-getrandom-o1/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type RandomizedSet struct {
    
}

func Constructor() RandomizedSet {

	return RandomizedSet{}
}

func (r *RandomizedSet) Insert(val int) bool {
    
}

func (r *RandomizedSet) Remove(val int) bool {
    
}

func (r *RandomizedSet) GetRandom() (ans int) {

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
		case "insert":
			methodParams := MustSplitArray(params[i])
			val := Deserialize[int](methodParams[0])
			ans := Serialize(obj.Insert(val))
			output = append(output, ans)
		case "remove":
			methodParams := MustSplitArray(params[i])
			val := Deserialize[int](methodParams[0])
			ans := Serialize(obj.Remove(val))
			output = append(output, ans)
		case "getRandom":
			ans := Serialize(obj.GetRandom())
			output = append(output, ans)
		}
	}
	fmt.Println("\noutput:", JoinArray(output))
}
