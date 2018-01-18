import "errors"

// 队列切片实现
type slicequeue struct {
	items []interface{}
}

func (sq *slicequeue) IsEmpty() bool {
	return len(sq.items) == 0
}

func (sq *slicequeue) Size() int {
	return len(sq.items)
}

func (sq *slicequeue) Enqueue(item interface{}) {
	sq.items = append(sq.items, item)
}

func (sq *slicequeue) Dequeue() (interface{}, error) {
	if sq.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}

	item = sq.items[0]
	sq.items = sq.items[1:]

	return item, nil
}