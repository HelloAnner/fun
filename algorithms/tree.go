package algorithms

import (
	"algorithms-go/data_structures"
)

type TreeNode = data_structures.TreeNode

// 树相关算法实现
// Tree related algorithms

// 101. 对称二叉树
// Symmetric Tree
// 判断二叉树是否镜像对称
// https://leetcode.cn/problems/symmetric-tree/description/
func IsSymmetric(root *data_structures.TreeNode) bool {
	if root == nil {
		return true
	}
	return isSameTree(root.Left, root.Right)
}

func isSameTree(p, q *data_structures.TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val &&
		isSameTree(p.Left, q.Right) &&
		isSameTree(p.Right, q.Left)
}

// 104. 二叉树的最大深度
// Maximum Depth of Binary Tree
// 计算二叉树的最大深度（根节点到最远叶子节点的最长路径上的节点数）
// https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/
func MaxDepth(root *data_structures.TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := MaxDepth(root.Left)
	rightDepth := MaxDepth(root.Right)
	if leftDepth > rightDepth {
		return 1 + leftDepth
	}
	return 1 + rightDepth
}

// 226. 翻转二叉树
// Invert Binary Tree
// 翻转二叉树的左右子树
// https://leetcode.cn/problems/invert-binary-tree/description/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

// 102. 二叉树的层序遍历
// Binary Tree Level Order Traversal
// 按层次遍历二叉树，返回每层的节点值
// https://leetcode.cn/problems/binary-tree-level-order-traversal/description/
func LevelOrder(root *data_structures.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*data_structures.TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := []int{}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			currentLevel = append(currentLevel, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, currentLevel)
	}

	return result
}

// 112. 路径总和
// Path Sum
// 判断二叉树中是否存在根节点到叶子节点的路径，其节点值之和等于目标值
// https://leetcode.cn/problems/path-sum/description/
func HasPathSum(root *data_structures.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	return HasPathSum(root.Left, targetSum-root.Val) ||
		HasPathSum(root.Right, targetSum-root.Val)
}

// 543. 二叉树的直径
// Diameter of Binary Tree
// 计算二叉树的直径（任意两个节点间最长路径的长度）
// https://leetcode.cn/problems/diameter-of-binary-tree/description/
func DiameterOfBinaryTree(root *data_structures.TreeNode) int {
	maxDiameter := 0

	var depth func(*data_structures.TreeNode) int
	depth = func(node *data_structures.TreeNode) int {
		if node == nil {
			return 0
		}

		leftDepth := depth(node.Left)
		rightDepth := depth(node.Right)

		if leftDepth+rightDepth > maxDiameter {
			maxDiameter = leftDepth + rightDepth
		}

		if leftDepth > rightDepth {
			return 1 + leftDepth
		}
		return 1 + rightDepth
	}

	depth(root)
	return maxDiameter
}

// 236. 二叉树的最近公共祖先
// Lowest Common Ancestor of a Binary Tree
// 找到二叉树中两个节点的最近公共祖先
// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/description/
func LowestCommonAncestor(root, p, q *data_structures.TreeNode) *data_structures.TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	return right
}

// 查找二叉树最深的一个节点
// Find Deepest Node in Binary Tree
// 找到二叉树中最深的一个叶子节点
func findDeepestNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var deepestNode *TreeNode
	maxDepth := -1

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth > maxDepth {
			maxDepth = depth
			deepestNode = node
		}
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, 0)
	return deepestNode
}
