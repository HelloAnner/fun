/**
 * 模拟算法相关实现
 * Simulation algorithms
 */

/**
 * 36. 有效的数独
 * Valid Sudoku
 */
export function isValidSudoku(board: string[][]): boolean {
    const rows = Array.from({ length: 9 }, () => new Set<string>());
    const cols = Array.from({ length: 9 }, () => new Set<string>());
    const boxes = Array.from({ length: 9 }, () => new Set<string>());
    
    for (let i = 0; i < 9; i++) {
        for (let j = 0; j < 9; j++) {
            const val = board[i][j];
            if (val === '.') continue;
            
            const boxIndex = Math.floor(i / 3) * 3 + Math.floor(j / 3);
            
            if (rows[i].has(val) || cols[j].has(val) || boxes[boxIndex].has(val)) {
                return false;
            }
            
            rows[i].add(val);
            cols[j].add(val);
            boxes[boxIndex].add(val);
        }
    }
    
    return true;
}

/**
 * 48. 旋转图像
 * Rotate Image
 */
export function rotate(matrix: number[][]): void {
    const n = matrix.length;
    
    // 转置矩阵
    for (let i = 0; i < n; i++) {
        for (let j = i; j < n; j++) {
            [matrix[i][j], matrix[j][i]] = [matrix[j][i], matrix[i][j]];
        }
    }
    
    // 反转每一行
    for (let i = 0; i < n; i++) {
        matrix[i].reverse();
    }
}

/**
 * 54. 螺旋矩阵
 * Spiral Matrix
 */
export function spiralOrder(matrix: number[][]): number[] {
    if (matrix.length === 0) return [];
    
    const result: number[] = [];
    let top = 0;
    let bottom = matrix.length - 1;
    let left = 0;
    let right = matrix[0].length - 1;
    
    while (top <= bottom && left <= right) {
        // 从左到右
        for (let j = left; j <= right; j++) {
            result.push(matrix[top][j]);
        }
        top++;
        
        // 从上到下
        for (let i = top; i <= bottom; i++) {
            result.push(matrix[i][right]);
        }
        right--;
        
        // 从右到左
        if (top <= bottom) {
            for (let j = right; j >= left; j--) {
                result.push(matrix[bottom][j]);
            }
            bottom--;
        }
        
        // 从下到上
        if (left <= right) {
            for (let i = bottom; i >= top; i--) {
                result.push(matrix[i][left]);
            }
            left++;
        }
    }
    
    return result;
}

/**
 * 73. 矩阵置零
 * Set Matrix Zeroes
 */
export function setZeroes(matrix: number[][]): void {
    const m = matrix.length;
    const n = matrix[0].length;
    let firstRowZero = false;
    let firstColZero = false;
    
    // 检查第一行是否有零
    for (let j = 0; j < n; j++) {
        if (matrix[0][j] === 0) {
            firstRowZero = true;
            break;
        }
    }
    
    // 检查第一列是否有零
    for (let i = 0; i < m; i++) {
        if (matrix[i][0] === 0) {
            firstColZero = true;
            break;
        }
    }
    
    // 使用第一行和第一列作为标记
    for (let i = 1; i < m; i++) {
        for (let j = 1; j < n; j++) {
            if (matrix[i][j] === 0) {
                matrix[i][0] = 0;
                matrix[0][j] = 0;
            }
        }
    }
    
    // 根据标记置零
    for (let i = 1; i < m; i++) {
        for (let j = 1; j < n; j++) {
            if (matrix[i][0] === 0 || matrix[0][j] === 0) {
                matrix[i][j] = 0;
            }
        }
    }
    
    // 处理第一行
    if (firstRowZero) {
        for (let j = 0; j < n; j++) {
            matrix[0][j] = 0;
        }
    }
    
    // 处理第一列
    if (firstColZero) {
        for (let i = 0; i < m; i++) {
            matrix[i][0] = 0;
        }
    }
}

/**
 * 274. H 指数
 * H-Index
 */
export function hIndex(citations: number[]): number {
    citations.sort((a, b) => b - a);
    let h = 0;
    
    for (let i = 0; i < citations.length; i++) {
        if (citations[i] >= i + 1) {
            h = i + 1;
        } else {
            break;
        }
    }
    
    return h;
}

/**
 * 1103. 分糖果 II
 * Distribute Candies to People
 */
export function distributeCandies(candies: number, numPeople: number): number[] {
    const result = new Array(numPeople).fill(0);
    let give = 1;
    let person = 0;
    
    while (candies > 0) {
        const toGive = Math.min(give, candies);
        result[person] += toGive;
        candies -= toGive;
        give++;
        person = (person + 1) % numPeople;
    }
    
    return result;
}

/**
 * 2974. 最小数字游戏
 * Minimum Number Game
 */
export function numberGame(nums: number[]): number[] {
    nums.sort((a, b) => a - b);
    const result: number[] = [];
    
    for (let i = 0; i < nums.length; i += 2) {
        result.push(nums[i + 1], nums[i]);
    }
    
    return result;
}

/**
 * 724. 寻找数组的中心下标
 * Find Pivot Index
 */
export function pivotIndex(nums: number[]): number {
    const totalSum = nums.reduce((sum, num) => sum + num, 0);
    let leftSum = 0;
    
    for (let i = 0; i < nums.length; i++) {
        if (leftSum === totalSum - leftSum - nums[i]) {
            return i;
        }
        leftSum += nums[i];
    }
    
    return -1;
}

/**
 * 1329. 将矩阵按对角线排序
 * Sort the Matrix Diagonally
 */
export function diagonalSort(mat: number[][]): number[][] {
    const m = mat.length;
    const n = mat[0].length;
    
    // 排序每条对角线
    for (let k = 0; k < m + n - 1; k++) {
        const diagonal: number[] = [];
        const positions: Array<[number, number]> = [];
        
        // 收集对角线元素
        for (let i = 0; i < m; i++) {
            for (let j = 0; j < n; j++) {
                if (i - j === k - n + 1) {
                    diagonal.push(mat[i][j]);
                    positions.push([i, j]);
                }
            }
        }
        
        // 排序并放回
        diagonal.sort((a, b) => a - b);
        for (let i = 0; i < positions.length; i++) {
            const [row, col] = positions[i];
            mat[row][col] = diagonal[i];
        }
    }
    
    return mat;
}