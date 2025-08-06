/**
 * 字符串相关算法实现
 * String related algorithms
 */

/**
 * 28. 找出字符串中第一个匹配项的下标
 * Find the Index of the First Occurrence in a String
 */
export function strStr(haystack: string, needle: string): number {
    if (needle.length === 0) return 0;
    if (needle.length > haystack.length) return -1;
    
    for (let i = 0; i <= haystack.length - needle.length; i++) {
        if (haystack.substring(i, i + needle.length) === needle) {
            return i;
        }
    }
    
    return -1;
}

/**
 * 6. Z 字形变换
 * Zigzag Conversion
 */
export function convert(s: string, numRows: number): string {
    if (numRows === 1) return s;
    
    const rows: string[] = new Array(Math.min(numRows, s.length)).fill('');
    let currentRow = 0;
    let goingDown = false;
    
    for (const char of s) {
        rows[currentRow] += char;
        
        if (currentRow === 0 || currentRow === numRows - 1) {
            goingDown = !goingDown;
        }
        
        currentRow += goingDown ? 1 : -1;
    }
    
    return rows.join('');
}

/**
 * 65. 有效数字
 * Valid Number
 */
export function isNumber(s: string): boolean {
    s = s.trim();
    let hasNum = false;
    let hasE = false;
    let hasDot = false;
    let hasSign = false;
    
    for (let i = 0; i < s.length; i++) {
        const char = s[i];
        
        if (char >= '0' && char <= '9') {
            hasNum = true;
        } else if (char === '.') {
            if (hasDot || hasE) return false;
            hasDot = true;
        } else if (char === 'e' || char === 'E') {
            if (hasE || !hasNum) return false;
            hasE = true;
            hasNum = false; // e后面必须有数字
        } else if (char === '+' || char === '-') {
            if (i !== 0 && s[i - 1] !== 'e' && s[i - 1] !== 'E') return false;
        } else {
            return false;
        }
    }
    
    return hasNum;
}

/**
 * 字符串翻转检查
 * Check if string can be shifted to match another
 */
export function canShift(a: string, b: string): boolean {
    if (a.length !== b.length) return false;
    return (a + a).includes(b);
}

/**
 * 字符串倒序修改
 * Reverse string modification
 */
export function reverseStringModify(s: string): string {
    const chars = s.split('');
    let left = 0;
    let right = chars.length - 1;
    
    while (left < right) {
        [chars[left], chars[right]] = [chars[right], chars[left]];
        left++;
        right--;
    }
    
    return chars.join('');
}

/**
 * 字母异位词分组
 * Group Anagrams
 */
export function groupAnagrams(strs: string[]): string[][] {
    const map = new Map<string, string[]>();
    
    for (const str of strs) {
        const key = str.split('').sort().join('');
        if (!map.has(key)) {
            map.set(key, []);
        }
        map.get(key)!.push(str);
    }
    
    return Array.from(map.values());
}

/**
 * 205. 同构字符串
 * Isomorphic Strings
 */
export function isIsomorphic(s: string, t: string): boolean {
    if (s.length !== t.length) return false;
    
    const mapS = new Map<string, string>();
    const mapT = new Map<string, string>();
    
    for (let i = 0; i < s.length; i++) {
        const charS = s[i];
        const charT = t[i];
        
        if (mapS.has(charS)) {
            if (mapS.get(charS) !== charT) return false;
        } else {
            mapS.set(charS, charT);
        }
        
        if (mapT.has(charT)) {
            if (mapT.get(charT) !== charS) return false;
        } else {
            mapT.set(charT, charS);
        }
    }
    
    return true;
}

/**
 * 242. 有效的字母异位词
 * Valid Anagram
 */
export function isAnagram(s: string, t: string): boolean {
    if (s.length !== t.length) return false;
    
    const count = new Map<string, number>();
    
    for (const char of s) {
        count.set(char, (count.get(char) || 0) + 1);
    }
    
    for (const char of t) {
        if (!count.has(char)) return false;
        count.set(char, count.get(char)! - 1);
        if (count.get(char) === 0) {
            count.delete(char);
        }
    }
    
    return count.size === 0;
}

/**
 * 290. 单词规律
 * Word Pattern
 */
export function wordPattern(pattern: string, s: string): boolean {
    const words = s.split(' ');
    if (pattern.length !== words.length) return false;
    
    const charToWord = new Map<string, string>();
    const wordToChar = new Map<string, string>();
    
    for (let i = 0; i < pattern.length; i++) {
        const char = pattern[i];
        const word = words[i];
        
        if (charToWord.has(char)) {
            if (charToWord.get(char) !== word) return false;
        } else {
            charToWord.set(char, word);
        }
        
        if (wordToChar.has(word)) {
            if (wordToChar.get(word) !== char) return false;
        } else {
            wordToChar.set(word, char);
        }
    }
    
    return true;
}