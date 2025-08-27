package algorithms

import (
	"math"
	"sort"
	"strings"
)

// 贪心算法相关实现
// Greedy algorithms

// 55. 跳跃游戏
// Jump Game
// 判断是否能够从数组的第一个位置跳到最后一个位置
// https://leetcode.cn/problems/jump-game/description/
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

// 45. 跳跃游戏 II
// Jump Game II
// 找到从数组第一个位置跳到最后一个位置的最小跳跃次数
// https://leetcode.cn/problems/jump-game-ii/description/
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

// 121. 买卖股票的最佳时机
// Best Time to Buy and Sell Stock
// 在股票价格数组中找到一次交易的最大利润
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/
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

// 135. 分发糖果
// Candy
// 根据孩子的评分分发糖果，满足相邻孩子评分高的糖果更多
// https://leetcode.cn/problems/candy/description/
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

// 2740. 找出分区值
// Find the Value of the Partition
// 将数组分成两个子数组，找到两个子数组之间最小差值的绝对值
// https://leetcode.cn/problems/find-the-value-of-the-partition/description/
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

// 68. 文本左右对齐
// Text Justification
// 将单词数组按照最大宽度进行文本左右对齐排版
// https://leetcode.cn/problems/text-justification/description/
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

// 665. 非递减数列
// Non-decreasing Array
// 判断是否最多修改一个元素就能使数组变成非递减数列
// https://leetcode.cn/problems/non-decreasing-array/description/
func checkPossibility(nums []int) bool {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			count++
			if count > 1 {
				return false
			}
			if i > 0 && nums[i+1] < nums[i-1] {
				// [3, 4, 2], i=1, nums[0]=3, nums[2]=2. 3 > 2，不能把4改成2，只能把2改成4.
				nums[i+1] = nums[i]
			} else {
				// [2, 4, 3] -> [2, 3, 3]
				nums[i] = nums[i+1]
			}
		}
	}
	return true
}
