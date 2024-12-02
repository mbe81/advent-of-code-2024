package calc

func AbsDifference(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func Sign(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}
