package repository

import (
	"sync"
)

type repository[T any] struct {
	sync.RWMutex
	m map[string]*T
}

func (r *repository[T]) Save(id string, t *T) error {
	defer r.Unlock()
	r.Lock()
	r.m[id] = t
	return nil
}

func (r *repository[T]) FindOne(id string) (*T, error) {
	defer r.RUnlock()
	r.RLock()
	return r.m[id], nil
}

func (r *repository[T]) FindAll() ([]*T, error) {
	defer r.RUnlock()
	r.RLock()
	slice := make([]*T, 0, len(r.m))
	for _, v := range r.m {
		slice = append(slice, v)
	}
	return slice, nil
}

func NewInMemory[T any]() Repository[T] {
	return &repository[T]{
		m: make(map[string]*T),
	}
}
