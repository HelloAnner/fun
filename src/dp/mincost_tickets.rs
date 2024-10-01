use super::Solution;

/**
在一个火车旅行很受欢迎的国度，你提前一年计划了一些火车旅行。在接下来的一年里，你要旅行的日子将以一个名为 days 的数组给出。每一项是一个从 1 到 365 的整数。

火车票有 三种不同的销售方式 ：

一张 为期一天 的通行证售价为 costs[0] 美元；
一张 为期七天 的通行证售价为 costs[1] 美元；
一张 为期三十天 的通行证售价为 costs[2] 美元。
通行证允许数天无限制的旅行。 例如，如果我们在第 2 天获得一张 为期 7 天 的通行证，那么我们可以连着旅行 7 天：第 2 天、第 3 天、第 4 天、第 5 天、第 6 天、第 7 天和第 8 天。

返回 你想要完成在给定的列表 days 中列出的每一天的旅行所需要的最低消费 。


https://leetcode.cn/problems/minimum-cost-for-tickets/description/?envType=daily-question&envId=2024-10-01

根据上面的讨论，定义 dfs(i) 表示 1 到 i 天的最小花费。

如果第 i 天不在 days 中，那么问题变成 1 到 i−1 天的最小花费，即

dfs(i)=dfs(i−1)
如果第 i 天在 days 中，分类讨论：

在第 i 天购买为期 1 天的通行证，接下来需要解决的问题为：1 到 i−1 天的最小花费，即 dfs(i)=dfs(i−1)+costs[0]。
在第 i−6 天购买为期 7 天的通行证，接下来需要解决的问题为：1 到 i−7 天的最小花费，即 dfs(i)=dfs(i−7)+costs[1]。
在第 i−29 天购买为期 30 天的通行证，接下来需要解决的问题为：1 到 i−30 天的最小花费，即 dfs(i)=dfs(i−30)+costs[2]。
这三种情况取最小值，就得到了 dfs(i)，即

dfs(i)=min(dfs(i−1)+costs[0],dfs(i−7)+costs[1],dfs(i−30)+costs[2])

 */

impl Solution {
    #[allow(dead_code)]
    pub fn mincost_tickets(days: Vec<i32>, costs: Vec<i32>) -> i32 {
        let last_day = *days.last().unwrap();
        let mut is_travel = vec![false; last_day as usize + 1];
        for &day in &days {
            is_travel[day as usize] = true;
        }
        let mut memo = vec![0; last_day as usize + 1];
        Self::dfs_mincost_tickets(last_day, &is_travel, &costs, &mut memo)
    }

    fn dfs_mincost_tickets(
        i: i32,
        is_travel: &Vec<bool>,
        costs: &Vec<i32>,
        memo: &mut Vec<i32>,
    ) -> i32 {
        if i <= 0 {
            return 0;
        }
        let index = i as usize;
        if memo[index] > 0 {
            return memo[index];
        }
        if !is_travel[index] {
            memo[index] = Self::dfs_mincost_tickets(i - 1, is_travel, costs, memo);
            return memo[index];
        }
        memo[index] = i32::min(
            Self::dfs_mincost_tickets(i - 1, is_travel, costs, memo) + costs[0],
            i32::min(
                Self::dfs_mincost_tickets(i - 7, is_travel, costs, memo) + costs[1],
                Self::dfs_mincost_tickets(i - 30, is_travel, costs, memo) + costs[2],
            ),
        );
        memo[index]
    }
}

#[test]
fn test_mincost_tickets() {
    assert_eq!(
        11,
        Solution::mincost_tickets(vec![1, 4, 6, 7, 8, 20], vec![2, 7, 15])
    )
}
