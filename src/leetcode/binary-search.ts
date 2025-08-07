/**
 * 二分查找相关算法实现
 * Binary Search related algorithms
 */

/**
 * 704. 二分查找
 * Binary Search
 */
export function search(nums: number[], target: number): number {
    let left = 0;
    let right = nums.length - 1;
    
    while (left <= right) {
        const mid = Math.floor((left + right) / 2);
        
        if (nums[mid] === target) {
            return mid;
        } else if (nums[mid] < target) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    
    return -1;
}

/**
 * 69. x 的平方根
 * Sqrt(x)
 */
export function mySqrt(x: number): number {
    if (x === 0) return 0;
    
    let left = 1;
    let right = x;
    
    while (left <= right) {
        const mid = Math.floor((left + right) / 2);
        const square = mid * mid;
        
        if (square === x) {
            return mid;
        } else if (square < x) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    
    return right;
}

/**
 * 35. 搜索插入位置
 * Search Insert Position
 */
export function searchInsert(nums: number[], target: number): number {
    let left = 0;
    let right = nums.length - 1;
    
    while (left <= right) {
        const mid = Math.floor((left + right) / 2);
        
        if (nums[mid] === target) {
            return mid;
        } else if (nums[mid] < target) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    
    return left;
}

/**
 * 34. 在排序数组中查找元素的第一个和最后一个位置
 * Find First and Last Position of Element in Sorted Array
 */
export function searchRange(nums: number[], target: number): number[] {
    const findFirst = (nums: number[], target: number): number => {
        let left = 0;
        let right = nums.length - 1;
        let result = -1;
        
        while (left <= right) {
            const mid = Math.floor((left + right) / 2);
            
            if (nums[mid] === target) {
                result = mid;
                right = mid - 1; // 继续向左查找
            } else if (nums[mid] < target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        
        return result;
    };
    
    const findLast = (nums: number[], target: number): number => {
        let left = 0;
        let right = nums.length - 1;
        let result = -1;
        
        while (left <= right) {
            const mid = Math.floor((left + right) / 2);
            
            if (nums[mid] === target) {
                result = mid;
                left = mid + 1; // 继续向右查找
            } else if (nums[mid] < target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        
        return result;
    };
    
    return [findFirst(nums, target), findLast(nums, target)];
}

/**
 * 3143. 正方形中的最多点数
 * Maximum Points Inside the Square
 */
export function maxPointsInsideSquare(points: number[][], s: string): number {
    const n = points.length;
    const distances: Array<{ dist: number, char: string }> = [];
    
    for (let i = 0; i < n; i++) {
        const [x, y] = points[i];
        const dist = Math.max(Math.abs(x), Math.abs(y));
        distances.push({ dist, char: s[i] });
    }
    
    distances.sort((a, b) => a.dist - b.dist);
    
    let left = 0;
    let right = distances[n - 1].dist;
    let result = 0;
    
    while (left <= right) {
        const mid = Math.floor((left + right) / 2);
        
        if (isValidSquareSize(distances, mid)) {
            result = countPointsInSquare(distances, mid);
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    
    return result;
}

function isValidSquareSize(distances: Array<{ dist: number, char: string }>, size: number): boolean {
    const seen = new Set<string>();
    
    for (const { dist, char } of distances) {
        if (dist <= size) {
            if (seen.has(char)) {
                return false;
            }
            seen.add(char);
        }
    }
    
    return true;
}

function countPointsInSquare(distances: Array<{ dist: number, char: string }>, size: number): number {
    const seen = new Set<string>();
    
    for (const { dist, char } of distances) {
        if (dist <= size) {
            seen.add(char);
        }
    }
    
    return seen.size;
}