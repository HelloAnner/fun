package algorithms

import "fmt"

/**
 * 动态规划相关算法实现
 * Dynamic Programming related algorithms
 */

/**
 * 70. 爬楼梯
 * Climbing Stairs
 */
func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}

	prev2 := 1
	prev1 := 2

	for i := 3; i <= n; i++ {
		current := prev1 + prev2
		prev2 = prev1
		prev1 = current
	}

	return prev1
}

/**
 * 322. 零钱兑换
 * Coin Change
 */
func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				if dp[i-coin]+1 < dp[i] {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

/**
 * 300. 最长递增子序列
 * Longest Increasing Subsequence
 */
func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}

	maxLen := dp[0]
	for _, length := range dp {
		if length > maxLen {
			maxLen = length
		}
	}

	return maxLen
}

/**
 * 198. 打家劫舍
 * House Robber
 */
func Rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	prev2 := nums[0]
	prev1 := max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		current := max(prev1, prev2+nums[i])
		prev2 = prev1
		prev1 = current
	}

	return prev1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
 * 139. 单词拆分
 * Word Break
 */
func WordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

/**
 * 1035. 不相交的线
 * Uncrossed Lines
 */
func MaxUncrossedLines(nums1, nums2 []int) int {
	m := len(nums1)
	n := len(nums2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

/**
 * 1216. 验证回文字符串 III
 * Valid Palindrome III
 */
func IsValidPalindrome(s string, k int) bool {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 计算最长回文子序列的长度
	for length := 1; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if i == j {
				dp[i][j] = 1
			} else if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	// 需要删除的字符数 = 总长度 - 最长回文子序列长度
	return n-dp[0][n-1] <= k
}

/**
 * 1186. 删除一次得到子数组最大和
 * Maximum Subarray Sum with One Deletion
 */
func MaximumSum(arr []int) int {
	n := len(arr)
	if n == 1 {
		return arr[0]
	}
	
	// dp0[i] 表示以 arr[i] 结尾且没有删除元素的最大子数组和
	// dp1[i] 表示以 arr[i] 结尾且删除了一个元素的最大子数组和
	dp0 := make([]int, n)
	dp1 := make([]int, n)
	
	dp0[0] = arr[0]
	dp1[0] = 0 // 不能删除唯一的元素
	
	result := arr[0]
	
	for i := 1; i < n; i++ {
		// 没有删除元素的情况
		dp0[i] = max(arr[i], dp0[i-1]+arr[i])
		
		// 删除了一个元素的情况
		// 要么删除当前元素（从 dp0[i-1] 转移）
		// 要么不删除当前元素（从 dp1[i-1] 转移）
		dp1[i] = max(dp0[i-1], dp1[i-1]+arr[i])
		
		result = max(result, max(dp0[i], dp1[i]))
	}
	
	return result
}

/**
 * 2708. 一个小组的最大实力值
 * Maximum Strength of a Group
 */
func MaxStrength(nums []int) int64 {
	if len(nums) == 1 {
		return int64(nums[0])
	}
	
	// mn 表示当前最小值，mx 表示当前最大值
	mn := int64(nums[0])
	mx := int64(nums[0])
	
	for i := 1; i < len(nums); i++ {
		x := int64(nums[i])
		// 计算新的最小值和最大值
		newMn := minInt64(mn, x, mn*x, mx*x)
		newMx := maxInt64(mx, x, mn*x, mx*x)
		mn = newMn
		mx = newMx
	}
	
	return mx
}

// minInt64 返回多个 int64 值中的最小值
func minInt64(values ...int64) int64 {
	if len(values) == 0 {
		return 0
	}
	min := values[0]
	for _, v := range values[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

// maxInt64 返回多个 int64 值中的最大值
func maxInt64(values ...int64) int64 {
	if len(values) == 0 {
		return 0
	}
	max := values[0]
	for _, v := range values[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// DemonstrateDynamicProgrammingAlgorithms 演示动态规划算法
func DemonstrateDynamicProgrammingAlgorithms() {
	fmt.Println("=== 动态规划算法演示 ===")
	
	// 不相交的线
	nums1 := []int{1, 4, 2}
	nums2 := []int{1, 2, 4}
	lines := MaxUncrossedLines(nums1, nums2)
	fmt.Printf("数组1: %v, 数组2: %v\n", nums1, nums2)
	fmt.Printf("最大不相交线数: %d\n", lines)
	
	// 删除一次得到子数组最大和
	arr := []int{1, -2, 0, 3}
	maxSum := MaximumSum(arr)
	fmt.Printf("数组: %v\n", arr)
	fmt.Printf("删除一次后的最大子数组和: %d\n", maxSum)
	
	// 一个小组的最大实力值
	nums := []int{3, -1, -5, 2, 5, -9}
	maxStrength := MaxStrength(nums)
	fmt.Printf("数组: %v\n", nums)
	fmt.Printf("最大实力值: %d\n", maxStrength)
	fmt.Println()
}