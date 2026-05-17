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
	adj := make(map[byte][]byte, 0)
	for i := 0; i < len(words)-1; i++ {
		currWd := words[i]
		nextWd := words[i+1]

		minLen := min(len(currWd), len(nextWd))

		for j := 0; j < minLen; j++ {
			if currWd[j] != nextWd[j] {
				from := currWd[j]
				to := nextWd[j]

				adj[from] = append(adj[from], to)
			}
		}
	}
}
