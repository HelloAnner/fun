package algorithms

// 位操作相关算法实现
// Bit Manipulation related algorithms

// 136. 只出现一次的数字
// Single Number
// 在数组中找出只出现一次的数字（其他数字都出现两次）
// https://leetcode.cn/problems/single-number/description/
func SingleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

// 137. 只出现一次的数字 II
// Single Number II
// 在数组中找出只出现一次的数字（其他数字都出现三次）
// https://leetcode.cn/problems/single-number-ii/description/
func SingleNumberII(nums []int) int {
	ones := 0
	twos := 0

	for _, num := range nums {
		ones = (ones ^ num) & ^twos
		twos = (twos ^ num) & ^ones
	}

	return ones
}

// 260. 只出现一次的数字 III
// Single Number III
// 在数组中找出两个只出现一次的数字（其他数字都出现两次）
// https://leetcode.cn/problems/single-number-iii/description/
func SingleNumberIII(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	// 找到最右边的1
	rightmostBit := xor & (-xor)

	num1 := 0
	num2 := 0

	for _, num := range nums {
		if num&rightmostBit != 0 {
			num1 ^= num
		} else {
			num2 ^= num
		}
	}

	return []int{num1, num2}
}

// 191. 位1的个数
// Number of 1 Bits
// 计算一个无符号整数的二进制表示中1的个数
// https://leetcode.cn/problems/number-of-1-bits/description/
func HammingWeight(n uint32) int {
	count := 0
	for n != 0 {
		count++
		n &= n - 1 // 清除最右边的1
	}
	return count
}

// 338. 比特位计数
// Counting Bits
// 计算从0到n的每个数字的二进制表示中1的个数
// https://leetcode.cn/problems/counting-bits/description/
func CountBits(n int) []int {
	result := make([]int, n+1)

	for i := 1; i <= n; i++ {
		result[i] = result[i>>1] + (i & 1)
	}

	return result
}

// 461. 汉明距离
// Hamming Distance
// 计算两个整数二进制表示中不同位的个数
// https://leetcode.cn/problems/hamming-distance/description/
func HammingDistance(x, y int) int {
	xor := x ^ y
	count := 0

	for xor != 0 {
		count++
		xor &= xor - 1
	}

	return count
}

// 477. 汉明距离总和
// Total Hamming Distance
// 计算数组中所有数对之间的汉明距离总和
// https://leetcode.cn/problems/total-hamming-distance/description/
func TotalHammingDistance(nums []int) int {
	total := 0
	n := len(nums)

	for i := 0; i < 32; i++ {
		ones := 0
		for _, num := range nums {
			ones += (num >> i) & 1
		}
		total += ones * (n - ones)
	}

	return total
}

// 190. 颠倒二进制位
// Reverse Bits
// 颠倒给定的32位无符号整数的二进制位
// https://leetcode.cn/problems/reverse-bits/description/
func ReverseBits(n uint32) uint32 {
	var result uint32 = 0
	for i := 0; i < 32; i++ {
		result = (result << 1) | (n & 1)
		n >>= 1
	}
	return result
}

// 231. 2的幂
// Power of Two
// 判断一个整数是否是2的幂次方
// https://leetcode.cn/problems/power-of-two/description/
func IsPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

// 342. 4的幂
// Power of Four
// 判断一个整数是否是4的幂次方
// https://leetcode.cn/problems/power-of-four/description/
func IsPowerOfFour(n int) bool {
	return n > 0 && (n&(n-1)) == 0 && (n&0x55555555) != 0
}

// 371. 两整数之和
// Sum of Two Integers
// 不使用+和-运算符，计算两整数之和
// https://leetcode.cn/problems/sum-of-two-integers/description/
func GetSum(a, b int) int {
	for b != 0 {
		carry := (a & b) << 1
		a = a ^ b
		b = carry
	}
	return a
}

// 389. 找不同
// Find the Difference
// 在字符串t中找到比字符串s多出的那个字符
// https://leetcode.cn/problems/find-the-difference/description/
func FindTheDifference(s, t string) byte {
	xor := 0

	for _, char := range s {
		xor ^= int(char)
	}

	for _, char := range t {
		xor ^= int(char)
	}

	return byte(xor)
}

// 405. 数字转换为十六进制数
// Convert a Number to Hexadecimal
// 将给定的整数转换为十六进制字符串表示
// https://leetcode.cn/problems/convert-a-number-to-hexadecimal/description/
func ToHex(num int) string {
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
