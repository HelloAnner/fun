// Created by anner at 2025/09/27 08:56
// leetgo: 1.4.15
// https://leetcode.cn/problems/design-spreadsheet/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type Spreadsheet struct {
    
}

func Constructor(rows int) Spreadsheet {

	return Spreadsheet{}
}

func (s *Spreadsheet) SetCell(cell string, value int)  {
    
}

func (s *Spreadsheet) ResetCell(cell string)  {
    
}

func (s *Spreadsheet) GetValue(formula string) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	ops := Deserialize[[]string](ReadLine(stdin))
	params := MustSplitArray(ReadLine(stdin))
	output := make([]string, 0, len(ops))
	output = append(output, "null")

	constructorParams := MustSplitArray(params[0])
	rows := Deserialize[int](constructorParams[0])
	obj := Constructor(rows)

	for i := 1; i < len(ops); i++ {
		switch ops[i] {
		case "setCell":
			methodParams := MustSplitArray(params[i])
			cell := Deserialize[string](methodParams[0])
			value := Deserialize[int](methodParams[1])
			obj.SetCell(cell, value)
			output = append(output, "null")
		case "resetCell":
			methodParams := MustSplitArray(params[i])
			cell := Deserialize[string](methodParams[0])
			obj.ResetCell(cell)
			output = append(output, "null")
		case "getValue":
			methodParams := MustSplitArray(params[i])
			formula := Deserialize[string](methodParams[0])
			ans := Serialize(obj.GetValue(formula))
			output = append(output, ans)
		}
	}
	fmt.Println("\noutput:", JoinArray(output))
}
