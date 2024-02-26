package component

type IDGenerator struct {
	nextID int
}

func (gen *IDGenerator) NextID() int {
	gen.nextID++
	return gen.nextID
}
