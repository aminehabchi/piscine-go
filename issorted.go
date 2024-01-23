package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if a[0] > a[len(a)-1] {
		for i := 0; i < len(a)-1; i++ {
			if f(a[i], a[i+1]) < 0 {
				return false
			}
		}
	} else {
		for i := 0; i < len(a)-1; i++ {
			if f(a[i], a[i+1]) > 0 {
				return false
			}
		}
	}
	return true
}

func f(x, y int) int {
	if x > y {
		return 1
	}
	if x == y {
		return 0
	}

	return -1
}
