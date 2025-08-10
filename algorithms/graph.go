package algorithms

import (
	"algorithms-go/data_structures"
)

/**
 * 图相关算法实现
 * Graph related algorithms
 */

/**
 * 133. 克隆图
 * Clone Graph
 */
func CloneGraph(node *data_structures.GraphNode) *data_structures.GraphNode {
	if node == nil {
		return nil
	}

	visited := make(map[*data_structures.GraphNode]*data_structures.GraphNode)

	var dfs func(*data_structures.GraphNode) *data_structures.GraphNode
	dfs = func(node *data_structures.GraphNode) *data_structures.GraphNode {
		if clone, exists := visited[node]; exists {
			return clone
		}

		clone := data_structures.NewGraphNode(node.Val)
		visited[node] = clone

		for _, neighbor := range node.Neighbors {
			clone.Neighbors = append(clone.Neighbors, dfs(neighbor))
		}

		return clone
	}

	return dfs(node)
}

/**
 * 752. 打开转盘锁
 * Open the Lock
 */
func OpenLock(deadends []string, target string) int {
	deadSet := make(map[string]bool)
	for _, dead := range deadends {
		deadSet[dead] = true
	}

	if deadSet["0000"] {
		return -1
	}
	if target == "0000" {
		return 0
	}

	queue := []string{"0000"}
	visited := map[string]bool{"0000": true}
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		steps++

		for i := 0; i < size; i++ {
			current := queue[0]
			queue = queue[1:]

			for _, next := range getNextStates(current) {
				if next == target {
					return steps
				}
				if !deadSet[next] && !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}
	}

	return -1
}

func getNextStates(state string) []string {
	result := []string{}
	chars := []rune(state)

	for i := 0; i < 4; i++ {
		digit := int(chars[i] - '0')

		// 向上转动
		chars[i] = rune('0' + (digit+1)%10)
		result = append(result, string(chars))

		// 向下转动
		chars[i] = rune('0' + (digit+9)%10)
		result = append(result, string(chars))

		// 恢复原值
		chars[i] = rune('0' + digit)
	}

	return result
}

/**
 * 547. 省份数量
 * Number of Provinces
 */
func FindCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)
	provinces := 0

	var dfs func(int)
	dfs = func(i int) {
		visited[i] = true
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 && !visited[j] {
				dfs(j)
			}
		}
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i)
			provinces++
		}
	}

	return provinces
}

/**
 * 207. 课程表
 * Course Schedule
 */
func CanFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	// 构建图和入度数组
	for _, prereq := range prerequisites {
		course, pre := prereq[0], prereq[1]
		graph[pre] = append(graph[pre], course)
		inDegree[course]++
	}

	// 找到所有入度为0的课程
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	completed := 0

	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		completed++

		for _, nextCourse := range graph[course] {
			inDegree[nextCourse]--
			if inDegree[nextCourse] == 0 {
				queue = append(queue, nextCourse)
			}
		}
	}

	return completed == numCourses
}

/**
 * 1971. 寻找图中是否存在路径
 * Find if Path Exists in Graph
 */
func ValidPath(n int, edges [][]int, source, destination int) bool {
	if source == destination {
		return true
	}

	graph := make([][]int, n)

	// 构建邻接表
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	visited := make([]bool, n)

	var dfs func(int) bool
	dfs = func(node int) bool {
		if node == destination {
			return true
		}

		visited[node] = true

		for _, neighbor := range graph[node] {
			if !visited[neighbor] && dfs(neighbor) {
				return true
			}
		}

		return false
	}

	return dfs(source)
}