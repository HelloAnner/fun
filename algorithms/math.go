package algorithms

import (
	"math"
)

// NumberOfRightTriangles 计算直角三角形数目
// 给你一个二维 boolean 矩阵 grid，返回使用 grid 中的 3 个元素可以构建的直角三角形数目
func NumberOfRightTriangles(grid [][]int) int64 {
	n, m := len(grid), len(grid[0])
	col := make([]int, m)
	
	// 统计每列的1的个数
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			col[j] += grid[i][j]
		}
	}
	
	var res int64
	for i := 0; i < n; i++ {
		// 统计当前行的1的个数
		row := 0
		for j := 0; j < m; j++ {
			row += grid[i][j]
		}
		
		// 枚举每个点作为直角顶点
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				res += int64((row - 1) * (col[j] - 1))
			}
		}
	}
	
	return res
}

// MySqrtFloat 计算平方根（浮点数版本）
// 使用二分查找计算平方根
func MySqrtFloat(n float64) float64 {
	if n < 0 {
		return -1 // 负数没有实数平方根
	}
	if n == 0 {
		return 0
	}
	
	low, high := 0.0, n
	if n < 1 {
		high = 1
	}
	
	for high-low > 1e-9 {
		mid := (low + high) / 2
		square := mid * mid
		if math.Abs(square-n) < 1e-9 {
			return mid
		} else if square < n {
			low = mid
		} else {
			high = mid
		}
	}
	
	return (low + high) / 2
}

// FindMaximumPower 找出最大的可达成数字
// 给你两个正整数 num 和 t，返回可以通过对 num 执行最多 t 次操作得到的最大整数
func FindMaximumPower(num, t int) int {
	// 这是一个数学问题，需要根据具体的操作规则来实现
	// 这里提供一个基础框架
	result := num
	for i := 0; i < t; i++ {
		// 根据具体的操作规则来修改 result
		// 这里需要根据题目的具体要求来实现
		result++
	}
	return result
}

// CountBeautifulPairs 美丽下标对的数目
// 计算满足条件的美丽下标对数目
func CountBeautifulPairs(nums []int) int {
	count := 0
	n := len(nums)
	
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// 获取第一个数的第一位数字
			first := nums[i]
			for first >= 10 {
				first /= 10
			}
			
			// 获取第二个数的最后一位数字
			last := nums[j] % 10
			
			// 检查是否互质
			if gcd(first, last) == 1 {
				count++
			}
		}
	}
	
	return count
}

// gcd 计算最大公约数
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// DoubleModularExponentiation 双模幂运算
// 计算 (base^exp1 * base^exp2) % mod
func DoubleModularExponentiation(base, exp1, exp2, mod int) int {
	result1 := modularExponentiation(base, exp1, mod)
	result2 := modularExponentiation(base, exp2, mod)
	return (result1 * result2) % mod
}

// modularExponentiation 模幂运算
func modularExponentiation(base, exp, mod int) int {
	result := 1
	base = base % mod
	
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		exp = exp >> 1
		base = (base * base) % mod
	}
	
	return result
}

// MaxPrimeDistance 质数的最大距离
// 找出数组中质数的最大距离
func MaxPrimeDistance(nums []int) int {
	primes := make([]int, 0)
	
	// 找出所有质数的索引
	for i, num := range nums {
		if isPrime(num) {
			primes = append(primes, i)
		}
	}
	
	if len(primes) < 2 {
		return 0
	}
	
	// 返回第一个和最后一个质数的距离
	return primes[len(primes)-1] - primes[0]
}

// isPrime 判断是否为质数
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	
	return true
}