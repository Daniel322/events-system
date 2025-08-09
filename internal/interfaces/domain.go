package interfaces

type Factory[Entity any, CreateData any, UpdateData any] interface {
	Create(data CreateData) (*Entity, error)
	Update(t *Entity, data UpdateData) (*Entity, error)
}
