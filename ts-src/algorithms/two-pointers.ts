/**
 * 双指针相关算法实现
 * Two pointers related algorithms
 */

/**
 * 977. 有序数组的平方
 * Squares of a Sorted Array
 */
export function sortedSquares(nums: number[]): number[] {
    const n = nums.length;
    const result = new Array(n);
    let left = 0;
    let right = n - 1;
    
    for (let i = n - 1; i >= 0; i--) {
        const leftSquare = nums[left] * nums[left];
        const rightSquare = nums[right] * nums[right];
        
        if (leftSquare > rightSquare) {
            result[i] = leftSquare;
            left++;
        } else {
            result[i] = rightSquare;
            right--;
        }
    }
    
    return result;
}

/**
 * 88. 合并两个有序数组
 * Merge Sorted Array
 */
export function merge(nums1: number[], m: number, nums2: number[], n: number): void {
    let i = m - 1;
    let j = n - 1;
    let k = m + n - 1;
    
    while (i >= 0 && j >= 0) {
        if (nums1[i] > nums2[j]) {
            nums1[k] = nums1[i];
            i--;
        } else {
            nums1[k] = nums2[j];
            j--;
        }
        k--;
    }
    
    while (j >= 0) {
        nums1[k] = nums2[j];
        j--;
        k--;
    }
}

/**
 * 125. 验证回文串
 * Valid Palindrome
 */
export function isPalindrome(s: string): boolean {
    const cleaned = s.toLowerCase().replace(/[^a-z0-9]/g, '');
    let left = 0;
    let right = cleaned.length - 1;
    
    while (left < right) {
        if (cleaned[left] !== cleaned[right]) {
            return false;
        }
        left++;
        right--;
    }
    
    return true;
}

/**
 * 392. 判断子序列
 * Is Subsequence
 */
export function isSubsequence(s: string, t: string): boolean {
    let i = 0;
    let j = 0;
    
    while (i < s.length && j < t.length) {
        if (s[i] === t[j]) {
            i++;
        }
        j++;
    }
    
    return i === s.length;
}

/**
 * 151. 反转字符串中的单词
 * Reverse Words in a String
 */
export function reverseWords(s: string): string {
    return s.trim().split(/\s+/).reverse().join(' ');
}