package medium

//? time complexity: O(3^N * 4^M) where N is the no of digits mapping to three characters and M is the no of digits mapping to four characters
//? aux space: O(X) where X is the length of digits array due to recursion stack (depth of recursion is always N)
func LetterCombinations(digits string) []string {

	if len(digits) == 0 {
		return nil
	}

	//! O(1) space cuz space taken by map does not scale with input size
	phoneMap := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
		'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
	}

	var result []string

	//? depth refers to the current level of the recursion tree
	var backtrack func(depth int, path string)

	backtrack = func(depth int, path string) {

		if depth == len(digits) {
			result = append(result, path)
			return
		}
		//* retrieve the characters that the current digit represents
		letters := phoneMap[digits[depth]]
		//? perform dfs on each branching path 
		for i := range len(letters) {
			backtrack(depth+1, path+string(letters[i]))
		}
	}

	backtrack(0, "")
	return result
}