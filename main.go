package main

import (
	"fmt"
	"leetcode/hard"
)

func main() {
	grid := [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 2}}
	fmt.Println(hard.UniquePaths3(grid))

}
