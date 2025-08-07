/**
 * 滑动窗口相关算法实现
 * Sliding Window related algorithms
 */

/**
 * 3. 无重复字符的最长子串
 * Longest Substring Without Repeating Characters
 */
export function lengthOfLongestSubstring(s: string): number {
    const charSet = new Set<string>();
    let left = 0;
    let maxLength = 0;
    
    for (let right = 0; right < s.length; right++) {
        while (charSet.has(s[right])) {
            charSet.delete(s[left]);
            left++;
        }
        charSet.add(s[right]);
        maxLength = Math.max(maxLength, right - left + 1);
    }
    
    return maxLength;
}

/**
 * 76. 最小覆盖子串
 * Minimum Window Substring
 */
export function minWindow(s: string, t: string): string {
    if (s.length < t.length) return "";
    
    const need = new Map<string, number>();
    const window = new Map<string, number>();
    
    // 统计 t 中字符频次
    for (const char of t) {
        need.set(char, (need.get(char) || 0) + 1);
    }
    
    let left = 0;
    let right = 0;
    let valid = 0;
    let start = 0;
    let len = Infinity;
    
    while (right < s.length) {
        const c = s[right];
        right++;
        
        if (need.has(c)) {
            window.set(c, (window.get(c) || 0) + 1);
            if (window.get(c) === need.get(c)) {
                valid++;
            }
        }
        
        while (valid === need.size) {
            if (right - left < len) {
                start = left;
                len = right - left;
            }
            
            const d = s[left];
            left++;
            
            if (need.has(d)) {
                if (window.get(d) === need.get(d)) {
                    valid--;
                }
                window.set(d, window.get(d)! - 1);
            }
        }
    }
    
    return len === Infinity ? "" : s.substring(start, start + len);
}

/**
 * 209. 长度最小的子数组
 * Minimum Size Subarray Sum
 */
export function minSubArrayLen(target: number, nums: number[]): number {
    let left = 0;
    let sum = 0;
    let minLen = Infinity;
    
    for (let right = 0; right < nums.length; right++) {
        sum += nums[right];
        
        while (sum >= target) {
            minLen = Math.min(minLen, right - left + 1);
            sum -= nums[left];
            left++;
        }
    }
    
    return minLen === Infinity ? 0 : minLen;
}

/**
 * 219. 存在重复元素 II
 * Contains Duplicate II
 */
export function containsNearbyDuplicate(nums: number[], k: number): boolean {
    const map = new Map<number, number>();
    
    for (let i = 0; i < nums.length; i++) {
        if (map.has(nums[i]) && i - map.get(nums[i])! <= k) {
            return true;
        }
        map.set(nums[i], i);
    }
    
    return false;
}

/**
 * 438. 找到字符串中所有字母异位词
 * Find All Anagrams in a String
 */
export function findAnagrams(s: string, p: string): number[] {
    const result: number[] = [];
    if (s.length < p.length) return result;
    
    const need = new Map<string, number>();
    const window = new Map<string, number>();
    
    // 统计 p 中字符频次
    for (const char of p) {
        need.set(char, (need.get(char) || 0) + 1);
    }
    
    let left = 0;
    let right = 0;
    let valid = 0;
    
    while (right < s.length) {
        const c = s[right];
        right++;
        
        if (need.has(c)) {
            window.set(c, (window.get(c) || 0) + 1);
            if (window.get(c) === need.get(c)) {
                valid++;
            }
        }
        
        while (right - left >= p.length) {
            if (valid === need.size) {
                result.push(left);
            }
            
            const d = s[left];
            left++;
            
            if (need.has(d)) {
                if (window.get(d) === need.get(d)) {
                    valid--;
                }
                window.set(d, window.get(d)! - 1);
            }
        }
    }
    
    return result;
}

/**
 * 1052. 爱生气的书店老板
 * Grumpy Bookstore Owner
 */
export function maxSatisfied(customers: number[], grumpy: number[], minutes: number): number {
    let satisfied = 0;
    
    // 计算不使用技巧时的满意顾客数
    for (let i = 0; i < customers.length; i++) {
        if (grumpy[i] === 0) {
            satisfied += customers[i];
        }
    }
    
    // 滑动窗口找到最大增益
    let maxGain = 0;
    let currentGain = 0;
    
    // 初始窗口
    for (let i = 0; i < minutes; i++) {
        if (grumpy[i] === 1) {
            currentGain += customers[i];
        }
    }
    maxGain = currentGain;
    
    // 滑动窗口
    for (let i = minutes; i < customers.length; i++) {
        if (grumpy[i] === 1) {
            currentGain += customers[i];
        }
        if (grumpy[i - minutes] === 1) {
            currentGain -= customers[i - minutes];
        }
        maxGain = Math.max(maxGain, currentGain);
    }
    
    return satisfied + maxGain;
}

/**
 * 2024. 考试的最大困扰度
 * Maximize the Confusion of an Exam
 */
export function maxConsecutiveAnswers(answerKey: string, k: number): number {
    const maxConsecutive = (char: string): number => {
        let left = 0;
        let count = 0;
        let maxLen = 0;
        
        for (let right = 0; right < answerKey.length; right++) {
            if (answerKey[right] !== char) {
                count++;
            }
            
            while (count > k) {
                if (answerKey[left] !== char) {
                    count--;
                }
                left++;
            }
            
            maxLen = Math.max(maxLen, right - left + 1);
        }
        
        return maxLen;
    };
    
    return Math.max(maxConsecutive('T'), maxConsecutive('F'));
}

/**
 * 2009. 使数组连续的最少操作数
 * Minimum Number of Operations to Make Array Continuous
 */
export function minOperations(nums: number[]): number {
    const n = nums.length;
    const uniqueNums = [...new Set(nums)].sort((a, b) => a - b);
    let maxCount = 0;
    let left = 0;
    
    for (let right = 0; right < uniqueNums.length; right++) {
        while (uniqueNums[right] - uniqueNums[left] >= n) {
            left++;
        }
        maxCount = Math.max(maxCount, right - left + 1);
    }
    
    return n - maxCount;
}

/**
 * 2831. 找出最长等值子数组
 * Find the Longest Equal Subarray
 */
export function longestEqualSubarray(nums: number[], k: number): number {
    const positions = new Map<number, number[]>();
    
    // 记录每个值的所有位置
    for (let i = 0; i < nums.length; i++) {
        if (!positions.has(nums[i])) {
            positions.set(nums[i], []);
        }
        positions.get(nums[i])!.push(i);
    }
    
    let maxLength = 0;
    
    // 对每个值使用滑动窗口
    for (const [value, pos] of positions) {
        let left = 0;
        
        for (let right = 0; right < pos.length; right++) {
            // 计算需要删除的元素数量
            const windowSize = pos[right] - pos[left] + 1;
            const equalElements = right - left + 1;
            const toDelete = windowSize - equalElements;
            
            while (toDelete > k) {
                left++;
                const newWindowSize = pos[right] - pos[left] + 1;
                const newEqualElements = right - left + 1;
                const newToDelete = newWindowSize - newEqualElements;
                if (newToDelete <= k) break;
            }
            
            maxLength = Math.max(maxLength, right - left + 1);
        }
    }
    
    return maxLength;
}