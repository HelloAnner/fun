package algorithms

/**
 * 位操作相关算法实现
 * Bit Manipulation related algorithms
 */

/**
 * 136. 只出现一次的数字
 * Single Number
 */
func SingleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

/**
 * 137. 只出现一次的数字 II
 * Single Number II
 */
func SingleNumberII(nums []int) int {
	ones := 0
	twos := 0

	for _, num := range nums {
		ones = (ones ^ num) & ^twos
		twos = (twos ^ num) & ^ones
	}

	return ones
}

/**
 * 260. 只出现一次的数字 III
 * Single Number III
 */
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

/**
 * 191. 位1的个数
 * Number of 1 Bits
 */
func HammingWeight(n uint32) int {
	count := 0
	for n != 0 {
		count++
		n &= n - 1 // 清除最右边的1
	}
	return count
}

/**
 * 338. 比特位计数
 * Counting Bits
 */
func CountBits(n int) []int {
	result := make([]int, n+1)

	for i := 1; i <= n; i++ {
		result[i] = result[i>>1] + (i & 1)
	}

	return result
}

/**
 * 461. 汉明距离
 * Hamming Distance
 */
func HammingDistance(x, y int) int {
	xor := x ^ y
	count := 0

	for xor != 0 {
		count++
		xor &= xor - 1
	}

	return count
}

/**
 * 477. 汉明距离总和
 * Total Hamming Distance
 */
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

/**
 * 190. 颠倒二进制位
 * Reverse Bits
 */
func ReverseBits(n uint32) uint32 {
	var result uint32 = 0
	for i := 0; i < 32; i++ {
		result = (result << 1) | (n & 1)
		n >>= 1
	}
	return result
}

/**
 * 231. 2的幂
 * Power of Two
 */
func IsPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

/**
 * 342. 4的幂
 * Power of Four
 */
func IsPowerOfFour(n int) bool {
	return n > 0 && (n&(n-1)) == 0 && (n&0x55555555) != 0
}

/**
 * 371. 两整数之和
 * Sum of Two Integers
 */
func GetSum(a, b int) int {
	for b != 0 {
		carry := (a & b) << 1
		a = a ^ b
		b = carry
	}
	return a
}

/**
 * 389. 找不同
 * Find the Difference
 */
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

/**
 * 405. 数字转换为十六进制数
 * Convert a Number to Hexadecimal
 */
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