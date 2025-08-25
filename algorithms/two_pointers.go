package algorithms

import (
	"regexp"
	"strings"
)

/**
 * 双指针相关算法实现
 * Two pointers related algorithms
 */

/**
 * 977. 有序数组的平方
 * Squares of a Sorted Array
 */
func SortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	left, right := 0, n-1

	for i := n - 1; i >= 0; i-- {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]

		if leftSquare > rightSquare {
			result[i] = leftSquare
			left++
		} else {
			result[i] = rightSquare
			right--
		}
	}

	return result
}

/**
 * 88. 合并两个有序数组
 * Merge Sorted Array
 */
func Merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

/**
 * 125. 验证回文串
 * Valid Palindrome
 */
func IsPalindrome(s string) bool {
	// 清理字符串：转换为小写并只保留字母和数字
	reg := regexp.MustCompile(`[^a-zA-Z0-9]`)
	cleaned := reg.ReplaceAllString(strings.ToLower(s), "")

	left, right := 0, len(cleaned)-1

	for left < right {
		if cleaned[left] != cleaned[right] {
			return false
		}
		left++
		right--
	}

	return true
}

/**
 * 392. 判断子序列
 * Is Subsequence
 */
func IsSubsequence(s, t string) bool {
	i, j := 0, 0

	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}

	return i == len(s)
}

/**
 * 151. 反转字符串中的单词
 * Reverse Words in a String
 */
func ReverseWords(s string) string {
	// 去除首尾空格并按空格分割
	words := strings.Fields(s)

	// 反转单词数组
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}
