package piscine

func Split(s, sep string) []string {
	b := ""
	p := ""
	words := []string{}
	x := 0
	for i := 0; i < len(s); i++ {
		if x > 0 {
			x--
			continue
		}
		p = ""
		if i < len(s)-len(sep)+1 {
			for j := 0; j < len(sep); j++ {
				p = p + string(s[i+j])
			}
		}

		if p == sep && i == 0 {
			p = ""
			x = len(sep) - 1
			continue
		}

		if p != sep {
			b = b + string(s[i])
		}
		if p == sep || i == len(s)-1 {

			words = append(words, b)
			b = ""
			x = len(sep) - 1
		}

	}
	return words
}
