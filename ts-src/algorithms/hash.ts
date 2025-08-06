/**
 * 哈希表相关算法实现
 * Hash table related algorithms
 */

/**
 * 2766. 重新放置石块
 * Relocate Marbles
 */
export function relocateMarbles(nums: number[], moveFrom: number[], moveTo: number[]): number[] {
    const stoneSet = new Set(nums);
    
    for (let i = 0; i < moveFrom.length; i++) {
        const from = moveFrom[i];
        const to = moveTo[i];
        stoneSet.delete(from);
        stoneSet.add(to);
    }
    
    return Array.from(stoneSet).sort((a, b) => a - b);
}

/**
 * 1. 两数之和
 * Two Sum
 */
export function twoSum(nums: number[], target: number): number[] {
    const map = new Map<number, number>();
    
    for (let i = 0; i < nums.length; i++) {
        const complement = target - nums[i];
        if (map.has(complement)) {
            return [map.get(complement)!, i];
        }
        map.set(nums[i], i);
    }
    
    return [];
}

/**
 * 202. 快乐数
 * Happy Number
 */
export function isHappy(n: number): boolean {
    const seen = new Set<number>();
    
    while (n !== 1 && !seen.has(n)) {
        seen.add(n);
        n = getNext(n);
    }
    
    return n === 1;
}

function getNext(n: number): number {
    let totalSum = 0;
    while (n > 0) {
        const digit = n % 10;
        n = Math.floor(n / 10);
        totalSum += digit * digit;
    }
    return totalSum;
}

/**
 * 997. 找到小镇的法官
 * Find the Town Judge
 */
export function findJudge(n: number, trust: number[][]): number {
    if (trust.length < n - 1) {
        return -1;
    }
    
    const inDegree = new Array(n + 1).fill(0);
    const outDegree = new Array(n + 1).fill(0);
    
    for (const [a, b] of trust) {
        outDegree[a]++;
        inDegree[b]++;
    }
    
    for (let i = 1; i <= n; i++) {
        if (inDegree[i] === n - 1 && outDegree[i] === 0) {
            return i;
        }
    }
    
    return -1;
}

/**
 * 1436. 旅行终点站
 * Destination City
 */
export function destCity(paths: string[][]): string {
    const outgoing = new Set<string>();
    
    for (const [from, to] of paths) {
        outgoing.add(from);
    }
    
    for (const [from, to] of paths) {
        if (!outgoing.has(to)) {
            return to;
        }
    }
    
    return "";
}