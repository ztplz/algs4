package stack

import "errors"

type slicestack struct {
	items []interface{}
}

func (ss *slicestack) IsEmpty() bool {
	return len(ss.items) == 0
}

func (ss *slicestack) Size() int {
	return len(ss.items)
}

func (ss *slicestack) Push(item interface{}) {
	ss.items = append(ss.items, item)
}

func (ss *slicestack) Pop() (interface{}, error) {
	if ss.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	item := ss.items[len(ss.items)-1]
	ss.items[len(ss.items)-1] = nil
	ss.items = ss.items[:len(ss.items)-1]

	return item, nil
}
