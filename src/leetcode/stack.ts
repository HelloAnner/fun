/**
 * 栈相关算法实现
 * Stack-related algorithms
 */

/**
 * 150. 逆波兰表达式求值
 * Evaluate Reverse Polish Notation
 */
export function evalRPN(tokens: string[]): number {
    const opToBinaryFn: { [key: string]: (x: number, y: number) => number } = {
        '+': (x, y) => x + y,
        '-': (x, y) => x - y,
        '*': (x, y) => x * y,
        '/': (x, y) => Math.trunc(x / y) // 向零截断
    };

    const stack: number[] = [];
    
    for (const token of tokens) {
        if (token in opToBinaryFn) {
            const num2 = stack.pop()!;
            const num1 = stack.pop()!;
            const result = opToBinaryFn[token](num1, num2);
            stack.push(result);
        } else {
            stack.push(parseInt(token));
        }
    }
    
    return stack[0];
}

/**
 * 最小栈实现
 * Min Stack implementation
 */
export class MinStack {
    private stack: number[];
    private minStack: number[];

    constructor() {
        this.stack = [];
        this.minStack = [];
    }

    push(val: number): void {
        this.stack.push(val);
        if (this.minStack.length === 0 || val <= this.getMin()) {
            this.minStack.push(val);
        }
    }

    pop(): void {
        const val = this.stack.pop();
        if (val === this.getMin()) {
            this.minStack.pop();
        }
    }

    top(): number {
        return this.stack[this.stack.length - 1];
    }

    getMin(): number {
        return this.minStack[this.minStack.length - 1];
    }
}

/**
 * 有效的括号
 * Valid Parentheses
 */
export function isValid(s: string): boolean {
    const stack: string[] = [];
    const mapping: { [key: string]: string } = {
        ')': '(',
        '}': '{',
        ']': '['
    };

    for (const char of s) {
        if (char in mapping) {
            const topElement = stack.length === 0 ? '#' : stack.pop()!;
            if (mapping[char] !== topElement) {
                return false;
            }
        } else {
            stack.push(char);
        }
    }

    return stack.length === 0;
}