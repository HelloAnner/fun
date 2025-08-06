/**
 * 树算法测试
 * Tree algorithms tests
 */

import { TreeNode } from '../../data-structures/tree-node';
import { isSymmetric, maxDepth, invertTree, levelOrder, hasPathSum, diameterOfBinaryTree } from '../tree';

describe('Tree Algorithms', () => {
    describe('isSymmetric', () => {
        test('should check if tree is symmetric', () => {
            // [1,2,2,3,4,4,3]
            const root1 = new TreeNode(1);
            root1.left = new TreeNode(2);
            root1.right = new TreeNode(2);
            root1.left.left = new TreeNode(3);
            root1.left.right = new TreeNode(4);
            root1.right.left = new TreeNode(4);
            root1.right.right = new TreeNode(3);
            expect(isSymmetric(root1)).toBe(true);

            // [1,2,2,null,3,null,3]
            const root2 = new TreeNode(1);
            root2.left = new TreeNode(2);
            root2.right = new TreeNode(2);
            root2.left.right = new TreeNode(3);
            root2.right.right = new TreeNode(3);
            expect(isSymmetric(root2)).toBe(false);
        });
    });

    describe('maxDepth', () => {
        test('should calculate maximum depth', () => {
            // [3,9,20,null,null,15,7]
            const root = new TreeNode(3);
            root.left = new TreeNode(9);
            root.right = new TreeNode(20);
            root.right.left = new TreeNode(15);
            root.right.right = new TreeNode(7);
            expect(maxDepth(root)).toBe(3);

            // [1,null,2]
            const root2 = new TreeNode(1);
            root2.right = new TreeNode(2);
            expect(maxDepth(root2)).toBe(2);
        });
    });

    describe('levelOrder', () => {
        test('should return level order traversal', () => {
            // [3,9,20,null,null,15,7]
            const root = new TreeNode(3);
            root.left = new TreeNode(9);
            root.right = new TreeNode(20);
            root.right.left = new TreeNode(15);
            root.right.right = new TreeNode(7);
            expect(levelOrder(root)).toEqual([[3], [9, 20], [15, 7]]);
        });
    });

    describe('hasPathSum', () => {
        test('should check if path sum exists', () => {
            // [5,4,8,11,null,13,4,7,2,null,null,null,1]
            const root = new TreeNode(5);
            root.left = new TreeNode(4);
            root.right = new TreeNode(8);
            root.left.left = new TreeNode(11);
            root.left.left.left = new TreeNode(7);
            root.left.left.right = new TreeNode(2);
            root.right.left = new TreeNode(13);
            root.right.right = new TreeNode(4);
            root.right.right.right = new TreeNode(1);
            
            expect(hasPathSum(root, 22)).toBe(true);
            expect(hasPathSum(root, 5)).toBe(false);
        });
    });

    describe('diameterOfBinaryTree', () => {
        test('should calculate diameter of binary tree', () => {
            // [1,2,3,4,5]
            const root = new TreeNode(1);
            root.left = new TreeNode(2);
            root.right = new TreeNode(3);
            root.left.left = new TreeNode(4);
            root.left.right = new TreeNode(5);
            expect(diameterOfBinaryTree(root)).toBe(3);
        });
    });
});