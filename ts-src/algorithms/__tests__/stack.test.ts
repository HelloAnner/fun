/**
 * 栈算法测试
 * Stack algorithms tests
 */

import { evalRPN, MinStack, isValid } from '../stack';

describe('Stack Algorithms', () => {
    describe('evalRPN', () => {
        test('should evaluate reverse polish notation correctly', () => {
            expect(evalRPN(["2", "1", "+", "3", "*"])).toBe(9);
            expect(evalRPN(["4", "13", "5", "/", "+"])).toBe(6);
            expect(evalRPN(["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"])).toBe(22);
        });
    });

    describe('MinStack', () => {
        test('should work as a min stack', () => {
            const minStack = new MinStack();
            minStack.push(-2);
            minStack.push(0);
            minStack.push(-3);
            expect(minStack.getMin()).toBe(-3);
            minStack.pop();
            expect(minStack.top()).toBe(0);
            expect(minStack.getMin()).toBe(-2);
        });
    });

    describe('isValid', () => {
        test('should validate parentheses correctly', () => {
            expect(isValid("()")).toBe(true);
            expect(isValid("()[]{}")).toBe(true);
            expect(isValid("(]")).toBe(false);
            expect(isValid("([)]")).toBe(false);
            expect(isValid("{[]}")).toBe(true);
        });
    });
});