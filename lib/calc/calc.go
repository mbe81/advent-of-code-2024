package calc

func AbsDifference(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
