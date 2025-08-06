/**
 * 排序相关算法实现
 * Sorting related algorithms
 */

/**
 * 快速排序
 * Quick Sort
 */
export function quickSort(arr: number[]): number[] {
    if (arr.length <= 1) return arr;
    
    const pivot = arr[Math.floor(arr.length / 2)];
    const left = arr.filter(x => x < pivot);
    const middle = arr.filter(x => x === pivot);
    const right = arr.filter(x => x > pivot);
    
    return [...quickSort(left), ...middle, ...quickSort(right)];
}

/**
 * 归并排序
 * Merge Sort
 */
export function mergeSort(arr: number[]): number[] {
    if (arr.length <= 1) return arr;
    
    const mid = Math.floor(arr.length / 2);
    const left = mergeSort(arr.slice(0, mid));
    const right = mergeSort(arr.slice(mid));
    
    return merge(left, right);
}

function merge(left: number[], right: number[]): number[] {
    const result: number[] = [];
    let i = 0, j = 0;
    
    while (i < left.length && j < right.length) {
        if (left[i] <= right[j]) {
            result.push(left[i]);
            i++;
        } else {
            result.push(right[j]);
            j++;
        }
    }
    
    return result.concat(left.slice(i)).concat(right.slice(j));
}

/**
 * 堆排序
 * Heap Sort
 */
export function heapSort(arr: number[]): number[] {
    const result = [...arr];
    const n = result.length;
    
    // 构建最大堆
    for (let i = Math.floor(n / 2) - 1; i >= 0; i--) {
        heapify(result, n, i);
    }
    
    // 逐个提取元素
    for (let i = n - 1; i > 0; i--) {
        [result[0], result[i]] = [result[i], result[0]];
        heapify(result, i, 0);
    }
    
    return result;
}

function heapify(arr: number[], n: number, i: number): void {
    let largest = i;
    const left = 2 * i + 1;
    const right = 2 * i + 2;
    
    if (left < n && arr[left] > arr[largest]) {
        largest = left;
    }
    
    if (right < n && arr[right] > arr[largest]) {
        largest = right;
    }
    
    if (largest !== i) {
        [arr[i], arr[largest]] = [arr[largest], arr[i]];
        heapify(arr, n, largest);
    }
}

/**
 * 冒泡排序
 * Bubble Sort
 */
export function bubbleSort(arr: number[]): number[] {
    const result = [...arr];
    const n = result.length;
    
    for (let i = 0; i < n - 1; i++) {
        for (let j = 0; j < n - i - 1; j++) {
            if (result[j] > result[j + 1]) {
                [result[j], result[j + 1]] = [result[j + 1], result[j]];
            }
        }
    }
    
    return result;
}

/**
 * 插入排序
 * Insertion Sort
 */
export function insertionSort(arr: number[]): number[] {
    const result = [...arr];
    
    for (let i = 1; i < result.length; i++) {
        const key = result[i];
        let j = i - 1;
        
        while (j >= 0 && result[j] > key) {
            result[j + 1] = result[j];
            j--;
        }
        
        result[j + 1] = key;
    }
    
    return result;
}

/**
 * 选择排序
 * Selection Sort
 */
export function selectionSort(arr: number[]): number[] {
    const result = [...arr];
    const n = result.length;
    
    for (let i = 0; i < n - 1; i++) {
        let minIdx = i;
        
        for (let j = i + 1; j < n; j++) {
            if (result[j] < result[minIdx]) {
                minIdx = j;
            }
        }
        
        if (minIdx !== i) {
            [result[i], result[minIdx]] = [result[minIdx], result[i]];
        }
    }
    
    return result;
}

/**
 * 计数排序
 * Counting Sort
 */
export function countingSort(arr: number[]): number[] {
    if (arr.length === 0) return arr;
    
    const max = Math.max(...arr);
    const min = Math.min(...arr);
    const range = max - min + 1;
    const count = new Array(range).fill(0);
    const result = new Array(arr.length);
    
    // 计数
    for (const num of arr) {
        count[num - min]++;
    }
    
    // 累积计数
    for (let i = 1; i < range; i++) {
        count[i] += count[i - 1];
    }
    
    // 构建结果数组
    for (let i = arr.length - 1; i >= 0; i--) {
        result[count[arr[i] - min] - 1] = arr[i];
        count[arr[i] - min]--;
    }
    
    return result;
}