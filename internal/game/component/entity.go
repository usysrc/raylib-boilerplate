package component

type Entity int

func NewEntity(id int) Entity {
	e := Entity(id)
	return e
}
