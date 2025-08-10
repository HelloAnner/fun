package data_structures

import "errors"

// BitArray 位数组实现
// 一个空间高效的数组，每个索引位置存储 0 或 1
type BitArray struct {
	size  int
	array []int32
}

// NewBitArray 初始化指定大小的位数组
func NewBitArray(size int) *BitArray {
	if size <= 0 {
		return nil
	}
	
	// 每个 int32 可以存储 32 位
	arraySize := (size + 31) / 32
	return &BitArray{
		size:  size,
		array: make([]int32, arraySize),
	}
}

// Set 设置指定索引位置的值（0 或 1）
func (ba *BitArray) Set(i int, val int) error {
	if i < 0 || i >= ba.size {
		return errors.New("index out of range")
	}
	if val != 0 && val != 1 {
		return errors.New("value must be either 0 or 1")
	}
	
	idx := i / 32
	bitOffset := i % 32
	mask := int32(1 << bitOffset)
	
	if val == 1 {
		ba.array[idx] |= mask
	} else {
		ba.array[idx] &= ^mask
	}
	
	return nil
}

// Get 获取指定索引位置的值
func (ba *BitArray) Get(i int) (int, error) {
	if i < 0 || i >= ba.size {
		return 0, errors.New("index out of range")
	}
	
	idx := i / 32
	bitOffset := i % 32
	mask := int32(1 << bitOffset)
	
	if ba.array[idx]&mask != 0 {
		return 1, nil
	}
	return 0, nil
}

// Size 返回位数组的大小
func (ba *BitArray) Size() int {
	return ba.size
}

// Clear 清空所有位
func (ba *BitArray) Clear() {
	for i := range ba.array {
		ba.array[i] = 0
	}
}

// Count 统计值为 1 的位数
func (ba *BitArray) Count() int {
	count := 0
	for i := 0; i < ba.size; i++ {
		if val, _ := ba.Get(i); val == 1 {
			count++
		}
	}
	return count
}

// String 返回位数组的字符串表示
func (ba *BitArray) String() string {
	result := make([]byte, ba.size)
	for i := 0; i < ba.size; i++ {
		if val, _ := ba.Get(i); val == 1 {
			result[i] = '1'
		} else {
			result[i] = '0'
		}
	}
	return string(result)
}