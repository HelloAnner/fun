package algorithms

import (
	"sort"
	"strconv"
	"strings"
	"algorithms-go/data_structures"
)

// CombinationSum 组合总和
// 找出 candidates 中可以使数字和为目标数 target 的所有不同组合
func CombinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var path []int
	
	sort.Ints(candidates) // 排序以便剪枝
	
	var dfs func(begin int, target int)
	dfs = func(begin int, target int) {
		if target == 0 {
			// 找到一个有效组合
			combination := make([]int, len(path))
			copy(combination, path)
			result = append(result, combination)
			return
		}
		
		for i := begin; i < len(candidates); i++ {
			if target-candidates[i] < 0 {
				break // 剪枝：后面的数字更大，不可能满足条件
			}
			
			path = append(path, candidates[i])
			dfs(i, target-candidates[i]) // 可以重复使用同一个数字
			path = path[:len(path)-1]    // 回溯
		}
	}
	
	dfs(0, target)
	return result
}

// CombinationSumIII 组合总和 III
// 找出所有相加之和为 n 的 k 个数的组合，只使用数字1到9
func CombinationSumIII(k int, n int) [][]int {
	var result [][]int
	var path []int
	
	var dfs func(start int, remaining int, count int)
	dfs = func(start int, remaining int, count int) {
		if count == k {
			if remaining == 0 {
				combination := make([]int, len(path))
				copy(combination, path)
				result = append(result, combination)
			}
			return
		}
		
		for i := start; i <= 9; i++ {
			if remaining-i < 0 {
				break // 剪枝
			}
			
			path = append(path, i)
			dfs(i+1, remaining-i, count+1)
			path = path[:len(path)-1] // 回溯
		}
	}
	
	dfs(1, n, 0)
	return result
}

// BinaryTreePaths 二叉树的所有路径
// 返回从根节点到叶子节点的所有路径
func BinaryTreePaths(root *data_structures.TreeNode) []string {
	if root == nil {
		return []string{}
	}
	
	var result []string
	var path []string
	
	var dfs func(node *data_structures.TreeNode)
	dfs = func(node *data_structures.TreeNode) {
		if node == nil {
			return
		}
		
		path = append(path, strconv.Itoa(node.Val))
		
		// 如果是叶子节点，添加路径到结果
		if node.Left == nil && node.Right == nil {
			result = append(result, strings.Join(path, "->"))
		} else {
			// 继续遍历左右子树
			dfs(node.Left)
			dfs(node.Right)
		}
		
		// 回溯
		path = path[:len(path)-1]
	}
	
	dfs(root)
	return result
}

// AllPathsSourceTarget 所有可能的路径
// 找到从节点 0 到节点 n-1 的所有可能路径
func AllPathsSourceTarget(graph [][]int) [][]int {
	var result [][]int
	var path []int
	target := len(graph) - 1
	
	var dfs func(node int)
	dfs = func(node int) {
		path = append(path, node)
		
		if node == target {
			// 到达目标节点，添加路径到结果
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			result = append(result, pathCopy)
		} else {
			// 继续遍历邻接节点
			for _, neighbor := range graph[node] {
				dfs(neighbor)
			}
		}
		
		// 回溯
		path = path[:len(path)-1]
	}
	
	dfs(0)
	return result
}

// IsShiftString 字符串轮转
// 判断字符串 A 是否可以通过轮转得到字符串 B
func IsShiftString(A, B string) bool {
	if len(A) != len(B) {
		return false
	}
	
	doubleA := A + A
	return strings.Contains(doubleA, B)
}

// Permute 全排列
// 给定一个不含重复数字的数组 nums，返回其所有可能的全排列
func Permute(nums []int) [][]int {
	var result [][]int
	var path []int
	used := make([]bool, len(nums))
	
	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			permutation := make([]int, len(path))
			copy(permutation, path)
			result = append(result, permutation)
			return
		}
		
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			
			path = append(path, nums[i])
			used[i] = true
			dfs()
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	
	dfs()
	return result
}

// Subsets 子集
// 给你一个整数数组 nums，返回该数组所有可能的子集
func Subsets(nums []int) [][]int {
	var result [][]int
	var path []int
	
	var dfs func(start int)
	dfs = func(start int) {
		// 当前路径就是一个子集
		subset := make([]int, len(path))
		copy(subset, path)
		result = append(result, subset)
		
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	
	dfs(0)
	return result
}

// LetterCombinations 电话号码的字母组合
// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合
func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	
	phoneMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	
	var result []string
	var path []byte
	
	var dfs func(index int)
	dfs = func(index int) {
		if index == len(digits) {
			result = append(result, string(path))
			return
		}
		
		letters := phoneMap[digits[index]]
		for i := 0; i < len(letters); i++ {
			path = append(path, letters[i])
			dfs(index + 1)
			path = path[:len(path)-1]
		}
	}
	
	dfs(0)
	return result
}