package repository

type Repository[T any] interface {
	Save(id string, t *T) error
	FindOne(id string) (*T, error)
	FindAll() ([]*T, error)
}
