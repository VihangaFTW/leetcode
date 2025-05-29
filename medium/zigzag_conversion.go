package medium

import "strings"

func ZigzagConversion(s string, numRows int) string {
	n := len(s)

	if numRows == 1 {
		return s
	}
	// initialize stringBuilder instances in memory: each row gets its own stringBuilder
	sbs := make([]strings.Builder, numRows)

	//? row depends on dir which is is either -1 or 1: 1 indicates go down a row; -1 indicates go up a row
	row, dir := 0, 0

	// loop through each character in given string and add each character to its respective stringBuilder
	for i := range n {
		char := s[i]
		//? evaluate row that needs to be extended (initially dir = 0 so we extend row 0)
		row += dir
		// append character to its string builder
		sbs[row].WriteByte(char)

		// since we increment row inside the loop first, row needs to adjusted for next iteration when at the edge boundaries to flip direction of movement to maintain the zigzag pattern
		if row == 0 || row == numRows-1 {
			//? currently at top row so go down next iteration
			if row == 0 {
				dir = 1

			} else { //? currently at bottom row so go up next iteration
				dir = -1
			}
		}

	}
	// combine the strings from all rows to build the result string
	var result strings.Builder
	for i := range numRows {
		result.WriteString(sbs[i].String())
	}

	return result.String()

}
