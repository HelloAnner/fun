package algorithms

import (
	"algorithms-go/data_structures"
	"sort"
)

// MagicDictionary 魔法字典
type MagicDictionary struct {
	words []string
}

// NewMagicDictionary 构造魔法字典
func NewMagicDictionary() *MagicDictionary {
	return &MagicDictionary{
		words: make([]string, 0),
	}
}

// BuildDict 使用字符串数组 dictionary 设定该数据结构
func (md *MagicDictionary) BuildDict(dictionary []string) {
	md.words = make([]string, len(dictionary))
	copy(md.words, dictionary)
}

// Search 给定一个字符串 searchWord，判定能否只将字符串中一个字母换成另一个字母，使得所形成的新字符串能够与字典中的任一字符串匹配
func (md *MagicDictionary) Search(searchWord string) bool {
	for _, word := range md.words {
		if len(word) != len(searchWord) {
			continue
		}
		
		diffCount := 0
		for i := 0; i < len(word); i++ {
			if word[i] != searchWord[i] {
				diffCount++
				if diffCount > 1 {
					break
				}
			}
		}
		
		if diffCount == 1 {
			return true
		}
	}
	
	return false
}

// PivotIndex 寻找数组的中心下标
// 返回数组的中心下标，如果不存在中心下标，返回 -1
func PivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	
	leftSum := 0
	for i, num := range nums {
		if leftSum == total-leftSum-num {
			return i
		}
		leftSum += num
	}
	
	return -1
}

// CountStudents 在既定时间做作业的学生人数
// 给你两个整数数组 startTime 和 endTime，以及一个整数 queryTime
func CountStudents(startTime []int, endTime []int, queryTime int) int {
	count := 0
	for i := 0; i < len(startTime); i++ {
		if startTime[i] <= queryTime && queryTime <= endTime[i] {
			count++
		}
	}
	return count
}

// SortTheMatrix 将矩阵按对角线排序
// 给你一个 m * n 的整数矩阵 mat，请你将同一条对角线上的元素按升序排序后，返回排好序的矩阵
func SortTheMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
		copy(result[i], mat[i])
	}
	
	// 处理从左上角开始的对角线
	for start := 0; start < m; start++ {
		diagonal := make([]int, 0)
		positions := make([][2]int, 0)
		
		i, j := start, 0
		for i < m && j < n {
			diagonal = append(diagonal, result[i][j])
			positions = append(positions, [2]int{i, j})
			i++
			j++
		}
		
		sort.Ints(diagonal)
		for k, pos := range positions {
			result[pos[0]][pos[1]] = diagonal[k]
		}
	}
	
	// 处理从第一行开始的对角线（除了左上角）
	for start := 1; start < n; start++ {
		diagonal := make([]int, 0)
		positions := make([][2]int, 0)
		
		i, j := 0, start
		for i < m && j < n {
			diagonal = append(diagonal, result[i][j])
			positions = append(positions, [2]int{i, j})
			i++
			j++
		}
		
		sort.Ints(diagonal)
		for k, pos := range positions {
			result[pos[0]][pos[1]] = diagonal[k]
		}
	}
	
	return result
}

// MinOperationsToMakeSpecial 生成特殊数字的最少操作
// 给你一个下标从 0 开始的字符串 num，表示一个非负整数
func MinOperationsToMakeSpecial(num string) int {
	n := len(num)
	if n == 1 {
		return 1
	}
	
	minOps := n // 最坏情况：删除所有数字
	
	// 尝试构造以 00, 25, 50, 75 结尾的数字
	targets := []string{"00", "25", "50", "75"}
	
	for _, target := range targets {
		ops := findMinOperations(num, target)
		if ops < minOps {
			minOps = ops
		}
	}
	
	return minOps
}

// findMinOperations 找到构造特定后缀的最少操作数
func findMinOperations(num, target string) int {
	n := len(num)
	found := make([]int, 0)
	
	// 从右到左找到目标数字的位置
	for i := n - 1; i >= 0; i-- {
		if len(found) < 2 && string(num[i]) == string(target[1-len(found)]) {
			found = append(found, i)
		}
	}
	
	if len(found) < 2 {
		return n // 无法构造，删除所有数字
	}
	
	// 计算需要删除的数字个数
	return n - found[1] - 1
}

// FindMissingAndRepeatedValues 找出缺失和重复的数字
// 给你一个下标从 0 开始的二维整数矩阵 grid，大小为 n * n
func FindMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	count := make(map[int]int)
	
	// 统计每个数字出现的次数
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			count[grid[i][j]]++
		}
	}
	
	var repeated, missing int
	
	// 找出重复和缺失的数字
	for i := 1; i <= n*n; i++ {
		if count[i] == 2 {
			repeated = i
		} else if count[i] == 0 {
			missing = i
		}
	}
	
	return []int{repeated, missing}
}

// MinNumberGame 最小数字游戏
// Alice 和 Bob 正在玩一个游戏
func MinNumberGame(nums []int) []int {
	sort.Ints(nums)
	result := make([]int, len(nums))
	
	for i := 0; i < len(nums); i += 2 {
		// Alice 取最小值，Bob 取次小值
		// Bob 先放，Alice 后放
		result[i] = nums[i+1]   // Bob 的数字
		result[i+1] = nums[i]   // Alice 的数字
	}
	
	return result
}

// ConstructSquareMatrix 构造相同颜色的正方形
// 给你一个二维 3 x 3 的矩阵 grid，每个格子都是一个字符，要么是 'B'，要么是 'W'
func ConstructSquareMatrix(grid [][]byte) bool {
	// 检查所有可能的 2x2 子矩阵
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if grid[i][j] == grid[i][j+1] && 
			   grid[i][j] == grid[i+1][j] && 
			   grid[i][j] == grid[i+1][j+1] {
				return true
			}
		}
	}
	return false
}

// IsSpecialArray 特殊数组 I
// 如果数组中每一对相邻元素都是两个奇偶性不同的数字，则该数组被认为是一个特殊数组
func IsSpecialArray(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]%2 == nums[i+1]%2 {
			return false
		}
	}
	return true
}

// MinDepth 二叉树的最小深度
// 给定一个二叉树，找出其最小深度
func MinDepth(root *data_structures.TreeNode) int {
	if root == nil {
		return 0
	}
	
	if root.Left == nil && root.Right == nil {
		return 1
	}
	
	minD := int(^uint(0) >> 1) // 最大整数值
	
	if root.Left != nil {
		minD = minInt(minD, MinDepth(root.Left))
	}
	
	if root.Right != nil {
		minD = minInt(minD, MinDepth(root.Right))
	}
	
	return minD + 1
}

// minInt 返回两个整数中的较小值
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// SimplifyPath 简化路径
// 给你一个字符串 path，表示指向某一文件或目录的 Unix 风格 绝对路径
func SimplifyPath(path string) string {
	stack := make([]string, 0)
	components := splitPath(path)
	
	for _, component := range components {
		if component == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if component != "." && component != "" {
			stack = append(stack, component)
		}
	}
	
	if len(stack) == 0 {
		return "/"
	}
	
	result := ""
	for _, dir := range stack {
		result += "/" + dir
	}
	
	return result
}

// splitPath 分割路径
func splitPath(path string) []string {
	var components []string
	current := ""
	
	for _, char := range path {
		if char == '/' {
			if current != "" {
				components = append(components, current)
				current = ""
			}
		} else {
			current += string(char)
		}
	}
	
	if current != "" {
		components = append(components, current)
	}
	
	return components
}