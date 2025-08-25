package algorithms

import "math"

/**
 * 滑动窗口相关算法实现
 * Sliding Window related algorithms
 */

/**
 * 3. 无重复字符的最长子串
 * Longest Substring Without Repeating Characters
 */
func LengthOfLongestSubstring(s string) int {
	charSet := make(map[byte]bool)
	left := 0
	maxLength := 0

	for right := 0; right < len(s); right++ {
		for charSet[s[right]] {
			delete(charSet, s[left])
			left++
		}
		charSet[s[right]] = true
		if right-left+1 > maxLength {
			maxLength = right - left + 1
		}
	}

	return maxLength
}

/**
 * 76. 最小覆盖子串
 * Minimum Window Substring
 */
func MinWindow(s, t string) string {
	if len(s) < len(t) {
		return ""
	}

	need := make(map[byte]int)
	window := make(map[byte]int)

	// 统计 t 中字符频次
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right := 0, 0
	valid := 0
	start := 0
	length := math.MaxInt32

	for right < len(s) {
		c := s[right]
		right++

		if _, exists := need[c]; exists {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}

			d := s[left]
			left++

			if _, exists := need[d]; exists {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}

/**
 * 209. 长度最小的子数组
 * Minimum Size Subarray Sum
 */
func MinSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	minLen := math.MaxInt32

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		for sum >= target {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}

	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}

/**
 * 219. 存在重复元素 II
 * Contains Duplicate II
 */
func ContainsNearbyDuplicate(nums []int, k int) bool {
	numMap := make(map[int]int)

	for i, num := range nums {
		if lastIndex, exists := numMap[num]; exists && i-lastIndex <= k {
			return true
		}
		numMap[num] = i
	}

	return false
}

/**
 * 438. 找到字符串中所有字母异位词
 * Find All Anagrams in a String
 */
func FindAnagrams(s, p string) []int {
	result := []int{}
	if len(s) < len(p) {
		return result
	}

	need := make(map[byte]int)
	window := make(map[byte]int)

	// 统计 p 中字符频次
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}

	left, right := 0, 0
	valid := 0

	for right < len(s) {
		c := s[right]
		right++

		if _, exists := need[c]; exists {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for right-left >= len(p) {
			if valid == len(need) {
				result = append(result, left)
			}

			d := s[left]
			left++

			if _, exists := need[d]; exists {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	return result
}

/**
 * 1052. 爱生气的书店老板
 * Grumpy Bookstore Owner
 */
func MaxSatisfied(customers []int, grumpy []int, minutes int) int {
	satisfied := 0

	// 计算不使用技巧时的满意顾客数
	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			satisfied += customers[i]
		}
	}

	// 滑动窗口找到最大增益
	maxGain := 0
	currentGain := 0

	// 初始窗口
	for i := 0; i < minutes; i++ {
		if grumpy[i] == 1 {
			currentGain += customers[i]
		}
	}
	maxGain = currentGain

	// 滑动窗口
	for i := minutes; i < len(customers); i++ {
		if grumpy[i] == 1 {
			currentGain += customers[i]
		}
		if grumpy[i-minutes] == 1 {
			currentGain -= customers[i-minutes]
		}
		if currentGain > maxGain {
			maxGain = currentGain
		}
	}

	return satisfied + maxGain
}
