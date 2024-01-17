package piscine

func Join(strs []string, sep string) string {
	b := ""
	for i := 0; i < len(strs); i++ {
		b = b + strs[i]
		if i != len(strs)-1 {
			b = b + sep
		}
	}
	return b
}
