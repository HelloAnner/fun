/**
 * 贪心算法相关实现
 * Greedy algorithms
 */

/**
 * 55. 跳跃游戏
 * Jump Game
 */
export function canJump(nums: number[]): boolean {
    let maxReach = 0;
    
    for (let i = 0; i < nums.length; i++) {
        if (i > maxReach) {
            return false;
        }
        maxReach = Math.max(maxReach, i + nums[i]);
    }
    
    return true;
}

/**
 * 45. 跳跃游戏 II
 * Jump Game II
 */
export function jump(nums: number[]): number {
    let jumps = 0;
    let currentEnd = 0;
    let farthest = 0;
    
    for (let i = 0; i < nums.length - 1; i++) {
        farthest = Math.max(farthest, i + nums[i]);
        
        if (i === currentEnd) {
            jumps++;
            currentEnd = farthest;
        }
    }
    
    return jumps;
}

/**
 * 121. 买卖股票的最佳时机
 * Best Time to Buy and Sell Stock
 */
export function maxProfit(prices: number[]): number {
    let minPrice = prices[0];
    let maxProfit = 0;
    
    for (let i = 1; i < prices.length; i++) {
        if (prices[i] < minPrice) {
            minPrice = prices[i];
        } else {
            maxProfit = Math.max(maxProfit, prices[i] - minPrice);
        }
    }
    
    return maxProfit;
}

/**
 * 135. 分发糖果
 * Candy
 */
export function candy(ratings: number[]): number {
    const n = ratings.length;
    const candies = new Array(n).fill(1);
    
    // 从左到右遍历
    for (let i = 1; i < n; i++) {
        if (ratings[i] > ratings[i - 1]) {
            candies[i] = candies[i - 1] + 1;
        }
    }
    
    // 从右到左遍历
    for (let i = n - 2; i >= 0; i--) {
        if (ratings[i] > ratings[i + 1]) {
            candies[i] = Math.max(candies[i], candies[i + 1] + 1);
        }
    }
    
    return candies.reduce((sum, candy) => sum + candy, 0);
}

/**
 * 2740. 找出分区值
 * Find the Value of the Partition
 */
export function findValueOfPartition(nums: number[]): number {
    nums.sort((a, b) => a - b);
    let minDiff = Infinity;
    
    for (let i = 1; i < nums.length; i++) {
        minDiff = Math.min(minDiff, nums[i] - nums[i - 1]);
    }
    
    return minDiff;
}

/**
 * 68. 文本左右对齐
 * Text Justification
 */
export function fullJustify(words: string[], maxWidth: number): string[] {
    const result: string[] = [];
    let currentLine: string[] = [];
    let currentLength = 0;
    
    for (const word of words) {
        // 检查是否可以添加当前单词到当前行
        if (currentLength + word.length + currentLine.length > maxWidth) {
            // 处理当前行
            result.push(justifyLine(currentLine, maxWidth, false));
            currentLine = [word];
            currentLength = word.length;
        } else {
            currentLine.push(word);
            currentLength += word.length;
        }
    }
    
    // 处理最后一行（左对齐）
    if (currentLine.length > 0) {
        result.push(justifyLine(currentLine, maxWidth, true));
    }
    
    return result;
}

function justifyLine(words: string[], maxWidth: number, isLastLine: boolean): string {
    if (words.length === 1 || isLastLine) {
        // 左对齐
        const line = words.join(' ');
        return line + ' '.repeat(maxWidth - line.length);
    }
    
    const totalChars = words.reduce((sum, word) => sum + word.length, 0);
    const totalSpaces = maxWidth - totalChars;
    const gaps = words.length - 1;
    const spacesPerGap = Math.floor(totalSpaces / gaps);
    const extraSpaces = totalSpaces % gaps;
    
    let result = '';
    for (let i = 0; i < words.length - 1; i++) {
        result += words[i];
        result += ' '.repeat(spacesPerGap);
        if (i < extraSpaces) {
            result += ' ';
        }
    }
    result += words[words.length - 1];
    
    return result;
}