/**
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

实现 MinStack 类:

MinStack() 初始化堆栈对象。
void push(int val) 将元素val推入堆栈。
void pop() 删除堆栈顶部的元素。
int top() 获取堆栈顶部的元素。
int getMin() 获取堆栈中的最小元素。


https://leetcode.cn/problems/min-stack/description/
 */

struct MinStack {
    stk: Vec<i32>,
    min_stk: Vec<i32>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MinStack {
    #[allow(dead_code)]
    fn new() -> Self {
        Self {
            stk: vec![],
            min_stk: vec![],
        }
    }

    #[allow(dead_code)]
    fn push(&mut self, val: i32) {
        self.stk.push(val);
        if self.min_stk.is_empty() || val <= *self.min_stk.last().unwrap() {
            self.min_stk.push(val);
        }
    }

    #[allow(dead_code)]
    fn pop(&mut self) {
        if self.stk.pop().unwrap() == *self.min_stk.last().unwrap() {
            self.min_stk.pop();
        }
    }

    #[allow(dead_code)]
    fn top(&self) -> i32 {
        *self.stk.last().unwrap()
    }

    #[allow(dead_code)]
    fn get_min(&self) -> i32 {
        *self.min_stk.last().unwrap()
    }
}
