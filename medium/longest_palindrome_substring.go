package medium

// LongestPalindromeDp finds the longest palindromic substring in a given string.
// The algorithm uses a DP approach.
// Time complexity: O(N^2) where N is the length of s
// Space Complexity : O(N^2) where N is the length of s.
func LongestPalindromeDp(s string) string {

	n := len(s)

	// array to store the indices of the current longest palindrome
	longestSubstring := [2]int{0, 0}
	maxLen := 1

	//* memo[i][j] = is s[i...j] a palindrome?
	// initialize n rows
	memo := make([][]bool, n)

	// initialize n columns
	for i := range n {
		memo[i] = make([]bool, n)
	}

	//? each character is by definition the longest substrings of length 1
	for i := range n {
		memo[i][i] = true
	}

	//? for two character substrings; we need to check if they are the same character or not for them to be the longest palindrome of length 2
	for i := range n - 1 {
		memo[i][i+1] = s[i] == s[i+1]
		if memo[i][i+1] && 2 > maxLen {
			maxLen, longestSubstring[0], longestSubstring[1] = 2, i, i+1
		}
	}

	//? fill up the slots for substring length > 2
	for length := 3; length <= n; length++ {
		for start := 0; start+length-1 < n; start++ {
			end := start + length - 1
			memo[start][end] = s[start] == s[end] && memo[start+1][end-1]
			if memo[start][end] && length > maxLen {
				maxLen, longestSubstring[0], longestSubstring[1] = length, start, end
			}
		}
	}

	return s[longestSubstring[0] : longestSubstring[1]+1]
}

func LongestPalindromeOptimal(s string) string {
	start, end := 0, 0

	n := len(s)

	if n < 1 {
		return ""
	}

	// the longest palindrome substring can start from any index so have to check every position
	for i := range n {

		//? go automatically performs integer divison
		//! possible odd length palindrome
		len1 := expandFromIndex(s, i, i)
		//! possible even length palindrome
		len2 := expandFromIndex(s, i, i+1)
		maxLen := maxInt(len1, len2)
		if maxLen > end-start+1 {
			start, end = i-((maxLen-1)/2), i+(maxLen/2)
		}

	}

	return s[start : end+1]
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// expandFromIndex returns the length of the longest palidrome starting at the given pointers
func expandFromIndex(s string, left int, right int) int {

	n := len(s)
	if (right < left) || n < 1 {
		return 0
	}

	for (left <= right) && (left >= 0 && right < n) && (s[left] == s[right]) {
		left--
		right++
	}

	return right - left - 1
}

