/**
 * 算法演示文件
 * Algorithm demonstration
 */

import { TreeNode } from './data-structures/tree-node';
import { GraphNode } from './data-structures/graph-node';
import { 
    evalRPN, 
    MinStack, 
    isValid 
} from './leetcode/stack';
import { 
    twoSum, 
    relocateMarbles, 
    isHappy 
} from './leetcode/hash';
import { 
    sortedSquares, 
    isPalindrome, 
    merge 
} from './leetcode/two-pointers';
import { 
    isSymmetric, 
    maxDepth, 
    levelOrder 
} from './leetcode/tree';
import { 
    cloneGraph, 
    openLock 
} from './leetcode/graph';
import { 
    climbStairs, 
    coinChange, 
    rob 
} from './leetcode/dynamic-programming';
import { 
    canJump, 
    maxProfit 
} from './leetcode/greedy';
import { 
    search, 
    mySqrt 
} from './leetcode/binary-search';
import { 
    quickSort, 
    mergeSort 
} from './leetcode/sorting';
import { 
    strStr, 
    convert, 
    isAnagram 
} from './leetcode/string';
import { 
    isValidSudoku, 
    spiralOrder, 
    pivotIndex 
} from './leetcode/simulation';

console.log('🚀 TypeScript 算法演示开始\n');

// 栈算法演示
console.log('📚 栈算法演示:');
console.log('逆波兰表达式 ["2","1","+","3","*"]:', evalRPN(["2", "1", "+", "3", "*"]));

const minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
console.log('最小栈最小值:', minStack.getMin());

console.log('有效括号 "()[]{}": ', isValid("()[]{}"));
console.log('');

// 哈希算法演示
console.log('🔍 哈希算法演示:');
console.log('两数之和 [2,7,11,15], target=9:', twoSum([2, 7, 11, 15], 9));
console.log('重新放置石块:', relocateMarbles([1, 6, 7, 8], [1, 7, 2], [2, 9, 5]));
console.log('快乐数 19:', isHappy(19));
console.log('');

// 双指针算法演示
console.log('👆 双指针算法演示:');
console.log('有序数组的平方 [-4,-1,0,3,10]:', sortedSquares([-4, -1, 0, 3, 10]));
console.log('验证回文串 "A man a plan a canal Panama":', isPalindrome("A man a plan a canal Panama"));
console.log('');

// 树算法演示
console.log('🌳 树算法演示:');
const treeRoot = new TreeNode(1);
treeRoot.left = new TreeNode(2);
treeRoot.right = new TreeNode(2);
treeRoot.left.left = new TreeNode(3);
treeRoot.left.right = new TreeNode(4);
treeRoot.right.left = new TreeNode(4);
treeRoot.right.right = new TreeNode(3);

console.log('对称二叉树检查:', isSymmetric(treeRoot));

const treeRoot2 = new TreeNode(3);
treeRoot2.left = new TreeNode(9);
treeRoot2.right = new TreeNode(20);
treeRoot2.right.left = new TreeNode(15);
treeRoot2.right.right = new TreeNode(7);

console.log('二叉树最大深度:', maxDepth(treeRoot2));
console.log('层序遍历:', levelOrder(treeRoot2));
console.log('');

// 图算法演示
console.log('🕸️ 图算法演示:');
const graphNode1 = new GraphNode(1);
const graphNode2 = new GraphNode(2);
const graphNode3 = new GraphNode(3);
const graphNode4 = new GraphNode(4);

graphNode1.neighbors = [graphNode2, graphNode4];
graphNode2.neighbors = [graphNode1, graphNode3];
graphNode3.neighbors = [graphNode2, graphNode4];
graphNode4.neighbors = [graphNode1, graphNode3];

const clonedGraph = cloneGraph(graphNode1);
console.log('图克隆成功，克隆节点值:', clonedGraph?.val);

console.log('打开转盘锁步数:', openLock(["0201", "0101", "0102", "1212", "2002"], "0202"));
console.log('');

// 动态规划演示
console.log('🎯 动态规划演示:');
console.log('爬楼梯 n=5:', climbStairs(5));
console.log('零钱兑换 coins=[1,3,4], amount=6:', coinChange([1, 3, 4], 6));
console.log('打家劫舍 [2,7,9,3,1]:', rob([2, 7, 9, 3, 1]));
console.log('');

// 贪心算法演示
console.log('💰 贪心算法演示:');
console.log('跳跃游戏 [2,3,1,1,4]:', canJump([2, 3, 1, 1, 4]));
console.log('买卖股票最大利润 [7,1,5,3,6,4]:', maxProfit([7, 1, 5, 3, 6, 4]));
console.log('');

// 二分查找演示
console.log('🔍 二分查找演示:');
console.log('二分查找 nums=[-1,0,3,5,9,12], target=9:', search([-1, 0, 3, 5, 9, 12], 9));
console.log('平方根 x=8:', mySqrt(8));
console.log('');

// 排序算法演示
console.log('📊 排序算法演示:');
const unsortedArray = [64, 34, 25, 12, 22, 11, 90];
console.log('原数组:', unsortedArray);
console.log('快速排序:', quickSort([...unsortedArray]));
console.log('归并排序:', mergeSort([...unsortedArray]));
console.log('');

// 字符串算法演示
console.log('📝 字符串算法演示:');
console.log('字符串匹配 "hello" in "hello world":', strStr("hello world", "hello"));
console.log('Z字形变换 "PAYPALISHIRING", numRows=3:', convert("PAYPALISHIRING", 3));
console.log('字母异位词检查 "anagram" vs "nagaram":', isAnagram("anagram", "nagaram"));
console.log('');

// 模拟算法演示
console.log('🎮 模拟算法演示:');
const sudokuBoard = [
    ["5","3",".",".","7",".",".",".","."],
    ["6",".",".","1","9","5",".",".","."],
    [".","9","8",".",".",".",".","6","."],
    ["8",".",".",".","6",".",".",".","3"],
    ["4",".",".","8",".","3",".",".","1"],
    ["7",".",".",".","2",".",".",".","6"],
    [".","6",".",".",".",".","2","8","."],
    [".",".",".","4","1","9",".",".","5"],
    [".",".",".",".","8",".",".","7","9"]
];
console.log('数独验证:', isValidSudoku(sudokuBoard));

const matrix = [[1,2,3],[4,5,6],[7,8,9]];
console.log('螺旋矩阵:', spiralOrder(matrix));

console.log('数组中心下标 [1,7,3,6,5,6]:', pivotIndex([1, 7, 3, 6, 5, 6]));

console.log('\n✅ 所有算法演示完成！');