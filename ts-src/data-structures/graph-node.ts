/**
 * 图节点定义
 * Definition for a graph node
 */
export class GraphNode {
    val: number;
    neighbors: GraphNode[];
    
    constructor(val?: number, neighbors?: GraphNode[]) {
        this.val = val ?? 0;
        this.neighbors = neighbors ?? [];
    }
}