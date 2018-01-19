package stack

import "errors"

type linkstack struct {
	len   int
	first *node
}

type node struct {
	item interface{}
	next *node
}

func (ls *linkstack) IsEmpty() bool {
	return ls.first == nil
}

func (ls *linkstack) Size() int {
	return ls.len
}

func (ls *linkstack) Push(item interface{}) {
	oldfirst := ls.first
	ls.first = &node{item: item, next: oldfirst}
	ls.len++
}

func (ls *linkstack) Pop() (interface{}, error) {
	if ls.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	item := ls.first.item
	ls.first = ls.first.next
	ls.len--

	return item, nil
}
