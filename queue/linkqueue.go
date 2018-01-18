package queue

// 队列链表实现

import "errors"

type linkqueue struct {
	len   int
	first *node
	last  *node
}

type node struct {
	item interface{}
	next *node
}

func (lq *linkqueue) IsEmpty() bool {
	return lq.first == nil
}

func (lq *linkqueue) Size() int {
	return lq.len
}

func (lq *linkqueue) Enqueue(item interface{}) {
	oldlast := lq.last
	lq.last = &node{item: item, next: nil}

	if lq.IsEmpty() {
		lq.first = lq.last
	} else {
		oldlast.next = q.last
	}
}

func (lq *linkqueue) Dequeue() (interface{}, error) {
	if lq.IsEmpty() {
		return nil, errors.New("linkqueue is empty")
	}

	item := lq.first.item
	lq.frist = lq.first.next
	lq.len--

	if lq.IsEmpty() {
		lq.last = nil
	}

	return item, nil
}
