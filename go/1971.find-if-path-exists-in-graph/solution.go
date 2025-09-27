// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/find-if-path-exists-in-graph/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func validPath(n int, edges [][]int, source int, destination int) bool {
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

@lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	edges := Deserialize[[][]int](ReadLine(stdin))
	source := Deserialize[int](ReadLine(stdin))
	destination := Deserialize[int](ReadLine(stdin))
	ans := validPath(n, edges, source, destination)

	fmt.Println("\noutput:", Serialize(ans))
}
