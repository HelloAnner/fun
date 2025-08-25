package algorithms

/**
 *  Leetcode 1504 统计全 1 子矩形
 *  给你一个 m x n 的二进制矩阵 mat ，请你返回有多少个 子矩形 的元素全部都是 1 。
	https://leetcode.cn/problems/count-submatrices-with-all-ones/description/?envType=daily-question&envId=2025-08-21
*/

func numSubmat(mat [][]int) (ans int) {
	m, n := len(mat), len(mat[0])
	for top := 0; top < m; top++ { // 枚举上边界
		a := make([]int, n)
		for bottom := top; bottom < m; bottom++ { // 枚举下边界
			h := bottom - top + 1 // 高
			// 2348. 全 h 子数组的数目
			last := -1
			for j := 0; j < n; j++ {
				a[j] += mat[bottom][j] // 把 bottom 这一行的值加到 a 中
				if a[j] != h {
					last = j // 记录上一个非 h 元素的位置
				} else {
					ans += j - last
				}
			}
		}
	}
	return
}
