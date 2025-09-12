package medium

func HouseRobber(nums []int) int {
	// max loot from two previous houses
	two_houses_back, one_house_back := 0, 0

	// for each house, decide whether to rob it or not
	for _, n := range nums {
		// calculate total loot by robbing current house (meaning we cannot rob prev house)
		// or dont rob current house (meaning we can rob prev house)
		// take max to get best option of the two options
		newSum := maxNum(n + two_houses_back, one_house_back)
		two_houses_back = one_house_back
		one_house_back = newSum
	}

	return one_house_back
}


func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
