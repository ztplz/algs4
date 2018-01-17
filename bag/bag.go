package bag

type Bag interface {
	IsEmpty() bool
	Size() int
	Add(item interface{})
}

func NewSliceBag() Bag {
	return &slicebag{
		items: make([]interface{}, 0),
	}
}
