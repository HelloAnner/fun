use crate::sliding_window::Solution;

/*
给你一个 二进制 字符串 s 和一个整数 k。
如果一个 二进制字符串 满足以下任一条件，则认为该字符串满足 k 约束：
字符串中 0 的数量最多为 k。
字符串中 1 的数量最多为 k。
返回一个整数，表示 s 的所有满足 k 约束 的子字符串的数量。
https://leetcode.cn/problems/count-substrings-that-satisfy-k-constraint-i/description/?envType=daily-question&envId=2024-11-12
*/
impl Solution {
    pub fn count_k_constraint_substrings(s: String, k: i32) -> i32 {
        let s: Vec<char> = s.chars().collect();
        let mut ans = 0;
        let mut left = 0;
        let mut cnt = [0; 2];
        for i in 0..s.len() {
            cnt[(s[i] as i32 & 1) as usize] += 1;
            while cnt[0] > k && cnt[1] > k {
                cnt[(s[left] as i32 & 1) as usize] -= 1;
                left += 1;
            }
            ans += (i - left + 1) as i32;
        }
        ans
    }
}

#[test]
fn test_count_k_constraint_substrings() {
    assert_eq!(
        12,
        Solution::count_k_constraint_substrings("10101".to_string(), 1)
    );
}
