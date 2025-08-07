/**
 * 位操作相关算法实现
 * Bit Manipulation related algorithms
 */

/**
 * 136. 只出现一次的数字
 * Single Number
 */
export function singleNumber(nums: number[]): number {
    let result = 0;
    for (const num of nums) {
        result ^= num;
    }
    return result;
}

/**
 * 137. 只出现一次的数字 II
 * Single Number II
 */
export function singleNumberII(nums: number[]): number {
    let ones = 0;
    let twos = 0;
    
    for (const num of nums) {
        ones = (ones ^ num) & ~twos;
        twos = (twos ^ num) & ~ones;
    }
    
    return ones;
}

/**
 * 260. 只出现一次的数字 III
 * Single Number III
 */
export function singleNumberIII(nums: number[]): number[] {
    let xor = 0;
    for (const num of nums) {
        xor ^= num;
    }
    
    // 找到最右边的1
    const rightmostBit = xor & (-xor);
    
    let num1 = 0;
    let num2 = 0;
    
    for (const num of nums) {
        if (num & rightmostBit) {
            num1 ^= num;
        } else {
            num2 ^= num;
        }
    }
    
    return [num1, num2];
}

/**
 * 191. 位1的个数
 * Number of 1 Bits
 */
export function hammingWeight(n: number): number {
    let count = 0;
    while (n !== 0) {
        count++;
        n &= n - 1; // 清除最右边的1
    }
    return count;
}

/**
 * 338. 比特位计数
 * Counting Bits
 */
export function countBits(n: number): number[] {
    const result = new Array(n + 1).fill(0);
    
    for (let i = 1; i <= n; i++) {
        result[i] = result[i >> 1] + (i & 1);
    }
    
    return result;
}

/**
 * 461. 汉明距离
 * Hamming Distance
 */
export function hammingDistance(x: number, y: number): number {
    let xor = x ^ y;
    let count = 0;
    
    while (xor !== 0) {
        count++;
        xor &= xor - 1;
    }
    
    return count;
}

/**
 * 477. 汉明距离总和
 * Total Hamming Distance
 */
export function totalHammingDistance(nums: number[]): number {
    let total = 0;
    const n = nums.length;
    
    for (let i = 0; i < 32; i++) {
        let ones = 0;
        for (const num of nums) {
            ones += (num >> i) & 1;
        }
        total += ones * (n - ones);
    }
    
    return total;
}

/**
 * 190. 颠倒二进制位
 * Reverse Bits
 */
export function reverseBits(n: number): number {
    let result = 0;
    for (let i = 0; i < 32; i++) {
        result = (result << 1) | (n & 1);
        n >>= 1;
    }
    return result >>> 0; // 无符号右移确保正数
}

/**
 * 231. 2的幂
 * Power of Two
 */
export function isPowerOfTwo(n: number): boolean {
    return n > 0 && (n & (n - 1)) === 0;
}

/**
 * 342. 4的幂
 * Power of Four
 */
export function isPowerOfFour(n: number): boolean {
    return n > 0 && (n & (n - 1)) === 0 && (n & 0x55555555) !== 0;
}

/**
 * 371. 两整数之和
 * Sum of Two Integers
 */
export function getSum(a: number, b: number): number {
    while (b !== 0) {
        const carry = (a & b) << 1;
        a = a ^ b;
        b = carry;
    }
    return a;
}

/**
 * 389. 找不同
 * Find the Difference
 */
export function findTheDifference(s: string, t: string): string {
    let xor = 0;
    
    for (const char of s) {
        xor ^= char.charCodeAt(0);
    }
    
    for (const char of t) {
        xor ^= char.charCodeAt(0);
    }
    
    return String.fromCharCode(xor);
}

/**
 * 405. 数字转换为十六进制数
 * Convert a Number to Hexadecimal
 */
export function toHex(num: number): string {
    if (num === 0) return "0";
    
    const hex = "0123456789abcdef";
    let result = "";
    
    // 处理负数
    if (num < 0) {
        num = 0x100000000 + num; // 转换为32位无符号数
    }
    
    while (num > 0) {
        result = hex[num & 15] + result;
        num >>>= 4;
    }
    
    return result;
}

/**
 * 421. 数组中两个数的最大异或值
 * Maximum XOR of Two Numbers in an Array
 */
export function findMaximumXOR(nums: number[]): number {
    let maxXor = 0;
    let mask = 0;
    
    for (let i = 30; i >= 0; i--) {
        mask |= (1 << i);
        const prefixes = new Set<number>();
        
        for (const num of nums) {
            prefixes.add(num & mask);
        }
        
        const candidate = maxXor | (1 << i);
        
        for (const prefix of prefixes) {
            if (prefixes.has(candidate ^ prefix)) {
                maxXor = candidate;
                break;
            }
        }
    }
    
    return maxXor;
}

/**
 * 476. 数字的补数
 * Number Complement
 */
export function findComplement(num: number): number {
    let mask = 1;
    let temp = num;
    
    while (temp > 0) {
        mask <<= 1;
        temp >>= 1;
    }
    
    return (mask - 1) ^ num;
}

/**
 * 693. 交替位二进制数
 * Binary Number with Alternating Bits
 */
export function hasAlternatingBits(n: number): boolean {
    const xor = n ^ (n >> 1);
    return (xor & (xor + 1)) === 0;
}

/**
 * 762. 二进制表示中质数个计算置位
 * Prime Number of Set Bits in Binary Representation
 */
export function countPrimeSetBits(left: number, right: number): number {
    const isPrime = (n: number): boolean => {
        if (n < 2) return false;
        if (n === 2) return true;
        if (n % 2 === 0) return false;
        
        for (let i = 3; i * i <= n; i += 2) {
            if (n % i === 0) return false;
        }
        return true;
    };
    
    let count = 0;
    
    for (let i = left; i <= right; i++) {
        const setBits = hammingWeight(i);
        if (isPrime(setBits)) {
            count++;
        }
    }
    
    return count;
}

/**
 * 868. 二进制间距
 * Binary Gap
 */
export function binaryGap(n: number): number {
    let maxGap = 0;
    let lastOne = -1;
    
    for (let i = 0; i < 32; i++) {
        if ((n >> i) & 1) {
            if (lastOne !== -1) {
                maxGap = Math.max(maxGap, i - lastOne);
            }
            lastOne = i;
        }
    }
    
    return maxGap;
}

/**
 * 1318. 或运算的最小翻转次数
 * Minimum Flips to Make a OR b Equal to c
 */
export function minFlips(a: number, b: number, c: number): number {
    let flips = 0;
    
    while (a > 0 || b > 0 || c > 0) {
        const bitA = a & 1;
        const bitB = b & 1;
        const bitC = c & 1;
        
        if (bitC === 0) {
            flips += bitA + bitB;
        } else {
            if (bitA === 0 && bitB === 0) {
                flips++;
            }
        }
        
        a >>= 1;
        b >>= 1;
        c >>= 1;
    }
    
    return flips;
}