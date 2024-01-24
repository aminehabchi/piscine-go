package piscine

func Unmatch(a []int) int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	x := 0
	for i := 0; i < len(a); i++ {
		if x == 1 {
			x = 0
			continue
		}
		if i == len(a)-1 {
			return a[len(a)-1]
		}
		if a[i] != a[i+1] {
			return a[i]
		}
		if a[i] == a[i+1] {
			x = 1
			continue
		}
	}
	return -1
}
