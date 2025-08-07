/**
 * 树相关算法实现
 * Tree related algorithms
 */

import { TreeNode } from '../data-structures/tree-node';

/**
 * 101. 对称二叉树
 * Symmetric Tree
 */
export function isSymmetric(root: TreeNode | null): boolean {
    if (!root) return true;
    return isSameTree(root.left, root.right);
}

function isSameTree(p: TreeNode | null, q: TreeNode | null): boolean {
    if (!p || !q) {
        return p === q;
    }
    return (
        p.val === q.val &&
        isSameTree(p.left, q.right) &&
        isSameTree(p.right, q.left)
    );
}

/**
 * 104. 二叉树的最大深度
 * Maximum Depth of Binary Tree
 */
export function maxDepth(root: TreeNode | null): number {
    if (!root) return 0;
    return 1 + Math.max(maxDepth(root.left), maxDepth(root.right));
}

/**
 * 226. 翻转二叉树
 * Invert Binary Tree
 */
export function invertTree(root: TreeNode | null): TreeNode | null {
    if (!root) return null;
    
    const temp = root.left;
    root.left = root.right;
    root.right = temp;
    
    invertTree(root.left);
    invertTree(root.right);
    
    return root;
}

/**
 * 102. 二叉树的层序遍历
 * Binary Tree Level Order Traversal
 */
export function levelOrder(root: TreeNode | null): number[][] {
    if (!root) return [];
    
    const result: number[][] = [];
    const queue: TreeNode[] = [root];
    
    while (queue.length > 0) {
        const levelSize = queue.length;
        const currentLevel: number[] = [];
        
        for (let i = 0; i < levelSize; i++) {
            const node = queue.shift()!;
            currentLevel.push(node.val);
            
            if (node.left) queue.push(node.left);
            if (node.right) queue.push(node.right);
        }
        
        result.push(currentLevel);
    }
    
    return result;
}

/**
 * 112. 路径总和
 * Path Sum
 */
export function hasPathSum(root: TreeNode | null, targetSum: number): boolean {
    if (!root) return false;
    
    if (!root.left && !root.right) {
        return root.val === targetSum;
    }
    
    return (
        hasPathSum(root.left, targetSum - root.val) ||
        hasPathSum(root.right, targetSum - root.val)
    );
}

/**
 * 543. 二叉树的直径
 * Diameter of Binary Tree
 */
export function diameterOfBinaryTree(root: TreeNode | null): number {
    let maxDiameter = 0;
    
    function depth(node: TreeNode | null): number {
        if (!node) return 0;
        
        const leftDepth = depth(node.left);
        const rightDepth = depth(node.right);
        
        maxDiameter = Math.max(maxDiameter, leftDepth + rightDepth);
        
        return 1 + Math.max(leftDepth, rightDepth);
    }
    
    depth(root);
    return maxDiameter;
}

/**
 * 236. 二叉树的最近公共祖先
 * Lowest Common Ancestor of a Binary Tree
 */
export function lowestCommonAncestor(
    root: TreeNode | null,
    p: TreeNode | null,
    q: TreeNode | null
): TreeNode | null {
    if (!root || root === p || root === q) {
        return root;
    }
    
    const left = lowestCommonAncestor(root.left, p, q);
    const right = lowestCommonAncestor(root.right, p, q);
    
    if (left && right) {
        return root;
    }
    
    return left || right;
}