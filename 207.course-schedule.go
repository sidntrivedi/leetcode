package main

/*
 * @lc app=leetcode id=207 lang=golang
 *
 * [207] Course Schedule
 *
 * https://leetcode.com/problems/course-schedule/description/
 *
 * algorithms
 * Medium (51.37%)
 * Likes:    18032
 * Dislikes: 872
 * Total Accepted:    2.6M
 * Total Submissions: 5.1M
 * Testcase Example:  '2\n[[1,0]]'
 *
 * There are a total of numCourses courses you have to take, labeled from 0 to
 * numCourses - 1. You are given an array prerequisites where prerequisites[i]
 * = [ai, bi] indicates that you must take course bi first if you want to take
 * course ai.
 *
 *
 * For example, the pair [0, 1], indicates that to take course 0 you have to
 * first take course 1.
 *
 *
 * Return true if you can finish all courses. Otherwise, return false.
 *
 *
 * Example 1:
 *
 *
 * Input: numCourses = 2, prerequisites = [[1,0]]
 * Output: true
 * Explanation: There are a total of 2 courses to take.
 * To take course 1 you should have finished course 0. So it is possible.
 *
 *
 * Example 2:
 *
 *
 * Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
 * Output: false
 * Explanation: There are a total of 2 courses to take.
 * To take course 1 you should have finished course 0, and to take course 0 you
 * should also have finished course 1. So it is impossible.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= numCourses <= 2000
 * 0 <= prerequisites.length <= 5000
 * prerequisites[i].length == 2
 * 0 <= ai, bi < numCourses
 * All the pairs prerequisites[i] are unique.
 *
 *
 */

// @lc code=start

/*
1. Iterate all the courses from 0 to numCourses-1
2. For each course, do a DFS
3. DUring DFS and its neighbours DFS, keep status of node as 1.
4. Once DFS completes for a node return status 2 and true
5. If any false is received, return false
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	state := make([]int, numCourses)

	// Create the graph adjacency list and
	// populate the state slice.
	for _, p := range prerequisites {
		course := p[0]
		prereq := p[1]
		graph[prereq] = append(graph[prereq], course)
	}

	var dfs func(course int) bool
	dfs = func(course int) bool {
		if state[course] == 1 {
			return false
		}
		if state[course] == 2 {
			return true
		}

		state[course] = 1

		for _, next := range graph[course] {
			if !dfs(next) {
				return false
			}
		}

		state[course] = 2
		return true
	}

	for course := 0; course < numCourses; course++ {
		if !dfs(course) {
			return false
		}
	}
	return true
}

// @lc code=end
