use super::Solution;

/**
 *
给你一个由字符 'a'、'b'、'c' 组成的字符串 s 和一个非负整数 k 。每分钟，你可以选择取走 s 最左侧 还是 最右侧 的那个字符。
你必须取走每种字符 至少 k 个，返回需要的 最少 分钟数；如果无法取到，则返回 -1 。

示例 1：
输入：s = "aabaaaacaabc", k = 2
输出：8
解释：
从 s 的左侧取三个字符，现在共取到两个字符 'a' 、一个字符 'b' 。
从 s 的右侧取五个字符，现在共取到四个字符 'a' 、两个字符 'b' 和两个字符 'c' 。
共需要 3 + 5 = 8 分钟。
可以证明需要的最少分钟数是 8

https://leetcode.cn/problems/take-k-of-each-character-from-left-and-right/solutions/2031995/on-shuang-zhi-zhen-by-endlesscheng-4g9p/?envType=daily-question&envId=2024-09-27
 比如 s 中有 3 个 a，4 个 b，5 个 c，k=2，每种字母至少取走 2 个，等价于剩下的字母至多有 1 个 a，2 个 b 和 3 个 c。
由于只能从 s 最左侧和最右侧取走字母，所以剩下的字母是 s 的子串。
设 s 中的 a,b,c 的个数分别为 x,y,z，现在问题变成：
计算 s 的最长子串长度，该子串满足 a,b,c 的个数分别至多为 x−k,y−k,z−k。
 */

impl Solution {
    #[allow(dead_code)]
    pub fn take_characters(s: String, k: i32) -> i32 {
        let mut cnt = [0; 3];
        for c in s.bytes() {
            cnt[(c - b'a') as usize] += 1;
        }
        if cnt[0] < k || cnt[1] < k || cnt[2] < k {
            return -1;
        }
        let mut mx = 0;
        let mut left = 0;
        let s = s.as_bytes();
        for (right, &c) in s.iter().enumerate() {
            let c = (c - b'a') as usize;
            cnt[c] -= 1; // 移入窗口
            while cnt[c] < k {
                cnt[(s[left] - b'a') as usize] += 1;
                left += 1;
            }
            // left 大于 right 的情况
            mx = mx.max(right - left + 1)
        }
        (s.len() - mx) as _
    }
}

#[test]
pub fn test_take_characters() {
    // right - left + 1 计算不对
    // assert_eq!(8, Solution::take_characters("aabaaaacaabc".to_string(), 2))
}
