package piscine

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}
	for l.Tail != nil {
		l.Head = n
		l.Tail = n
	}
	n.Next = l.Head
	l.Head = n
}
