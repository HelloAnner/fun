import { 
    singleNumber,
    singleNumberII,
    singleNumberIII,
    hammingWeight,
    countBits,
    hammingDistance,
    totalHammingDistance,
    reverseBits,
    isPowerOfTwo,
    isPowerOfFour,
    getSum,
    findTheDifference,
    toHex,
    findMaximumXOR,
    findComplement,
    hasAlternatingBits,
    countPrimeSetBits,
    binaryGap,
    minFlips
} from '../bit-manipulation';

describe('Bit Manipulation Algorithms', () => {
    test('singleNumber', () => {
        expect(singleNumber([2, 2, 1])).toBe(1);
        expect(singleNumber([4, 1, 2, 1, 2])).toBe(4);
        expect(singleNumber([1])).toBe(1);
    });
    
    test('singleNumberII', () => {
        expect(singleNumberII([2, 2, 3, 2])).toBe(3);
        expect(singleNumberII([0, 1, 0, 1, 0, 1, 99])).toBe(99);
    });
    
    test('singleNumberIII', () => {
        expect(singleNumberIII([1, 2, 1, 3, 2, 5]).sort()).toEqual([3, 5]);
        expect(singleNumberIII([-1, 0]).sort()).toEqual([-1, 0]);
        expect(singleNumberIII([0, 1]).sort()).toEqual([0, 1]);
    });
    
    test('hammingWeight', () => {
        expect(hammingWeight(0b00000000000000000000000000001011)).toBe(3);
        expect(hammingWeight(0b00000000000000000000000010000000)).toBe(1);
        expect(hammingWeight(0b11111111111111111111111111111101)).toBe(31);
    });
    
    test('countBits', () => {
        expect(countBits(2)).toEqual([0, 1, 1]);
        expect(countBits(5)).toEqual([0, 1, 1, 2, 1, 2]);
    });
    
    test('hammingDistance', () => {
        expect(hammingDistance(1, 4)).toBe(2);
        expect(hammingDistance(3, 1)).toBe(1);
    });
    
    test('totalHammingDistance', () => {
        expect(totalHammingDistance([4, 14, 2])).toBe(6);
        expect(totalHammingDistance([4, 14, 4])).toBe(4);
    });
    
    test('reverseBits', () => {
        expect(reverseBits(0b00000010100101000001111010011100)).toBe(0b00111001011110000010100101000000);
        expect(reverseBits(0b11111111111111111111111111111101)).toBe(0b10111111111111111111111111111111);
    });
    
    test('isPowerOfTwo', () => {
        expect(isPowerOfTwo(1)).toBe(true);
        expect(isPowerOfTwo(16)).toBe(true);
        expect(isPowerOfTwo(3)).toBe(false);
        expect(isPowerOfTwo(4)).toBe(true);
        expect(isPowerOfTwo(5)).toBe(false);
    });
    
    test('isPowerOfFour', () => {
        expect(isPowerOfFour(16)).toBe(true);
        expect(isPowerOfFour(5)).toBe(false);
        expect(isPowerOfFour(1)).toBe(true);
        expect(isPowerOfFour(8)).toBe(false);
    });
    
    test('getSum', () => {
        expect(getSum(1, 2)).toBe(3);
        expect(getSum(2, 3)).toBe(5);
        expect(getSum(-2, 3)).toBe(1);
    });
    
    test('findTheDifference', () => {
        expect(findTheDifference("abcd", "abcde")).toBe("e");
        expect(findTheDifference("", "y")).toBe("y");
        expect(findTheDifference("a", "aa")).toBe("a");
    });
    
    test('toHex', () => {
        expect(toHex(26)).toBe("1a");
        expect(toHex(-1)).toBe("ffffffff");
        expect(toHex(0)).toBe("0");
    });
    
    test('findMaximumXOR', () => {
        expect(findMaximumXOR([3, 10, 5, 25, 2, 8])).toBe(28);
        expect(findMaximumXOR([14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70])).toBe(127);
    });
    
    test('findComplement', () => {
        expect(findComplement(5)).toBe(2);
        expect(findComplement(1)).toBe(0);
    });
    
    test('hasAlternatingBits', () => {
        expect(hasAlternatingBits(5)).toBe(true);
        expect(hasAlternatingBits(7)).toBe(false);
        expect(hasAlternatingBits(11)).toBe(false);
        expect(hasAlternatingBits(10)).toBe(true);
    });
    
    test('countPrimeSetBits', () => {
        expect(countPrimeSetBits(6, 10)).toBe(4);
        expect(countPrimeSetBits(10, 15)).toBe(5);
    });
    
    test('binaryGap', () => {
        expect(binaryGap(22)).toBe(2);
        expect(binaryGap(8)).toBe(0);
        expect(binaryGap(5)).toBe(2);
    });
    
    test('minFlips', () => {
        expect(minFlips(2, 6, 5)).toBe(3);
        expect(minFlips(4, 2, 7)).toBe(1);
        expect(minFlips(1, 2, 3)).toBe(0);
    });
});