import { 
    lengthOfLongestSubstring,
    minWindow,
    minSubArrayLen,
    containsNearbyDuplicate,
    findAnagrams,
    maxSatisfied,
    maxConsecutiveAnswers,
    minOperations,
    longestEqualSubarray
} from '../sliding-window';

describe('Sliding Window Algorithms', () => {
    test('lengthOfLongestSubstring', () => {
        expect(lengthOfLongestSubstring("abcabcbb")).toBe(3);
        expect(lengthOfLongestSubstring("bbbbb")).toBe(1);
        expect(lengthOfLongestSubstring("pwwkew")).toBe(3);
        expect(lengthOfLongestSubstring("")).toBe(0);
        expect(lengthOfLongestSubstring("au")).toBe(2);
    });
    
    test('minWindow', () => {
        expect(minWindow("ADOBECODEBANC", "ABC")).toBe("BANC");
        expect(minWindow("a", "a")).toBe("a");
        expect(minWindow("a", "aa")).toBe("");
        expect(minWindow("ab", "b")).toBe("b");
    });
    
    test('minSubArrayLen', () => {
        expect(minSubArrayLen(7, [2, 3, 1, 2, 4, 3])).toBe(2);
        expect(minSubArrayLen(4, [1, 4, 4])).toBe(1);
        expect(minSubArrayLen(11, [1, 1, 1, 1, 1, 1, 1, 1])).toBe(0);
        expect(minSubArrayLen(15, [1, 2, 3, 4, 5])).toBe(5);
    });
    
    test('containsNearbyDuplicate', () => {
        expect(containsNearbyDuplicate([1, 2, 3, 1], 3)).toBe(true);
        expect(containsNearbyDuplicate([1, 0, 1, 1], 1)).toBe(true);
        expect(containsNearbyDuplicate([1, 2, 3, 1, 2, 3], 2)).toBe(false);
    });
    
    test('findAnagrams', () => {
        expect(findAnagrams("abab", "ab")).toEqual([0, 2]);
        expect(findAnagrams("abacbabc", "abc")).toEqual([1, 2, 3, 5]);
        expect(findAnagrams("aab", "ab")).toEqual([1]);
    });
    
    test('maxSatisfied', () => {
        expect(maxSatisfied([1, 0, 1, 2, 1, 1, 7, 5], [0, 1, 0, 1, 0, 1, 0, 1], 3)).toBe(16);
        expect(maxSatisfied([1], [0], 1)).toBe(1);
        expect(maxSatisfied([4, 10, 10], [1, 1, 0], 2)).toBe(24);
    });
    
    test('maxConsecutiveAnswers', () => {
        expect(maxConsecutiveAnswers("TTFF", 2)).toBe(4);
        expect(maxConsecutiveAnswers("TFFT", 1)).toBe(3);
        expect(maxConsecutiveAnswers("TTFTTFTT", 1)).toBe(5);
    });
    
    test('minOperations', () => {
        expect(minOperations([4, 2, 5, 3])).toBe(0);
        expect(minOperations([1, 2, 3, 5, 6])).toBe(1);
        expect(minOperations([1, 10, 100, 1000])).toBe(3);
    });
    
    test('longestEqualSubarray', () => {
        expect(longestEqualSubarray([1, 3, 2, 3, 1, 3], 3)).toBe(3);
        expect(longestEqualSubarray([1, 1, 2, 2, 1, 1], 2)).toBe(4);
        expect(longestEqualSubarray([1, 2, 1], 0)).toBe(1);
    });
});