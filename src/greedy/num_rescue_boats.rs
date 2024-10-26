use crate::greedy::Solution;
/**
给定数组 people 。people[i]表示第 i 个人的体重 ，船的数量不限，每艘船可以承载的最大重量为 limit。
每艘船最多可同时载两人，但条件是这些人的重量之和最多为 limit。
返回 承载所有人所需的最小船数 。

https://leetcode.cn/problems/boats-to-save-people/description/


*/
impl Solution {
    pub fn num_rescue_boats(mut people: Vec<i32>, limit: i32) -> i32 {
        let mut ans = 0;
        people.sort();
        let mut light = 0;
        let mut heavy = people.len() - 1;
        while light <= heavy {
            if people[light] + people[heavy] <= limit {
                light += 1;
            }
            heavy -= 1;
            ans += 1;
        }
        ans
    }
}
