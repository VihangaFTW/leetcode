package main

import (
	"fmt"
	"leetcode/medium"
)

func main() {
	grid := [][]int{{0,0,0}, {0,1,0}, {0,0,0}}
	fmt.Println(medium.UniquePaths2Optimized(grid))

}
