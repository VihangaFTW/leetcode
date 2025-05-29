package easy

import "strconv"

func PalindromeNumberStr(x int) bool {
	
	number := strconv.Itoa(x)
	n := len(number)
	left, right := 0, n-1

	for left <= right {
		if left < 0 || right == n || number[left] != number[right] {
			return false
		} 
		left ++ 
		right --
	}

	return true
}


func PalindromeNumber(x int) bool {
	// negative numbers can never be palindrome
	// every number that ends with a 0 (except 0 itself) is never a palindrome
	if x < 0 || (x != 0 && x%10 == 0){
		return false
	}
	//? maintain a reversedHalf container: move digits from x to the reversedHalf counter until we reach the middle
	// terminate when both values are identical
	reversedHalf := 0
	for reversedHalf < x {
		// x%10 returns the last digit
		reversedHalf = (reversedHalf*10) + (x%10)
		// truncate the last digit from x after processing
		x /= 10 
	}
	//? for even sized palindromes: x and half are identical
	//? for odd sized palindromes: x has one more digit than half
	// x/10 truncates the last digit 
	return (x == reversedHalf) || (x == reversedHalf/10) 
}



