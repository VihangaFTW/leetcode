package easy

func IsValidParentheses(s string) bool {

	//? space O(1) mapping does not scale with size of input string
	validMappings := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	//? space O(N) where N is the length of the input string
	closingParentheses := []byte{}

	for _, char := range s {
		//? current character is an opening paranthesis; add its corresponding closing parenthesis to stack
		closingChar, openExists := validMappings[byte(char)]
		if openExists {
			closingParentheses = append(closingParentheses, closingChar)
			continue
		}

		//? current character is a closing parenthesis, check if it matches
		//! no opening parenthesis seen before so there can never be a match
		if len(closingParentheses) == 0 {
			return false
		}
		
		closingChar = closingParentheses[len(closingParentheses)-1]
		closingParentheses = closingParentheses[:len(closingParentheses)-1]

		//* no corresponding opening parenthesis seen before
		if closingChar != byte(char) {
			return false
		}

	}

	return len(closingParentheses) == 0

}
