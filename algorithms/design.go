package algorithms

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * 380. O(1) 时间插入、删除和获取随机元素
 * Insert Delete GetRandom O(1)
 */
type RandomizedSet struct {
	nums    []int
	indices map[int]int
	rng     *rand.Rand
}

func NewRandomizedSet() *RandomizedSet {
	return &RandomizedSet{
		nums:    []int{},
		indices: make(map[int]int),
		rng:     rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, exists := rs.indices[val]; exists {
		return false
	}
	rs.indices[val] = len(rs.nums)
	rs.nums = append(rs.nums, val)
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	idx, exists := rs.indices[val]
	if !exists {
		return false
	}
	
	// 将最后一个元素移到要删除的位置
	lastElement := rs.nums[len(rs.nums)-1]
	rs.nums[idx] = lastElement
	rs.indices[lastElement] = idx
	
	// 删除最后一个元素
	rs.nums = rs.nums[:len(rs.nums)-1]
	delete(rs.indices, val)
	
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	if len(rs.nums) == 0 {
		return 0
	}
	return rs.nums[rs.rng.Intn(len(rs.nums))]
}

/**
 * 705. 设计哈希集合
 * Design HashSet
 */
type MyHashSet struct {
	size int
	data [][]int
}

func NewMyHashSet() *MyHashSet {
	size := 1000
	data := make([][]int, size)
	for i := range data {
		data[i] = []int{}
	}
	return &MyHashSet{
		size: size,
		data: data,
	}
}

func (hs *MyHashSet) hash(key int) int {
	return key % hs.size
}

func (hs *MyHashSet) Add(key int) {
	if hs.Contains(key) {
		return
	}
	idx := hs.hash(key)
	hs.data[idx] = append(hs.data[idx], key)
}

func (hs *MyHashSet) Remove(key int) {
	if !hs.Contains(key) {
		return
	}
	idx := hs.hash(key)
	for i, v := range hs.data[idx] {
		if v == key {
			// 删除元素
			hs.data[idx] = append(hs.data[idx][:i], hs.data[idx][i+1:]...)
			break
		}
	}
}

func (hs *MyHashSet) Contains(key int) bool {
	idx := hs.hash(key)
	for _, v := range hs.data[idx] {
		if v == key {
			return true
		}
	}
	return false
}

// DemonstrateDesignAlgorithms 演示设计算法
func DemonstrateDesignAlgorithms() {
	fmt.Println("=== 设计算法演示 ===")
	
	// RandomizedSet 演示
	fmt.Println("RandomizedSet 演示:")
	rs := NewRandomizedSet()
	fmt.Printf("插入 1: %t\n", rs.Insert(1))
	fmt.Printf("删除 2: %t\n", rs.Remove(2))
	fmt.Printf("插入 2: %t\n", rs.Insert(2))
	fmt.Printf("获取随机元素: %d\n", rs.GetRandom())
	fmt.Printf("删除 1: %t\n", rs.Remove(1))
	fmt.Printf("插入 2: %t\n", rs.Insert(2))
	fmt.Printf("获取随机元素: %d\n", rs.GetRandom())
	
	// MyHashSet 演示
	fmt.Println("\nMyHashSet 演示:")
	hs := NewMyHashSet()
	hs.Add(1)
	hs.Add(2)
	fmt.Printf("包含 1: %t\n", hs.Contains(1))
	fmt.Printf("包含 3: %t\n", hs.Contains(3))
	hs.Add(2)
	fmt.Printf("包含 2: %t\n", hs.Contains(2))
	hs.Remove(2)
	fmt.Printf("删除后包含 2: %t\n", hs.Contains(2))
	
	// IsValidNumber 演示
	fmt.Println("\n数字验证演示:")
	testNumbers := []string{"0", "e", ".", "2e10", "-90e3", "1e", "e3", "6e-1", "99e2.5", "53.5e93", "--6", "-+3", "95a54e53"}
	for _, num := range testNumbers {
		fmt.Printf("'%s' 是有效数字: %t\n", num, IsValidNumber(num))
	}
	fmt.Println()
}

/**
 * 65. 有效数字
 * Valid Number
 */
func IsValidNumber(s string) bool {
	// 去除首尾空格
	start, end := 0, len(s)-1
	for start <= end && s[start] == ' ' {
		start++
	}
	for start <= end && s[end] == ' ' {
		end--
	}
	
	if start > end {
		return false
	}
	
	s = s[start : end+1]
	
	seenDigit := false
	seenDot := false
	seenE := false
	
	for i, char := range s {
		if char >= '0' && char <= '9' {
			seenDigit = true
		} else if char == '.' {
			if seenDot || seenE {
				return false
			}
			seenDot = true
		} else if char == 'e' || char == 'E' {
			if seenE || !seenDigit {
				return false
			}
			seenE = true
			// e 后面必须出现数字
			seenDigit = false
		} else if char == '+' || char == '-' {
			if i > 0 && s[i-1] != 'e' && s[i-1] != 'E' {
				return false
			}
		} else {
			return false
		}
	}
	
	return seenDigit
}