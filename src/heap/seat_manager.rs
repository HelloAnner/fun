use std::collections::BinaryHeap;

/**

请你设计一个管理 n 个座位预约的系统，座位编号从 1 到 n 。

请你实现 SeatManager 类：

SeatManager(int n) 初始化一个 SeatManager 对象，它管理从 1 到 n 编号的 n 个座位。所有座位初始都是可预约的。
int reserve() 返回可以预约座位的 最小编号 ，此座位变为不可预约。
void unreserve(int seatNumber) 将给定编号 seatNumber 对应的座位变成可以预约。

https://leetcode.cn/problems/seat-reservation-manager/description/?envType=daily-question&envId=2024-09-30

 */
struct SeatManager {
    available: BinaryHeap<i32>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
#[allow(dead_code)]
impl SeatManager {
    fn new(n: i32) -> Self {
        let mut available = BinaryHeap::new();
        for i in 1..=n {
            // 取相反数，变为最小堆
            available.push(-i);
        }
        Self { available }
    }

    fn reserve(&mut self) -> i32 {
        -self.available.pop().unwrap()
    }

    fn unreserve(&mut self, seat_number: i32) {
        self.available.push(-seat_number);
    }
}
