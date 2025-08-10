package engineering

import (
	"math"
)

/**
 * 布隆过滤器实现
 * Bloom Filter implementation
 */

type BloomFilter struct {
	// m - 位数组大小
	size int

	// k - 哈希函数数量
	hashCount int

	// 底层原理是字节（Byte），而不是位（Bit）
	// 存在巨大的空间浪费
	// 空间效率只有理论值的 1/8，造成了 8 倍的空间浪费
	// 数组的长度是 size / 8，因为每个元素可以存 8 个位（0 或 1）
	bitArray []byte
}

/**
 * 构造一个高效的布隆过滤器
 * @param estimatedItemCount 预估要存储的元素数量 (n)
 * @param falsePositiveRate 期望的假阳性概率 (p)
 */
func NewBloomFilter(estimatedItemCount int, falsePositiveRate float64) *BloomFilter {
	if estimatedItemCount <= 0 {
		estimatedItemCount = 1000
	}
	if falsePositiveRate <= 0 {
		falsePositiveRate = 0.01
	}

	size, hashCount := getOptimalParams(estimatedItemCount, falsePositiveRate)
	byteArraySize := int(math.Ceil(float64(size) / 8))

	return &BloomFilter{
		size:      size,
		hashCount: hashCount,
		bitArray:  make([]byte, byteArraySize),
	}
}

/**
 * 向布隆过滤器中添加一个元素
 * @param item 要添加的元素
 */
func (bf *BloomFilter) Add(item string) {
	hashes := bf.getHashes(item)
	for _, hash := range hashes {
		byteIndex := hash / 8
		bitIndex := hash % 8
		// 使用按位或 "|" 操作将该位置为 1
		bf.bitArray[byteIndex] |= (1 << bitIndex)
	}
}

/**
 * 检查布隆过滤器是否可能包含该元素
 * @param item 要检查的元素
 * @returns 如果可能包含，则返回 true；否则返回 false
 */
func (bf *BloomFilter) MightContain(item string) bool {
	hashes := bf.getHashes(item)
	for _, hash := range hashes {
		byteIndex := hash / 8
		bitIndex := hash % 8
		// 使用按位与 "&" 操作检查该位是否为 1
		if (bf.bitArray[byteIndex] & (1 << bitIndex)) == 0 {
			return false
		}
	}
	return true
}

/**
 * 私有辅助方法：获取一个元素对应的 k 个哈希值（即位数组的索引）
 * 这里使用"双重哈希"技巧，仅用两个基础哈希函数来模拟 k 个哈希函数。
 * g_i(x) = (h1(x) + i * h2(x)) % m
 */
func (bf *BloomFilter) getHashes(item string) []int {
	hashes := make([]int, bf.hashCount)
	h1 := djb2(item)
	h2 := sdbm(item)

	for i := 0; i < bf.hashCount; i++ {
		hash := abs((h1 + i*h2) % bf.size)
		hashes[i] = hash
	}
	return hashes
}

// --- 基础哈希函数 ---
// 这些函数只需要速度快、分布好即可，不需要具备加密安全性。

/**
 * 一个简单高效的字符串哈希函数 (djb2)
 */
func djb2(str string) int {
	hash := 5381
	for i := 0; i < len(str); i++ {
		hash = (hash * 33) ^ int(str[i])
	}
	return hash
}

/**
 * 一个简单的哈希函数 (sdbm)
 */
func sdbm(str string) int {
	hash := 0
	for i := 0; i < len(str); i++ {
		hash = int(str[i]) + (hash << 6) + (hash << 16) - hash
	}
	return hash
}

/**
 * 根据预估元素数量和期望误判率，计算最优的 m 和 k
 */
func getOptimalParams(n int, p float64) (int, int) {
	m := int(math.Ceil(-(float64(n) * math.Log(p)) / (math.Log(2) * math.Log(2))))
	k := int(math.Round((float64(m) / float64(n)) * math.Log(2)))
	return m, k
}

/**
 * 辅助函数：计算绝对值
 */
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
