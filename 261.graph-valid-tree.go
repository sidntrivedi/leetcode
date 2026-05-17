package main

/*
You have a graph of n nodes labeled from 0 to n - 1.
You are given an integer n and a list of edges where edges[i] = [ai, bi] indicates that there is
an undirected edge between nodes ai and bi in the graph.

Return true if the edges of the given graph make up a valid tree, and false otherwise.
*/

/*
A graph is a tree if:
1. All the nodes are reachable from a node.
2. There's no cycle.
*/
func validTree(n int, edges [][]int) bool {
	// 1. First check len(edges) == n-1 since in undirected graph, adding one more edge will
	//   create a cycle.
	// 2. Build adjacency list.
	// 3. DFS/BFS from node 0.
	// 4. Return true only if you visited all n nodes.

	if len(edges) != n-1 {
		return false
	}

	// Create adjacency list.
	adj := make([][]int, n)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

	visited := make(map[int]bool)
	var dfs func(node int)
	dfs = func(node int) {
		if visited[node] {
			return
		}
		visited[node] = true

		for _, nei := range adj[node] {
			dfs(nei)
		}
	}

	dfs(0)

	if len(visited) == n {
		return true

	}
	return false
}
