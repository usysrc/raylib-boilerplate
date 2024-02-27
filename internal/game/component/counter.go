package component

type Counter struct {
	nextID int
}

func (gen *Counter) NextID() int {
	gen.nextID++
	return gen.nextID
}
