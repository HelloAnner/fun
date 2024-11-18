use crate::two_pointer::Solution;

/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k
同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。
https://leetcode.cn/problems/3sum/description/
*/
impl Solution {
    pub fn three_sum(mut nums: Vec<i32>) -> Vec<Vec<i32>> {
        nums.sort_unstable();
        let n = nums.len();
        let mut ans = vec![];
        for i in 0..n - 2 {
            let x = nums[i];
            if i > 0 && x == nums[i - 1] {
                continue;
            }
            if x + nums[i + 1] + nums[i + 2] > 0 {
                break;
            }
            if x + nums[n - 2] + nums[n - 1] < 0 {
                continue;
            }
            let mut j = i + 1;
            let mut k = n - 1;
            while j < k {
                let s = x + nums[j] + nums[k];
                if s > 0 {
                    k -= 1;
                } else if s < 0 {
                    j += 1;
                } else {
                    ans.push(vec![x, nums[j], nums[k]]);
                    j += 1;
                    while j < k && nums[j] == nums[j - 1] {
                        j += 1;
                    }
                    k -= 1;
                    while k > j && nums[k] == nums[k + 1] {
                        k -= 1;
                    }
                }
            }
        }
        ans
    }
}

#[test]
fn test_three_sum() {
    assert_eq!(
        vec![vec![-1, -1, 2], vec![-1, 0, 1]],
        Solution::three_sum(vec![-1, 0, 1, 2, -1, -4])
    )
}
