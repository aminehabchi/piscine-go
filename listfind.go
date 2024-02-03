package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	c := l.Head
	for c != nil {
		if comp(c.Data, ref) {
			return &c.Data
		}
		c = c.Next
	}
	return &c.Data
}
