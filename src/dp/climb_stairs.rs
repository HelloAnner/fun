// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
// https://leetcode.cn/problems/climbing-stairs/description/?envType=study-plan-v2&envId=top-interview-150
#[allow(dead_code)]
pub fn climb_stairs(n: i32) -> i32 {
    let mut f0 = 1;
    let mut f1 = 1;
    for i in 2..=n {
        let new_f = f1 + f0;
        f0 = f1;
        f1 = new_f;
    }
    f1
}
