use crate::hash::Solution;
use std::collections::HashSet;
/*
202. 快乐数
编写一个算法来判断一个数 n 是不是快乐数。
「快乐数」 定义为：
对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
如果这个过程 结果为 1，那么这个数就是快乐数。
如果 n 是 快乐数 就返回 true ；不是，则返回 false 。
https://leetcode.cn/problems/happy-number/description/
*/
impl Solution {
    pub fn is_happy(mut n: i32) -> bool {
        fn get_next(mut n: i32) -> i32 {
            let mut total_sum = 0;
            while n > 0 {
                let d = n % 10;
                n = n / 10;
                total_sum += d * d;
            }
            total_sum
        }
        let mut seen = HashSet::new();
        while n != 1 && !seen.contains(&n) {
            seen.insert(n);
            n = get_next(n);
        }
        n == 1
    }
}
