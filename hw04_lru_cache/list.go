package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	length int
	first  *ListItem
	last   *ListItem
}

func (l *list) zeroList(v interface{}) *ListItem {
	candidate := ListItem{
		Value: v,
		Next:  nil,
		Prev:  nil,
	}
	l.first = &candidate
	l.last = &candidate
	l.length++
	return &candidate
}

func (l list) Len() int {
	return l.length
}

func (l list) Front() *ListItem {
	return l.first
}

func (l list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	if l.length > 0 {
		candidate := ListItem{
			Value: v,
			Next:  l.first,
			Prev:  nil,
		}
		l.first.Prev = &candidate
		l.first = &candidate
		l.length++
		return &candidate
	}
	return l.zeroList(v)
}

func (l *list) PushBack(v interface{}) *ListItem {
	if l.length > 0 {
		candidate := ListItem{
			Value: v,
			Next:  nil,
			Prev:  l.last,
		}
		l.last.Next = &candidate
		l.last = &candidate
		l.length++
		return &candidate
	}
	return l.zeroList(v)
}

func (l *list) Remove(i *ListItem) {
	if l.length > 1 {
		nestedBlocks(i, l)
	} else if l.length == 1 {
		l.length = 0
		l.first = nil
		l.last = nil
	}
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev != nil {
		if i.Next != nil {
			i.Prev.Next = i.Next
			i.Next.Prev = i.Prev
		} else {
			i.Prev.Next = nil
		}
		i.Next = l.first
		i.Prev = nil
		l.first = i
	} else {
		return
	}
}

func nestedBlocks(i *ListItem, l *list) {
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		i.Prev.Next = nil
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		i.Next.Prev = nil
	}
	l.length--
}

func NewList() List {
	return new(list)
}
