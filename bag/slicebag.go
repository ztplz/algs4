package bag

type slicebag struct {
	items []interface{}
}

func (sb *slicebag) IsEmpty() bool {
	return len(sb.items) == 0
}

func (sb *slicebag) Size() int {
	return len(sb.items)
}

func (sb *slicebag) Add(item interface{}) {
	sb.items = append(sb.items, item)
}
