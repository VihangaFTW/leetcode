package medium

import (
	"math"
)

func ReverseInteger(x int32) int32 {

	const minVal int32 = math.MinInt32
	const maxVal int32 = math.MaxInt32

	const maxLastDigit int32 = maxVal % 10
	const minLastDigit int32 = minVal % 10

	var reversed int32 = 0

	for x != 0 {
		last_digit := int32(x % 10)
		//? check for overflow
		if (reversed > maxVal/10) || (reversed == maxVal/10 && last_digit == maxLastDigit) {
			return 0
		} else if (reversed < minVal/10) || (reversed == minVal/10 && last_digit == minLastDigit) {
			return 0
		}
		// add the new digit to number
		reversed = (reversed * 10) + last_digit
		// truncate the last digit from the number
		x = x/10
	}

	return reversed

}
