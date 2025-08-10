package data_structures

/**
 * 二叉树节点定义
 * Definition for a binary tree node
 */
type TreeNode struct {
	Val   int       `json:"val"`
	Left  *TreeNode `json:"left"`
	Right *TreeNode `json:"right"`
}

// NewTreeNode 创建新的二叉树节点
func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

// NewTreeNodeWithChildren 创建带子节点的二叉树节点
func NewTreeNodeWithChildren(val int, left, right *TreeNode) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  left,
		Right: right,
	}
}