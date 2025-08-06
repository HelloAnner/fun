/**
 * 哈希算法测试
 * Hash algorithms tests
 */

import { relocateMarbles, twoSum, isHappy, findJudge, destCity } from '../hash';

describe('Hash Algorithms', () => {
    describe('relocateMarbles', () => {
        test('should relocate marbles correctly', () => {
            expect(relocateMarbles([1, 6, 7, 8], [1, 7, 2], [2, 9, 5])).toEqual([5, 6, 8, 9]);
            expect(relocateMarbles([1, 1, 3, 3], [1, 3], [2, 2])).toEqual([2]);
        });
    });

    describe('twoSum', () => {
        test('should find two numbers that add up to target', () => {
            expect(twoSum([2, 7, 11, 15], 9)).toEqual([0, 1]);
            expect(twoSum([3, 2, 4], 6)).toEqual([1, 2]);
            expect(twoSum([3, 3], 6)).toEqual([0, 1]);
        });
    });

    describe('isHappy', () => {
        test('should determine if a number is happy', () => {
            expect(isHappy(19)).toBe(true);
            expect(isHappy(2)).toBe(false);
            expect(isHappy(1)).toBe(true);
        });
    });

    describe('findJudge', () => {
        test('should find the town judge', () => {
            expect(findJudge(2, [[1, 2]])).toBe(2);
            expect(findJudge(3, [[1, 3], [2, 3]])).toBe(3);
            expect(findJudge(3, [[1, 3], [2, 3], [3, 1]])).toBe(-1);
        });
    });

    describe('destCity', () => {
        test('should find destination city', () => {
            expect(destCity([["London", "New York"], ["New York", "Lima"], ["Lima", "Sao Paulo"]])).toBe("Sao Paulo");
            expect(destCity([["B", "C"], ["D", "B"], ["C", "A"]])).toBe("A");
        });
    });
});