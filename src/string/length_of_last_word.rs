use super::Solution;

/**
给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。

单词 是指仅由字母组成、不包含任何空格字符的最大
子字符串
。



示例 1：

输入：s = "Hello World"
输出：5
解释：最后一个单词是“World”，长度为 5。

https://leetcode.cn/problems/length-of-last-word/description/


 */

impl Solution {
    #[allow(dead_code)]
    pub fn length_of_last_word(s: String) -> i32 {
        let t = s.as_bytes();
        let mut i = s.len() - 1;
        while t[i] == b' ' {
            i -= 1;
        }
        let mut j = i as i32 - 1;
        while j >= 0 && t[j as usize] != b' ' {
            j -= 1;
        }
        i as i32 - j
    }
}
