/**
 * 图相关算法实现
 * Graph related algorithms
 */

import { GraphNode } from '../data-structures/graph-node';

/**
 * 133. 克隆图
 * Clone Graph
 */
export function cloneGraph(node: GraphNode | null): GraphNode | null {
    if (!node) return null;
    
    const visited = new Map<GraphNode, GraphNode>();
    
    function dfs(node: GraphNode): GraphNode {
        if (visited.has(node)) {
            return visited.get(node)!;
        }
        
        const clone = new GraphNode(node.val);
        visited.set(node, clone);
        
        for (const neighbor of node.neighbors) {
            clone.neighbors.push(dfs(neighbor));
        }
        
        return clone;
    }
    
    return dfs(node);
}

/**
 * 752. 打开转盘锁
 * Open the Lock
 */
export function openLock(deadends: string[], target: string): number {
    const deadSet = new Set(deadends);
    if (deadSet.has("0000")) return -1;
    if (target === "0000") return 0;
    
    const queue: string[] = ["0000"];
    const visited = new Set<string>(["0000"]);
    let steps = 0;
    
    while (queue.length > 0) {
        const size = queue.length;
        steps++;
        
        for (let i = 0; i < size; i++) {
            const current = queue.shift()!;
            
            for (const next of getNextStates(current)) {
                if (next === target) return steps;
                if (!deadSet.has(next) && !visited.has(next)) {
                    visited.add(next);
                    queue.push(next);
                }
            }
        }
    }
    
    return -1;
}

function getNextStates(state: string): string[] {
    const result: string[] = [];
    const chars = state.split('');
    
    for (let i = 0; i < 4; i++) {
        const digit = parseInt(chars[i]);
        
        // 向上转动
        chars[i] = ((digit + 1) % 10).toString();
        result.push(chars.join(''));
        
        // 向下转动
        chars[i] = ((digit + 9) % 10).toString();
        result.push(chars.join(''));
        
        // 恢复原值
        chars[i] = digit.toString();
    }
    
    return result;
}

/**
 * 547. 省份数量
 * Number of Provinces
 */
export function findCircleNum(isConnected: number[][]): number {
    const n = isConnected.length;
    const visited = new Array(n).fill(false);
    let provinces = 0;
    
    function dfs(i: number): void {
        visited[i] = true;
        for (let j = 0; j < n; j++) {
            if (isConnected[i][j] === 1 && !visited[j]) {
                dfs(j);
            }
        }
    }
    
    for (let i = 0; i < n; i++) {
        if (!visited[i]) {
            dfs(i);
            provinces++;
        }
    }
    
    return provinces;
}

/**
 * 207. 课程表
 * Course Schedule
 */
export function canFinish(numCourses: number, prerequisites: number[][]): boolean {
    const graph: number[][] = Array.from({ length: numCourses }, () => []);
    const inDegree = new Array(numCourses).fill(0);
    
    // 构建图和入度数组
    for (const [course, prereq] of prerequisites) {
        graph[prereq].push(course);
        inDegree[course]++;
    }
    
    // 找到所有入度为0的课程
    const queue: number[] = [];
    for (let i = 0; i < numCourses; i++) {
        if (inDegree[i] === 0) {
            queue.push(i);
        }
    }
    
    let completed = 0;
    
    while (queue.length > 0) {
        const course = queue.shift()!;
        completed++;
        
        for (const nextCourse of graph[course]) {
            inDegree[nextCourse]--;
            if (inDegree[nextCourse] === 0) {
                queue.push(nextCourse);
            }
        }
    }
    
    return completed === numCourses;
}

/**
 * 1971. 寻找图中是否存在路径
 * Find if Path Exists in Graph
 */
export function validPath(n: number, edges: number[][], source: number, destination: number): boolean {
    if (source === destination) return true;
    
    const graph: number[][] = Array.from({ length: n }, () => []);
    
    // 构建邻接表
    for (const [u, v] of edges) {
        graph[u].push(v);
        graph[v].push(u);
    }
    
    const visited = new Array(n).fill(false);
    
    function dfs(node: number): boolean {
        if (node === destination) return true;
        
        visited[node] = true;
        
        for (const neighbor of graph[node]) {
            if (!visited[neighbor] && dfs(neighbor)) {
                return true;
            }
        }
        
        return false;
    }
    
    return dfs(source);
}