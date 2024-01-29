package piscine

func ListLast(l *List) interface{} {
	var d interface{} = nil
	c := l.Head
	for c != nil {
		d = c.Data
		c = c.Next
	}

	return d
}
