package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return []int(nil)
	}
	arr := make([]int, max-min)
	arr[0] = min
	for i := 1; i < max-min; i++ {
		arr[i] += 1
	}
	return arr
}
