package hard

import "fmt"

func RegularExpMatching(s string, p string) bool {

	type Pair struct {
		A int
		B int
	}
	//? top down memoization: cache traversal results so we can exit out of recursive calls that call a cached path traversal
	//! worst case O(2^N) --> O(N*M) where N is length of the string and M is the length of the pattern
	// there can be only m * n unique sub problems and each of these sub problem are calculated exactly once.
	cache := make(map[Pair]bool)

	var dfs func(i int, j int) bool
	dfs = func(i int, j int) bool {
		fmt.Println("i, j", i, j)
		//? check cache if this traversal has been completed before
		value, exists := cache[Pair{i, j}]
		if exists {
			fmt.Println("cache found for", i, j)
			return value
		}

		// base case: end of pattern - we need s to be fully matched too
		if j >= len(p) {
			return i >= len(s)
		}

		// check if current characters match (and ensure i is in bounds)
		match := i < len(s) && (p[j] == '.' || s[i] == p[j])

		//! whenever we encounter a *: we have two un-skippable paths as either path can lead to a match:
		//* 1.keep repeating the character
		//* 2. ignore character and move onto next character in the pattern
		// handle case with '*'
		if j+1 < len(p) && p[j+1] == '*' {
			//* if current character match we skip
			//* if current character does not match we also skip
			//? current match status does not matter for this branch
			skipPattern := dfs(i, j+2)
			//?* current pattern has to match in this case if we want to continue repeating character
			usePattern := match && dfs(i+1, j)
			cache[Pair{i, j}] = skipPattern || usePattern
			return cache[Pair{i, j}]
		}

		// normal character match: both indices advance if current pattern character pair matches
		cache[Pair{i, j}] = match && dfs(i+1, j+1)
		return cache[Pair{i, j}]
	}

	return dfs(0, 0)
}
