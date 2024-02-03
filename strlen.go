package piscine

func StrLen(s string) int {
	b := []byte(s)
	var a int
	var c int = 0
	for i := 0; i < len(b); i++ {
		if b[i] <= 255 && b[i] >= 128 {
			if c == 1 {
				a++
				c = 0
			} else if c == 0 {
				c++
			}
		}

		if b[i] < 128 {
			a++
		}
	}
	return a
}
