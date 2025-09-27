// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/course-schedule/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func canFinish(numCourses int, prerequisites [][]int) bool {
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

@lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	numCourses := Deserialize[int](ReadLine(stdin))
	prerequisites := Deserialize[[][]int](ReadLine(stdin))
	ans := canFinish(numCourses, prerequisites)

	fmt.Println("\noutput:", Serialize(ans))
}
