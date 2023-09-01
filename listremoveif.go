package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	var prev *NodeL
	for current := l.Head; current != nil; current = current.Next {
		if current.Data == data_ref {
			if prev == nil {
				l.Head = current.Next
			} else {
				prev.Next = current.Next
			}
		} else {
			prev = current
		}
	}
	l.Tail = prev
}
