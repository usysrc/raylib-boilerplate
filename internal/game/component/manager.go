package component

type Manager struct {
	idCounter  *Counter
	Positions  map[Entity]*Position
	Velocities map[Entity]*Velocity
	Renders    map[Entity]*Render
	Tags       map[Entity]*Tag
}

func NewManager() *Manager {
	return &Manager{
		idCounter:  &Counter{},
		Positions:  make(map[Entity]*Position),
		Velocities: make(map[Entity]*Velocity),
		Renders:    make(map[Entity]*Render),
		Tags:       make(map[Entity]*Tag),
	}
}

func (m *Manager) NewEntity() Entity {
	return NewEntity(m.idCounter.NextID())
}
