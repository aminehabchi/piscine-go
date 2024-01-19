package piscine

func AppendRange(min, max int) []int {
	arr := []int{}
	if min > max {
		return arr
	}
	for i := min; i < max; i++ {
		arr = append(arr, i)
	}
	return arr
}
