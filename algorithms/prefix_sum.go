package algorithms

import "fmt"

// MinimumLevels 计算Alice获得比Bob更多分数所需的最少关卡数目
// 把0视作-1，找一个最短前缀，其元素和大于剩余元素和
func MinimumLevels(possible []int) int {
	n := len(possible)
	s := 0
	for _, x := range possible {
		if x == 1 {
			s += 2
		} else {
			s -= 2
		}
	}
	s -= 2 // 减去最后一个元素，因为Alice不能完成所有关卡
	
	pre := 0
	for i := 0; i < n-1; i++ {
		if possible[i] == 1 {
			pre += 2
		} else {
			pre -= 2
		}
		if pre > s-pre {
			return i + 1
		}
	}
	return -1
}

// IsArraySpecial 检查子数组是否为特殊数组（相邻元素奇偶性不同）
func IsArraySpecial(nums []int, queries [][]int) []bool {
	n := len(nums)
	// 构建前缀和数组，记录相邻元素奇偶性相同的位置
	prefixSum := make([]int, n)
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1]
		if nums[i]%2 == nums[i-1]%2 {
			prefixSum[i]++
		}
	}
	
	result := make([]bool, len(queries))
	for i, query := range queries {
		from, to := query[0], query[1]
		// 如果区间内没有相邻元素奇偶性相同，则为特殊数组
		result[i] = prefixSum[to] == prefixSum[from]
	}
	return result
}

// SummaryRanges 汇总区间，将连续的数字合并为区间
func SummaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	
	var result []string
	start := 0
	
	for i := 1; i <= len(nums); i++ {
		// 如果到达末尾或者当前数字不连续
		if i == len(nums) || nums[i] != nums[i-1]+1 {
			if start == i-1 {
				// 单个数字
				result = append(result, formatRange(nums[start], nums[start]))
			} else {
				// 区间
				result = append(result, formatRange(nums[start], nums[i-1]))
			}
			start = i
		}
	}
	
	return result
}

// formatRange 格式化区间字符串
func formatRange(start, end int) string {
	if start == end {
		return fmt.Sprintf("%d", start)
	}
	return fmt.Sprintf("%d->%d", start, end)
}

// DemonstratePrefixSumAlgorithms 演示前缀和算法
func DemonstratePrefixSumAlgorithms() {
	fmt.Println("=== 前缀和算法演示 ===")
	
	// 最少关卡数目
	alice := []int{1, 1, 0, 1}
	minLevels := MinimumLevels(alice)
	fmt.Printf("Alice: %v\n", alice)
	fmt.Printf("Alice需要完成的最少关卡数: %d\n", minLevels)
	
	// 特殊数组
	nums := []int{3, 4, 1, 2, 6}
	queries := [][]int{{0, 4}}
	results := IsArraySpecial(nums, queries)
	fmt.Printf("数组: %v, 查询: %v\n", nums, queries)
	fmt.Printf("特殊数组查询结果: %v\n", results)
	
	// 汇总区间
	nums2 := []int{0, 1, 2, 4, 5, 7}
	ranges := SummaryRanges(nums2)
	fmt.Printf("数组: %v\n", nums2)
	fmt.Printf("汇总区间: %v\n", ranges)
	fmt.Println()
}