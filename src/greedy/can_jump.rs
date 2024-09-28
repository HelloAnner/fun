use super::Solution;

/**

Given an integer list where each number represents the number of hops you can make, determine whether you can reach to the last index starting at index 0.
For example, [2, 0, 1, 0] returns True while [1, 1, 0, 1] returns False.


给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。

https://leetcode.cn/problems/jump-game/


 */

impl Solution {
    #[allow(dead_code)]
    pub fn can_jump(nums: Vec<i32>) -> bool {
        let n = nums.len();
        let mut right_most = 0;
        for (i, &jump) in nums.iter().enumerate() {
            if i <= right_most {
                right_most = right_most.max(i + jump as usize);
                if right_most >= n - 1 {
                    return true;
                }
            }
        }
        false
    }
}

#[test]
fn test_can_jump() {
    assert_eq!(true, Solution::can_jump(vec![2, 3, 1, 14]));
    assert_eq!(false, Solution::can_jump(vec![3, 2, 1, 0, 4]))
}
