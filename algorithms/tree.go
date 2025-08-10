package algorithms

import (
	"algorithms-go/data_structures"
)

/**
 * 树相关算法实现
 * Tree related algorithms
 */

/**
 * 101. 对称二叉树
 * Symmetric Tree
 */
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

/**
 * 104. 二叉树的最大深度
 * Maximum Depth of Binary Tree
 */
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

/**
 * 226. 翻转二叉树
 * Invert Binary Tree
 */
func InvertTree(root *data_structures.TreeNode) *data_structures.TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	InvertTree(root.Left)
	InvertTree(root.Right)

	return root
}

/**
 * 102. 二叉树的层序遍历
 * Binary Tree Level Order Traversal
 */
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

/**
 * 112. 路径总和
 * Path Sum
 */
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

/**
 * 543. 二叉树的直径
 * Diameter of Binary Tree
 */
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

/**
 * 236. 二叉树的最近公共祖先
 * Lowest Common Ancestor of a Binary Tree
 */
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