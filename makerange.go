package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return []int(nil)
	}

	arr := make([]int, max-min)

	count := 0

	for i := min; i < max; i++ {
		arr[count] = i
		count++
	}

	return arr
}
