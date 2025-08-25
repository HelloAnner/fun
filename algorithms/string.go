package algorithms

import (
	"sort"
	"strings"
)

/**
 * 字符串相关算法实现
 * String related algorithms
 */

/**
 * 28. 找出字符串中第一个匹配项的下标
 * Find the Index of the First Occurrence in a String
 */
func StrStr(haystack, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}

/**
 * 6. Z 字形变换
 * Zigzag Conversion
 */
func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]string, min(numRows, len(s)))
	currentRow := 0
	goingDown := false

	for _, char := range s {
		rows[currentRow] += string(char)

		if currentRow == 0 || currentRow == numRows-1 {
			goingDown = !goingDown
		}

		if goingDown {
			currentRow++
		} else {
			currentRow--
		}
	}

	return strings.Join(rows, "")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * 65. 有效数字
 * Valid Number
 */
func IsNumber(s string) bool {
	s = strings.TrimSpace(s)
	hasNum := false
	hasE := false
	hasDot := false

	for i, char := range s {
		if char >= '0' && char <= '9' {
			hasNum = true
		} else if char == '.' {
			if hasDot || hasE {
				return false
			}
			hasDot = true
		} else if char == 'e' || char == 'E' {
			if hasE || !hasNum {
				return false
			}
			hasE = true
			hasNum = false // e后面必须有数字
		} else if char == '+' || char == '-' {
			if i != 0 && s[i-1] != 'e' && s[i-1] != 'E' {
				return false
			}
		} else {
			return false
		}
	}

	return hasNum
}

/**
 * 字符串翻转检查
 * Check if string can be shifted to match another
 */
func CanShift(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	return strings.Contains(a+a, b)
}

/**
 * 字符串倒序修改
 * Reverse string modification
 */
func ReverseStringModify(s string) string {
	chars := []rune(s)
	left := 0
	right := len(chars) - 1

	for left < right {
		chars[left], chars[right] = chars[right], chars[left]
		left++
		right--
	}

	return string(chars)
}

/**
 * 字母异位词分组
 * Group Anagrams
 */
func GroupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, str := range strs {
		chars := []rune(str)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		key := string(chars)

		groups[key] = append(groups[key], str)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

/**
 * 205. 同构字符串
 * Isomorphic Strings
 */
func IsIsomorphic(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapS := make(map[rune]rune)
	mapT := make(map[rune]rune)

	sRunes := []rune(s)
	tRunes := []rune(t)

	for i := 0; i < len(sRunes); i++ {
		charS := sRunes[i]
		charT := tRunes[i]

		if mapped, exists := mapS[charS]; exists {
			if mapped != charT {
				return false
			}
		} else {
			mapS[charS] = charT
		}

		if mapped, exists := mapT[charT]; exists {
			if mapped != charS {
				return false
			}
		} else {
			mapT[charT] = charS
		}
	}

	return true
}

/**
 * 242. 有效的字母异位词
 * Valid Anagram
 */
func IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[rune]int)

	for _, char := range s {
		count[char]++
	}

	for _, char := range t {
		if count[char] == 0 {
			return false
		}
		count[char]--
		if count[char] == 0 {
			delete(count, char)
		}
	}

	return len(count) == 0
}

/**
 * 290. 单词规律
 * Word Pattern
 */
func WordPattern(pattern, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	charToWord := make(map[rune]string)
	wordToChar := make(map[string]rune)

	patternRunes := []rune(pattern)

	for i := 0; i < len(patternRunes); i++ {
		char := patternRunes[i]
		word := words[i]

		if mapped, exists := charToWord[char]; exists {
			if mapped != word {
				return false
			}
		} else {
			charToWord[char] = word
		}

		if mapped, exists := wordToChar[word]; exists {
			if mapped != char {
				return false
			}
		} else {
			wordToChar[word] = char
		}
	}

	return true
}
