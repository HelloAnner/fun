use super::Solution;

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。


示例 1：
输入：s = "()"
输出：true
 */

impl Solution {
    #[allow(dead_code)]
    pub fn is_valid(s: String) -> bool {
        if s.len() % 2 != 0 {
            return false;
        }
        let mut st = vec![];
        for c in s.bytes() {
            match c {
                b'(' => st.push(b')'),
                b'[' => st.push(b']'),
                b'{' => st.push(b'}'),
                _ => {
                    if st.pop() != Some(c) {
                        return false;
                    }
                }
            }
        }
        st.is_empty()
    }
}
