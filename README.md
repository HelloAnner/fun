# TypeScript 算法实现

这是一个TypeScript的项目，包含了各种经典算法和数据结构的实现。

## 项目结构

```
ts-src/
├── data-structures/          # 数据结构定义
│   ├── tree-node.ts         # 二叉树节点
│   ├── list-node.ts         # 链表节点
│   ├── graph-node.ts        # 图节点
│   └── index.ts             # 导出文件
├── algorithms/               # 算法实现
│   ├── stack.ts             # 栈相关算法
│   ├── hash.ts              # 哈希表相关算法
│   ├── two-pointers.ts      # 双指针算法
│   ├── tree.ts              # 树相关算法
│   ├── graph.ts             # 图相关算法
│   ├── dynamic-programming.ts # 动态规划算法
│   ├── greedy.ts            # 贪心算法
│   ├── binary-search.ts     # 二分查找算法
│   ├── sorting.ts           # 排序算法
│   ├── string.ts            # 字符串算法
│   ├── simulation.ts        # 模拟算法
│   ├── __tests__/           # 测试文件
│   └── index.ts             # 导出文件
└── index.ts                 # 主入口文件
```

## 安装依赖

```bash
npm install
```

## 运行和测试

```bash
# 编译TypeScript
npm run build

# 运行开发模式
npm run dev

# 运行测试
npm test

# 监听模式运行测试
npm run test:watch

# 清理编译文件
npm clean
```

## 已实现的算法

### 栈相关算法
- ✅ 150. 逆波兰表达式求值 (evalRPN)
- ✅ 最小栈实现 (MinStack)
- ✅ 有效的括号 (isValid)

### 哈希表相关算法
- ✅ 2766. 重新放置石块 (relocateMarbles)
- ✅ 1. 两数之和 (twoSum)
- ✅ 202. 快乐数 (isHappy)
- ✅ 997. 找到小镇的法官 (findJudge)
- ✅ 1436. 旅行终点站 (destCity)

### 双指针算法
- ✅ 977. 有序数组的平方 (sortedSquares)
- ✅ 88. 合并两个有序数组 (merge)
- ✅ 125. 验证回文串 (isPalindrome)
- ✅ 392. 判断子序列 (isSubsequence)
- ✅ 151. 反转字符串中的单词 (reverseWords)

### 树相关算法
- ✅ 101. 对称二叉树 (isSymmetric)
- ✅ 104. 二叉树的最大深度 (maxDepth)
- ✅ 226. 翻转二叉树 (invertTree)
- ✅ 102. 二叉树的层序遍历 (levelOrder)
- ✅ 112. 路径总和 (hasPathSum)
- ✅ 543. 二叉树的直径 (diameterOfBinaryTree)
- ✅ 236. 二叉树的最近公共祖先 (lowestCommonAncestor)

### 图相关算法
- ✅ 133. 克隆图 (cloneGraph)
- ✅ 752. 打开转盘锁 (openLock)
- ✅ 547. 省份数量 (findCircleNum)
- ✅ 207. 课程表 (canFinish)
- ✅ 1971. 寻找图中是否存在路径 (validPath)

### 动态规划算法
- ✅ 70. 爬楼梯 (climbStairs)
- ✅ 322. 零钱兑换 (coinChange)
- ✅ 300. 最长递增子序列 (lengthOfLIS)
- ✅ 198. 打家劫舍 (rob)
- ✅ 139. 单词拆分 (wordBreak)
- ✅ 1035. 不相交的线 (maxUncrossedLines)
- ✅ 1216. 验证回文字符串 III (isValidPalindrome)

### 贪心算法
- ✅ 55. 跳跃游戏 (canJump)
- ✅ 45. 跳跃游戏 II (jump)
- ✅ 121. 买卖股票的最佳时机 (maxProfit)
- ✅ 135. 分发糖果 (candy)
- ✅ 2740. 找出分区值 (findValueOfPartition)
- ✅ 68. 文本左右对齐 (fullJustify)

### 二分查找算法
- ✅ 704. 二分查找 (search)
- ✅ 69. x 的平方根 (mySqrt)
- ✅ 35. 搜索插入位置 (searchInsert)
- ✅ 34. 在排序数组中查找元素的第一个和最后一个位置 (searchRange)
- ✅ 3143. 正方形中的最多点数 (maxPointsInsideSquare)

### 排序算法
- ✅ 快速排序 (quickSort)
- ✅ 归并排序 (mergeSort)
- ✅ 堆排序 (heapSort)
- ✅ 冒泡排序 (bubbleSort)
- ✅ 插入排序 (insertionSort)
- ✅ 选择排序 (selectionSort)
- ✅ 计数排序 (countingSort)

### 字符串算法
- ✅ 28. 找出字符串中第一个匹配项的下标 (strStr)
- ✅ 6. Z 字形变换 (convert)
- ✅ 65. 有效数字 (isNumber)
- ✅ 字符串翻转检查 (canShift)
- ✅ 字母异位词分组 (groupAnagrams)
- ✅ 205. 同构字符串 (isIsomorphic)
- ✅ 242. 有效的字母异位词 (isAnagram)
- ✅ 290. 单词规律 (wordPattern)

### 模拟算法
- ✅ 36. 有效的数独 (isValidSudoku)
- ✅ 48. 旋转图像 (rotate)
- ✅ 54. 螺旋矩阵 (spiralOrder)
- ✅ 73. 矩阵置零 (setZeroes)
- ✅ 274. H 指数 (hIndex)
- ✅ 1103. 分糖果 II (distributeCandies)
- ✅ 2974. 最小数字游戏 (numberGame)
- ✅ 724. 寻找数组的中心下标 (pivotIndex)
- ✅ 1329. 将矩阵按对角线排序 (diagonalSort)

## 使用示例

```typescript
import { evalRPN, twoSum, isSymmetric, TreeNode } from './ts-src';

// 使用逆波兰表达式求值
const result = evalRPN(["2", "1", "+", "3", "*"]);
console.log(result); // 9

// 使用两数之和
const indices = twoSum([2, 7, 11, 15], 9);
console.log(indices); // [0, 1]

// 使用对称二叉树检查
const root = new TreeNode(1);
root.left = new TreeNode(2);
root.right = new TreeNode(2);
console.log(isSymmetric(root)); // true
```

## 特性

- 🚀 完整的TypeScript类型支持
- 🧪 全面的单元测试覆盖
- 📚 详细的代码注释和文档
- 🔧 现代化的开发工具配置
- 📦 模块化的代码组织结构
- ⚡ 高性能的算法实现

## 贡献

欢迎提交Issue和Pull Request来改进这个项目！

## 许可证

MIT License