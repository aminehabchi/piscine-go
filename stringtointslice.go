package piscine

func StringToIntSlice(str string) []int {
	s := []int{}

	if str == "" {
		return nil
	}
	for _, char := range str {
		s = append(s, int(char))
	}
	return s
}
