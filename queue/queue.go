package queue

// 队列
type Queue interface {
	enqueue(item interface{})
	dequeue() interface{}
	IsEmpty() bool
	Size() int
}

func NewSliceQueue() *slicequeue {
	return &slicequq{
		items: make([]interface{}, 0)
	}
}

func NewLinkQueue() *linkqueue {
	return &linkqueue{
		len: 0,
		first: nil,
		last: nil
	}
}