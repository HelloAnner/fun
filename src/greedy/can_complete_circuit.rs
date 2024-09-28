use std::i32;

use super::Solution;

/**

在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。

你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。

给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。

https://leetcode.cn/problems/gas-station/description/

直接从全局进行贪心选择，情况如下：

情况一：如果gas的总和小于cost总和，那么无论从哪里出发，一定是跑不了一圈的

情况二：rest[i] = gas[i]-cost[i]为一天剩下的油，i从0开始计算累加到最后一站，如果累加没有出现负数，说明从0出发，油就没有断过，那么0就是起点。

情况三：如果累加的最小值是负数，汽车就要从非0节点出发，从后向前，看哪个节点能把这个负数填平，能把这个负数填平的节点就是出发节点。

 */

impl Solution {
    #[allow(dead_code)]
    pub fn can_complete_circuit(gas: Vec<i32>, cost: Vec<i32>) -> i32 {
        let mut cur_sum = 0;
        let mut min_val = i32::MAX;
        for (i, &v) in gas.iter().enumerate() {
            let rest = v - cost[i];
            cur_sum += rest;
            if cur_sum < min_val {
                min_val = cur_sum;
            }
        }
        // 情况一
        if cur_sum < 0 {
            return -1;
        }
        // 情况二
        if min_val >= 0 {
            return 0;
        }
        for i in (0..gas.len()).rev() {
            let rest = gas[i] - cost[i];
            min_val += rest;
            if min_val >= 0 {
                return i as _;
            }
        }
        -1
    }
}

#[test]
fn test_can_complete_circuit() {
    assert_eq!(
        3,
        Solution::can_complete_circuit(vec![1, 2, 3, 4, 5], vec![3, 4, 5, 1, 2])
    )
}
