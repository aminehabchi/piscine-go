package piscine

func ConcatParams(args []string) string {
	str := ""
	for i := 0; i < len(args); i++ {
		str = str + args[i]
		if i != len(args)-1 {
			str = str + "\n"
		}
	}
	return str
}
