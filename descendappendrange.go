package piscine

func DescendAppendRange(max, min int) []int {
	a := []int{}
	if max < min {
		return a
	}
	for i := max; i > min; i-- {
		a = append(a, i)
	}
	return a
}
