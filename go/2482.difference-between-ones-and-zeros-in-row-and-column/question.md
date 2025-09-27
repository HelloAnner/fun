# [2482. 行和列中一和零的差值][link] (Medium)

[link]: https://leetcode.cn/problems/difference-between-ones-and-zeros-in-row-and-column/

给你一个下标从 **0** 开始的 `m x n` 二进制矩阵 `grid` 。

我们按照如下过程，定义一个下标从 **0** 开始的 `m x n` 差值矩阵 `diff` ：

- 令第 `i` 行一的数目为 `onesRowᵢ` 。
- 令第 `j` 列一的数目为 `onesColⱼ`。
- 令第 `i` 行零的数目为 `zerosRowᵢ` 。
- 令第 `j` 列零的数目为 `zerosColⱼ` 。
- `diff[i][j] = onesRowᵢ + onesColⱼ - zerosRowᵢ - zerosColⱼ`

请你返回差值矩阵 `diff` 。

**示例 1：**

![](https://assets.leetcode.com/uploads/2022/11/06/image-20221106171729-5.png)

```
输入：grid = [[0,1,1],[1,0,1],[0,0,1]]
输出：[[0,0,4],[0,0,4],[-2,-2,2]]
解释：
- diff[0][0] = onesRow₀ + onesCol₀ - zerosRow₀ - zerosCol₀ = 2 + 1 - 1 - 2 = 0
- diff[0][1] = onesRow₀ + onesCol₁ - zerosRow₀ - zerosCol₁ = 2 + 1 - 1 - 2 = 0
- diff[0][2] = onesRow₀ + onesCol₂ - zerosRow₀ - zerosCol₂ = 2 + 3 - 1 - 0 = 4
- diff[1][0] = onesRow₁ + onesCol₀ - zerosRow₁ - zerosCol₀ = 2 + 1 - 1 - 2 = 0
- diff[1][1] = onesRow₁ + onesCol₁ - zerosRow₁ - zerosCol₁ = 2 + 1 - 1 - 2 = 0
- diff[1][2] = onesRow₁ + onesCol₂ - zerosRow₁ - zerosCol₂ = 2 + 3 - 1 - 0 = 4
- diff[2][0] = onesRow₂ + onesCol₀ - zerosRow₂ - zerosCol₀ = 1 + 1 - 2 - 2 = -2
- diff[2][1] = onesRow₂ + onesCol₁ - zerosRow₂ - zerosCol₁ = 1 + 1 - 2 - 2 = -2
- diff[2][2] = onesRow₂ + onesCol₂ - zerosRow₂ - zerosCol₂ = 1 + 3 - 2 - 0 = 2
```

**示例 2：**

![](https://assets.leetcode.com/uploads/2022/11/06/image-20221106171747-6.png)

```
输入：grid = [[1,1,1],[1,1,1]]
输出：[[5,5,5],[5,5,5]]
解释：
- diff[0][0] = onesRow₀ + onesCol₀ - zerosRow₀ - zerosCol₀ = 3 + 2 - 0 - 0 = 5
- diff[0][1] = onesRow₀ + onesCol₁ - zerosRow₀ - zerosCol₁ = 3 + 2 - 0 - 0 = 5
- diff[0][2] = onesRow₀ + onesCol₂ - zerosRow₀ - zerosCol₂ = 3 + 2 - 0 - 0 = 5
- diff[1][0] = onesRow₁ + onesCol₀ - zerosRow₁ - zerosCol₀ = 3 + 2 - 0 - 0 = 5
- diff[1][1] = onesRow₁ + onesCol₁ - zerosRow₁ - zerosCol₁ = 3 + 2 - 0 - 0 = 5
- diff[1][2] = onesRow₁ + onesCol₂ - zerosRow₁ - zerosCol₂ = 3 + 2 - 0 - 0 = 5
```

**提示：**

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 10⁵`
- `1 <= m * n <= 10⁵`
- `grid[i][j]` 要么是 `0` ，要么是 `1` 。
