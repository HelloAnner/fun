use super::Solution;
/**
给你 n 个项目，编号从 0 到 n - 1 。同时给你一个整数数组 milestones ，其中每个 milestones[i] 表示第 i 个项目中的阶段任务数量。

你可以按下面两个规则参与项目中的工作：

1. 每周，你将会完成 某一个 项目中的 恰好一个 阶段任务。你每周都 必须 工作。
2. 在 连续的 两周中，你 不能 参与并完成同一个项目中的两个阶段任务。
一旦所有项目中的全部阶段任务都完成，或者仅剩余一个阶段任务都会导致你违反上面的规则，那么你将 停止工作 。注意，由于这些条件的限制，你可能无法完成所有阶段任务。

返回在不违反上面规则的情况下你 最多 能工作多少周。

https://leetcode.cn/problems/maximum-number-of-weeks-for-which-you-can-work/description/


考虑耗时最长的工作。假设我们需要 longest 周完成该工作，其余工作共计需要 rest 周完成。那么可以完成所有工作的充要条件是：
longest <=  rest + 1
如果 longest <= rest + 1 ， 那么可以完成的所有的工作
反之 ， 2 * rest + 1
 */

impl Solution {
    #[allow(dead_code)]
    pub fn number_of_weeks(milestones: Vec<i32>) -> i64 {
        let longest = *milestones.iter().max().unwrap() as i64;
        let rest = milestones.iter().map(|&n| n as i64).sum::<i64>() - longest;
        if longest > rest + 1 {
            rest * 2 + 1
        } else {
            longest + rest
        }
    }
}

#[test]
fn test_number_of_weeks() {
    assert_eq!(6, Solution::number_of_weeks(vec![1, 2, 3]))
}
