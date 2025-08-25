package data_structures

/**
 * 图节点定义
 * Definition for a graph node
 */
type GraphNode struct {
	Val       int          `json:"val"`
	Neighbors []*GraphNode `json:"neighbors"`
}

// NewGraphNode 创建新的图节点
func NewGraphNode(val int) *GraphNode {
	return &GraphNode{
		Val:       val,
		Neighbors: make([]*GraphNode, 0),
	}
}

// NewGraphNodeWithNeighbors 创建带邻居节点的图节点
func NewGraphNodeWithNeighbors(val int, neighbors []*GraphNode) *GraphNode {
	return &GraphNode{
		Val:       val,
		Neighbors: neighbors,
	}
}

// AddNeighbor 添加邻居节点
func (g *GraphNode) AddNeighbor(neighbor *GraphNode) {
	g.Neighbors = append(g.Neighbors, neighbor)
}
