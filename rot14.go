package piscine

func Rot14(s string) string {
	ss := []byte(s)
	b := ""
	a := 0
	for i := 0; i < len(ss); i++ {
		if ss[i] >= 65 && ss[i] <= 90 {

			if ss[i] >= 65 && ss[i] <= 76 {
				ss[i] = ss[i] + 14
			} else {
				a = int(ss[i]-77) + 65
				ss[i] = byte(a)
			}

		} else if ss[i] >= 97 && ss[i] <= 122 {

			if ss[i] >= 97 && ss[i] <= 108 {
				ss[i] = ss[i] + 14
			} else {
				a = int(ss[i]-109) + 97
				ss[i] = byte(a)
			}

		}
		b = b + string(ss[i])
	}
	return b
}
