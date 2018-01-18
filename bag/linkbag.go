package bag

import (
	"fmt"
)

type linkbag struct {
	len   int
	first *node
}

type node struct {
	item interface{}
	next *node
}

func newNode(item interface{}, next *node) *node {
	return &node{item: item, next: next}
}

func (lb *linkbag) IsEmpty() bool {
	return lb.first == nil
}

func (lb *linkbag) Size() int {
	return lb.len
}

func (lb *linkbag) Add(item interface{}) {
	// lb.first = &node{item: item, next: lb.first}
	lb.first = newNode(item, lb.first)
	lb.len++
	fmt.Println(lb.len)
}
