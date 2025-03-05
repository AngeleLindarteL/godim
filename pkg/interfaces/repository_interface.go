package interfaces

type Repository[T any] interface {
	FindOne(id string) (*T, error)
	FindAll() ([]*T, error)
	FindMany(filter Filter) ([]*T, error)
	CreateOne(entity *T) error
	UpdateOne(entity *T) error
	Delete(filter Filter) error
}
