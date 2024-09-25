use std::collections::HashSet;

use super::Solution;

/**
 * 给你一个字符串数组 ideas 表示在公司命名过程中使用的名字列表。公司命名流程如下：

从 ideas 中选择 2 个 不同 名字，称为 ideaA 和 ideaB 。
交换 ideaA 和 ideaB 的首字母。
如果得到的两个新名字 都 不在 ideas 中，那么 ideaA ideaB（串联 ideaA 和 ideaB ，中间用一个空格分隔）是一个有效的公司名字。
否则，不是一个有效的名字。
返回 不同 且有效的公司名字的数目。



示例 1：

输入：ideas = ["coffee","donuts","time","toffee"]
输出：6
解释：下面列出一些有效的选择方案：
- ("coffee", "donuts")：对应的公司名字是 "doffee conuts" 。
- ("donuts", "coffee")：对应的公司名字是 "conuts doffee" 。
- ("donuts", "time")：对应的公司名字是 "tonuts dime" 。
- ("donuts", "toffee")：对应的公司名字是 "tonuts doffee" 。
- ("time", "donuts")：对应的公司名字是 "dime tonuts" 。
- ("toffee", "donuts")：对应的公司名字是 "doffee tonuts" 。
因此，总共有 6 个不同的公司名字。

https://leetcode.cn/problems/naming-a-company/description/?envType=daily-question&envId=2024-09-25

按照首字母，把 ideas 分成（至多）26 组字符串。

例如 ideas=[aa,ab,ac,bc,bd,be] 分成如下两组（只记录去掉首字母后的字符串）：

A 组（集合）：{a,b,c}。
B 组（集合）：{c,d,e}。
分组后：

从 A 中选一个不等于 c 的字符串，这有 2 种选法。
从 B 中选一个不等于 c 的字符串，这有 2 种选法。
考虑两个字符串先后顺序（谁在左谁在右），有 2 种方法。
根据乘法原理，有 2×2×2=8 对符合要求的字符串。

由于无法选交集中的字符串，一般地，从 A 和 B 中可以选出

2⋅(∣A∣−∣A∩B∣)⋅(∣B∣−∣A∩B∣)
对符合要求的字符串。其中 ∣S∣ 表示集合 S 的大小。
 */

#[allow(dead_code)]
impl Solution {
    pub fn distinct_names(ideas: Vec<String>) -> i64 {
        let mut groups = vec![HashSet::new(); 26];
        for s in ideas {
            groups[(s.as_bytes()[0] - b'a') as usize].insert(s[1..].to_string());
        }

        let mut ans = 0i64;
        for a in 1..26 {
            for b in 0..a {
                let m = groups[a].iter().filter(|&s| groups[b].contains(s)).count();
                ans += (groups[a].len() - m) as i64 * (groups[b].len() - m) as i64;
            }
        }
        ans * 2
    }
}

#[test]
fn test_distinct_names() {
    let ideas = vec![
        "aaa".to_string(),
        "baa".to_string(),
        "caa".to_string(),
        "bbb".to_string(),
        "cbb".to_string(),
        "dbb".to_string(),
    ];
    assert_eq!(Solution::distinct_names(ideas), 2);
}
