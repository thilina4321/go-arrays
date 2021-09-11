package doubles

func DoubleNumber(number *int) int {
	result := *number * 2
	*number = 20
	return result
}
