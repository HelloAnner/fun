package algorithms

import (
	"fmt"
	"strconv"
)

/**
 * 栈相关算法实现
 * Stack-related algorithms
 */

/**
 * 150. 逆波兰表达式求值
 * Evaluate Reverse Polish Notation
 */
func EvalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		switch token {
		case "+":
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a+b)
		case "-":
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a-b)
		case "*":
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a*b)
		case "/":
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			// Go 中的整数除法自动向零截断
			stack = append(stack, a/b)
		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}

/**
 * 最小栈实现
 * Min Stack implementation
 */
type MinStack struct {
	stack    []int
	minStack []int
}

func NewMinStack() *MinStack {
	return &MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (ms *MinStack) Push(val int) {
	ms.stack = append(ms.stack, val)
	if len(ms.minStack) == 0 || val <= ms.GetMin() {
		ms.minStack = append(ms.minStack, val)
	}
}

func (ms *MinStack) Pop() {
	if len(ms.stack) == 0 {
		return
	}
	val := ms.stack[len(ms.stack)-1]
	ms.stack = ms.stack[:len(ms.stack)-1]
	if val == ms.GetMin() {
		ms.minStack = ms.minStack[:len(ms.minStack)-1]
	}
}

func (ms *MinStack) Top() int {
	if len(ms.stack) == 0 {
		return 0
	}
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	if len(ms.minStack) == 0 {
		return 0
	}
	return ms.minStack[len(ms.minStack)-1]
}

/**
 * 20. 有效的括号
 * Valid Parentheses
 */
func IsValid(s string) bool {
	stack := []rune{}
	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if openBracket, exists := mapping[char]; exists {
			// 这是一个闭括号
			if len(stack) == 0 || stack[len(stack)-1] != openBracket {
				return false
			}
			stack = stack[:len(stack)-1] // pop
		} else {
			// 这是一个开括号
			stack = append(stack, char) // push
		}
	}

	return len(stack) == 0
}

/**
 * 503. 下一个更大元素 II
 * Next Greater Element II (循环数组)
 */
func NextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	
	stack := []int{}
	
	// 模拟循环数组，遍历两遍
	for i := 2*n - 1; i >= 0; i-- {
		x := nums[i%n]
		// 维护单调递减栈
		for len(stack) > 0 && x >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		// 如果栈不为空且在第一轮遍历中，记录答案
		if len(stack) > 0 && i < n {
			ans[i] = stack[len(stack)-1]
		}
		stack = append(stack, x)
	}
	
	return ans
}

/**
 * 739. 每日温度
 * Daily Temperatures
 */
func DailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	stack := []int{} // 存储索引
	
	// 从右向左遍历
	for i := n - 1; i >= 0; i-- {
		t := temperatures[i]
		// 维护单调递减栈（存储索引）
		for len(stack) > 0 && t >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		// 如果栈不为空，计算天数差
		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	
	return ans
}

/**
 * 739. 每日温度 (从左向右的实现)
 * Daily Temperatures (left to right approach)
 */
func DailyTemperaturesLeftToRight(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	stack := []int{} // 存储索引
	
	for i, t := range temperatures {
		// 当前温度比栈顶温度高时，可以为栈中的元素找到答案
		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[j] = i - j
		}
		stack = append(stack, i)
	}
	
	return ans
}

// DemonstrateStackAlgorithms 演示栈算法
func DemonstrateStackAlgorithms() {
	fmt.Println("=== 栈算法演示 ===")
	
	// 逆波兰表达式求值
	tokens := []string{"2", "1", "+", "3", "*"}
	result := EvalRPN(tokens)
	fmt.Printf("逆波兰表达式 %v = %d\n", tokens, result)
	
	// 下一个更大元素 II
	nums := []int{1, 2, 1}
	nextGreater := NextGreaterElements(nums)
	fmt.Printf("数组: %v\n", nums)
	fmt.Printf("下一个更大元素: %v\n", nextGreater)
	
	// 每日温度
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	days := DailyTemperatures(temperatures)
	fmt.Printf("温度: %v\n", temperatures)
	fmt.Printf("等待天数: %v\n", days)
	fmt.Println()
}