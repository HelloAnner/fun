use super::Solution;

/**

给你一个二进制数组 nums 。

你可以对数组执行以下操作 任意 次（也可以 0 次）：

选择数组中 任意连续 3 个元素，并将它们 全部反转 。
反转 一个元素指的是将它的值从 0 变 1 ，或者从 1 变 0 。

请你返回将 nums 中所有元素变为 1 的 最少 操作次数。如果无法全部变成 1 ，返回 -1 。

示例 1：

输入：nums = [0,1,1,1,0,0]

输出：3

解释：
我们可以执行以下操作：

选择下标为 0 ，1 和 2 的元素并反转，得到 nums = [1,0,0,1,0,0] 。
选择下标为 1 ，2 和 3 的元素并反转，得到 nums = [1,1,1,0,0,0] 。
选择下标为 3 ，4 和 5 的元素并反转，得到 nums = [1,1,1,1,1,1] 。



https://leetcode.cn/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-i/description/?envType=daily-question&envId=2024-10-18

讨论是否需要对 i=0 执行操作：

如果 nums[0]=1，不需要操作，问题变成剩下 n−1 个数的子问题。
如果 nums[0]=0，一定要操作，问题变成剩下 n−1 个数的子问题。
接下来，讨论是否需要对 i=1 执行操作，处理方式同上。

依此类推，一直到 i=n−3 处理完后，还剩下 nums[n−2] 和 nums[n−1]，这两个数必须都等于 1，否则无法达成题目要求。
 */

impl Solution {
    #[allow(dead_code)]
    pub fn min_operations(mut nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut ans = 0;
        for i in 0..n - 2 {
            if nums[i] == 0 {
                nums[i + 1] ^= 1;
                nums[i + 2] ^= 1;
                ans += 1;
            }
        }

        if nums[n - 2] != 0 && nums[n - 1] != 0 {
            ans
        } else {
            -1
        }
    }
}
