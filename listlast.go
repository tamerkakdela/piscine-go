package piscine

func ListLast(l *List) interface{} {
	for l.Tail != nil {
		return l.Tail.Data
	}
	return nil
}
