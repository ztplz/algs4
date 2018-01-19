package stack

type Stack interface {
	IsEmpty() bool
	Size() int
	Push(item interface{})
	Pop() interface{}
}

func NewSliceStack() *slicestack {
	return &slicestack{
		items: make([]interface{}, 0),
	}
}

func NewLinkStack() *linkstack {
	return &linkstack{
		len:   0,
		first: nil,
	}
}
