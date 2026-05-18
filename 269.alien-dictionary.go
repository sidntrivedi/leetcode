package main

/*
There is a new alien language that uses the English alphabet. However, the order of the letters is unknown to you.

You are given a list of strings words from the alien language's dictionary.
Now it is claimed that the strings in words are sorted lexicographically by the rules of this new language.

If this claim is incorrect, and the given arrangement of string in words cannot
correspond to any order of letters, return "".

Otherwise, return a string of the unique letters in the new alien language sorted in
lexicographically increasing order by the new language's rules. If there are multiple solutions, return any of them.

Example 1:

Input: words = ["wrt","wrf","er","ett","rftt"]
Output: "wertf"
Example 2:

Input: words = ["z","x"]
Output: "zx"
Example 3:

Input: words = ["z","x","z"]
Output: ""
Explanation: The order is invalid, so return "".
*/

/*
Approach:
 1. Match the words in pairs and create a DAG out of it for
    each character.
 2. Once you have the DAG, get the ordering and do a topological
    sort over it.
*/
func alienOrder(words []string) string {

	// Create adjacency map.
	adj := make(map[byte]map[byte]bool)
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			if _, ok := adj[word[i]]; !ok {
				adj[word[i]] = make(map[byte]bool)
			}
		}
	}

	for i := 0; i < len(words)-1; i++ {
		currWd := words[i]
		nextWd := words[i+1]

		minLen := min(len(currWd), len(nextWd))
		foundOrdering := false

		for j := 0; j < minLen; j++ {
			if currWd[j] != nextWd[j] {
				from := currWd[j]
				to := nextWd[j]

				adj[from][to] = true
				foundOrdering = true
				break
			}
		}

		if !foundOrdering && len(currWd) > len(nextWd) {
			return ""
		}
	}

	// Now that we have the adjacency list, lets create
	// a ordering and then use topological sort.
	return topoSort(adj)
}

// topoSort returns the topological sort of the given graph.
// Uses BFS.
func topoSort(adj map[byte]map[byte]bool) string {
	indegree := make(map[byte]int)
	for ch := range adj {
		indegree[ch] = 0
	}

	for _, neighbors := range adj {
		for nei := range neighbors {
			indegree[nei]++
		}
	}

	queue := []byte{}
	for ch, degree := range indegree {
		if degree == 0 {
			queue = append(queue, ch)
		}
	}

	order := []byte{}
	for len(queue) > 0 {
		ch := queue[0]
		queue = queue[1:]
		order = append(order, ch)

		for nei := range adj[ch] {
			indegree[nei]--
			if indegree[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}

	if len(order) != len(indegree) {
		return ""
	}
	return string(order)
}
