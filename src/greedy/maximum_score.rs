/*
要求选手从 N 张卡牌中选出 cnt 张卡牌，若这 cnt 张卡牌数字总和为偶数，则选手成绩「有效」且得分为 cnt 张卡牌数字总和。
给定数组 cards 和 cnt，其中 cards[i] 表示第 i 张卡牌上的数字。 请帮参赛选手计算最大的有效得分。若不存在获取有效得分的卡牌方案，则返回 0。
https://leetcode.cn/problems/uOAnQW/description/

将 cards 从大到小排序后，先贪心的将后 cnt 个数字加起来，若此时 sum 为偶数，直接返回即可。

若此时答案为奇数，有两种方案:
在数组前面找到一个最大的奇数与后 cnt 个数中最小的偶数进行替换；
在数组前面找到一个最大的偶数与后 cnt 个数中最小的奇数进行替换。
两种方案选最大值即可。
*/
use crate::greedy::Solution;

impl Solution {
    pub fn maxmium_score(cards: Vec<i32>, cnt: i32) -> i32 {
        let mut cards = cards;
        cards.sort_by(|a, b| b.cmp(a));

        let mut ans = 0;
        let mut tmp = 0;
        let mut odd = -1;
        let mut even = -1;

        for i in 0..cnt as usize {
            tmp += cards[i];
            if cards[i] % 2 == 1 {
                odd = cards[i];
            } else {
                even = cards[i];
            }
        }

        if tmp % 2 == 0 {
            return tmp;
        }

        for i in cnt as usize..cards.len() {
            if cards[i] % 2 == 1 {
                if even != -1 {
                    ans = ans.max(tmp - even + cards[i]);
                }
            } else {
                if odd != -1 {
                    ans = ans.max(tmp - odd + cards[i]);
                }
            }
        }
        ans
    }
}
