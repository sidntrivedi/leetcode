package main

/*
Source: https://takeuforward.org/plus/dsa/problems/connected-components

Given a undirected Graph consisting of V vertices numbered from 0 to V-1 and E edges.
The ith edge is represented by [ai,bi], denoting a edge between vertex ai and bi.
We say two vertices u and v belong to a same component if there is a path from u to v
or v to u. Find the number of connected components in the graph.

A connected component is a subgraph of a graph in which there exists a path between
any two vertices, and no vertex of the subgraph shares an edge with a vertex outside of the subgraph.


Examples:
Input: V=4, edges=[[0,1],[1,2]]
Output: 2
Explanation: Vertices {0,1,2} forms the first component and vertex 3 forms the second component.

Input: V = 7, edges = [[0, 1], [1, 2], [2, 3], [4, 5]]
Output: 3
Explanation:
The edges [0, 1], [1, 2], [2, 3] form a connected component with vertices {0, 1, 2, 3}.
The edge [4, 5] forms another connected component with vertices {4, 5}.
Therefore, the graph has 3 connected components: {0, 1, 2, 3}, {4, 5}, and the isolated
vertices {6} (vertices 6 and any other unconnected vertices).
*/

func findNumberOfComponent(V int, edges [][]int) int {
	adj := make(map[int][]int, 0)
	visited := make([]bool, V)

	// Create adjacency list.
	for i := 0; i < len(edges); i++ {
		from := edges[i][0]
		to := edges[i][1]

		// Undirected.
		adj[from] = append(adj[from], to)
		adj[to] = append(adj[to], from)
	}

	components := 0
	// Start DFS from the first node and keep on marking visited array.
	var dfs func(node int)
	dfs = func(node int) {
		visited[node] = true

		for _, neighbour := range adj[node] {
			if !visited[neighbour] {
				dfs(neighbour)
			}
		}
	}

	for n := 0; n < V; n++ {
		if !visited[n] {
			components++
			dfs(n)
		}
	}

	return components
}
