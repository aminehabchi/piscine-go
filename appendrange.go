package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		ar := []int(nil)
		return ar
	}
	arr := []int{}
	for i := min; i < max; i++ {
		arr = append(arr, i)
	}
	return arr
}
