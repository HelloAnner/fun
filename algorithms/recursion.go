package algorithms

import "fmt"

// TowerOfHanoi 汉诺塔问题 - 递归解法
// n: 圆盘数量, source: 起始柱子, auxiliary: 辅助柱子, target: 目标柱子
func TowerOfHanoi(n int, source, auxiliary, target int) {
	if n > 0 {
		// 将 n-1 个圆盘从起始柱子移动到辅助柱子
		TowerOfHanoi(n-1, source, target, auxiliary)
		// 将最大的圆盘从起始柱子移动到目标柱子
		fmt.Printf("Move disk %d from rod %d to rod %d\n", n, source, target)
		// 将 n-1 个圆盘从辅助柱子移动到目标柱子
		TowerOfHanoi(n-1, auxiliary, source, target)
	}
}

// TowerOfHanoiSteps 汉诺塔问题 - 返回移动步骤而不是打印
func TowerOfHanoiSteps(n int, source, auxiliary, target int) []string {
	var steps []string
	towerOfHanoiHelper(n, source, auxiliary, target, &steps)
	return steps
}

func towerOfHanoiHelper(n int, source, auxiliary, target int, steps *[]string) {
	if n > 0 {
		// 将 n-1 个圆盘从起始柱子移动到辅助柱子
		towerOfHanoiHelper(n-1, source, target, auxiliary, steps)
		// 将最大的圆盘从起始柱子移动到目标柱子
		*steps = append(*steps, fmt.Sprintf("Move disk %d from rod %d to rod %d", n, source, target))
		// 将 n-1 个圆盘从辅助柱子移动到目标柱子
		towerOfHanoiHelper(n-1, auxiliary, source, target, steps)
	}
}

// Factorial 阶乘 - 递归实现
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// Fibonacci 斐波那契数列 - 递归实现
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// FibonacciMemo 斐波那契数列 - 带记忆化的递归实现
func FibonacciMemo(n int) int {
	memo := make(map[int]int)
	return fibonacciMemoHelper(n, memo)
}

func fibonacciMemoHelper(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if val, exists := memo[n]; exists {
		return val
	}
	memo[n] = fibonacciMemoHelper(n-1, memo) + fibonacciMemoHelper(n-2, memo)
	return memo[n]
}

// Power 幂运算 - 递归实现
func Power(base, exp int) int {
	if exp == 0 {
		return 1
	}
	if exp == 1 {
		return base
	}
	if exp%2 == 0 {
		half := Power(base, exp/2)
		return half * half
	}
	return base * Power(base, exp-1)
}

// DemonstrateRecursionAlgorithms 演示递归算法
func DemonstrateRecursionAlgorithms() {
	fmt.Println("=== 递归算法演示 ===")
	
	// 汉诺塔问题
	fmt.Println("汉诺塔问题 (3个圆盘):")
	TowerOfHanoi(3, 1, 2, 3)
	fmt.Println()
	
	// 阶乘
	n := 5
	fmt.Printf("阶乘 %d! = %d\n", n, Factorial(n))
	
	// 斐波那契数列
	fmt.Printf("斐波那契数列 F(%d) = %d\n", 10, Fibonacci(10))
	fmt.Printf("斐波那契数列 (记忆化) F(%d) = %d\n", 10, FibonacciMemo(10))
	
	// 幂运算
	fmt.Printf("幂运算 2^10 = %d\n", Power(2, 10))
	fmt.Println()
}