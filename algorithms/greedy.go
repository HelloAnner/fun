package algorithms

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

/**
 * 贪心算法相关实现
 * Greedy algorithms
 */

/**
 * 55. 跳跃游戏
 * Jump Game
 */
func CanJump(nums []int) bool {
	maxReach := 0

	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}
	}

	return true
}

/**
 * 45. 跳跃游戏 II
 * Jump Game II
 */
func Jump(nums []int) int {
	jumps := 0
	currentEnd := 0
	farthest := 0

	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}

		if i == currentEnd {
			jumps++
			currentEnd = farthest
		}
	}

	return jumps
}

/**
 * 121. 买卖股票的最佳时机
 * Best Time to Buy and Sell Stock
 */
func MaxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}

/**
 * 135. 分发糖果
 * Candy
 */
func Candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1
	}

	// 从左到右遍历
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// 从右到左遍历
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if candies[i+1]+1 > candies[i] {
				candies[i] = candies[i+1] + 1
			}
		}
	}

	total := 0
	for _, candy := range candies {
		total += candy
	}

	return total
}

/**
 * 2740. 找出分区值
 * Find the Value of the Partition
 */
func FindValueOfPartition(nums []int) int {
	sort.Ints(nums)
	minDiff := math.MaxInt32

	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}

/**
 * 68. 文本左右对齐
 * Text Justification
 */
func FullJustify(words []string, maxWidth int) []string {
	result := []string{}
	currentLine := []string{}
	currentLength := 0

	for _, word := range words {
		// 检查是否可以添加当前单词到当前行
		if currentLength+len(word)+len(currentLine) > maxWidth {
			// 处理当前行
			result = append(result, justifyLine(currentLine, maxWidth, false))
			currentLine = []string{word}
			currentLength = len(word)
		} else {
			currentLine = append(currentLine, word)
			currentLength += len(word)
		}
	}

	// 处理最后一行（左对齐）
	if len(currentLine) > 0 {
		result = append(result, justifyLine(currentLine, maxWidth, true))
	}

	return result
}

func justifyLine(words []string, maxWidth int, isLastLine bool) string {
	if len(words) == 1 || isLastLine {
		// 左对齐
		line := strings.Join(words, " ")
		return line + strings.Repeat(" ", maxWidth-len(line))
	}

	totalChars := 0
	for _, word := range words {
		totalChars += len(word)
	}

	totalSpaces := maxWidth - totalChars
	gaps := len(words) - 1
	spacesPerGap := totalSpaces / gaps
	extraSpaces := totalSpaces % gaps

	var result strings.Builder
	for i := 0; i < len(words)-1; i++ {
		result.WriteString(words[i])
		result.WriteString(strings.Repeat(" ", spacesPerGap))
		if i < extraSpaces {
			result.WriteString(" ")
		}
	}
	result.WriteString(words[len(words)-1])

	return result.String()
}

// DemonstrateCandyAlgorithm 演示分发糖果算法
func DemonstrateCandyAlgorithm() {
	fmt.Println("=== 贪心算法 - 分发糖果 ===")
	
	ratings := []int{1, 0, 2}
	result := Candy(ratings)
	fmt.Printf("评分: %v\n", ratings)
	fmt.Printf("最少糖果数: %d\n", result)
	
	ratings2 := []int{1, 2, 2}
	result2 := Candy(ratings2)
	fmt.Printf("评分: %v\n", ratings2)
	fmt.Printf("最少糖果数: %d\n", result2)
	
	// 跳跃游戏演示
	nums := []int{2, 3, 1, 1, 4}
	canJump := CanJump(nums)
	fmt.Printf("数组: %v, 能否跳到最后: %t\n", nums, canJump)
	
	// 买卖股票演示
	prices := []int{7, 1, 5, 3, 6, 4}
	profit := MaxProfit(prices)
	fmt.Printf("股票价格: %v, 最大利润: %d\n", prices, profit)
	fmt.Println()
}