package algorithms

import "sort"

/**
 * 哈希表相关算法实现
 * Hash table related algorithms
 */

/**
 * 2766. 重新放置石块
 * Relocate Marbles
 */
func RelocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
	stoneSet := make(map[int]bool)
	for _, num := range nums {
		stoneSet[num] = true
	}

	for i := 0; i < len(moveFrom); i++ {
		from := moveFrom[i]
		to := moveTo[i]
		delete(stoneSet, from)
		stoneSet[to] = true
	}

	result := make([]int, 0, len(stoneSet))
	for stone := range stoneSet {
		result = append(result, stone)
	}

	sort.Ints(result)
	return result
}

/**
 * 1. 两数之和
 * Two Sum
 */
func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if idx, exists := numMap[complement]; exists {
			return []int{idx, i}
		}
		numMap[num] = i
	}

	return []int{}
}

/**
 * 202. 快乐数
 * Happy Number
 */
func IsHappy(n int) bool {
	seen := make(map[int]bool)

	for n != 1 && !seen[n] {
		seen[n] = true
		n = getNext(n)
	}

	return n == 1
}

func getNext(n int) int {
	totalSum := 0
	for n > 0 {
		digit := n % 10
		n /= 10
		totalSum += digit * digit
	}
	return totalSum
}

/**
 * 997. 找到小镇的法官
 * Find the Town Judge
 */
func FindJudge(n int, trust [][]int) int {
	if len(trust) < n-1 {
		return -1
	}

	inDegree := make([]int, n+1)
	outDegree := make([]int, n+1)

	for _, t := range trust {
		a, b := t[0], t[1]
		outDegree[a]++
		inDegree[b]++
	}

	for i := 1; i <= n; i++ {
		if inDegree[i] == n-1 && outDegree[i] == 0 {
			return i
		}
	}

	return -1
}

/**
 * 1436. 旅行终点站
 * Destination City
 */
func DestCity(paths [][]string) string {
	outgoing := make(map[string]bool)

	for _, path := range paths {
		from := path[0]
		outgoing[from] = true
	}

	for _, path := range paths {
		to := path[1]
		if !outgoing[to] {
			return to
		}
	}

	return ""
}