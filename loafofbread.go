package piscine

func LoafOfBread(str string) string {
	var f []rune
	s := str
	if s == "" {
		f = append(f, '\n')
		return string(f)
	}

	if len(s) < 5 {
		f = append(f, []rune("Invalid Output\n")...)
		return string(f)
	}

	j := 0
	for i := 0; i < len(s); i++ {
		if j < 5 && rune(s[i]) == ' ' {
			continue
		}
		if j == 5 {
			if i != len(s)-1 && s[i+1] == ' ' {
				continue
			}
			if i == len(s)-1 {
				break
			}
			f = append(f, ' ')
			j = 0
			continue
		}
		f = append(f, rune(s[i]))
		j++
	}

	f = append(f, '\n')
	return string(f)
}
