package easy

import (
	"strings"
)

// ? time: O(N *M) where N is the length of the list and M is the size of the smallest string. Spaace: O(M)	for building string
func LongestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	prefix := new(strings.Builder)  //? space: O(M)

	minLen := len(strs[0])

	for i := 1; i < len(strs); i++ {  //? time: O(N)
		minLen = min(minLen, len(strs[i]))
	}

	for i := range minLen {  //? time: O(M * N)
		char := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != char {
				return prefix.String()
			}
		}
		prefix.WriteByte(char)
	}

	return prefix.String()
}
