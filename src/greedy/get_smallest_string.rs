use crate::greedy::Solution;

/*

给你一个仅由数字组成的字符串 s，在最多交换一次 相邻 且具有相同 奇偶性 的数字后，返回可以得到的字典序最小的字符串。
如果两个数字都是奇数或都是偶数，则它们具有相同的奇偶性。例如，5 和 9、2 和 4 奇偶性相同，而 6 和 9 奇偶性不同。

https://leetcode.cn/problems/lexicographically-smallest-string-after-a-swap/description/?envType=daily-question&envId=2024-11-02

只能交换一次。要使字典序最小，需要满足下面两个要求：

交换的两个数字，左边必须大于右边，否则交换不会使字典序变小。
交换的位置越靠左越好。比如示例 1 的 45320，交换 5 和 3 得到 43520，而交换更靠右的 2 和 0 得到 45302，这比 43520 更大。
*/
impl Solution {
    pub fn get_smallest_string(mut s: String) -> String {
        unsafe {
            let t = s.as_bytes_mut();
            for i in 1..t.len() {
                let x = t[i - 1];
                let y = t[i];
                if x > y && x % 2 == y % 2 {
                    t[i - 1] = y;
                    t[i] = x;
                    break;
                }
            }
            s
        }
    }
}
