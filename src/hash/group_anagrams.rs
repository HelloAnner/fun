use crate::hash::Solution;
use std::collections::HashMap;
/*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
https://leetcode.cn/problems/group-anagrams/description/
*/
impl Solution {
    pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
        let mut m = HashMap::new();
        for s in strs {
            let mut sorted_s = s.clone().into_bytes();
            sorted_s.sort_unstable();
            m.entry(sorted_s).or_insert(vec![]).push(s)
        }
        m.into_values().collect()
    }
}
