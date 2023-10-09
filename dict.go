package godict

import "sync"

type Dict[K comparable, V any] struct {
	m *sync.Map
}

// New creates a thread-safe dict mapped by K-key and V-value.
func New[K comparable, V any]() *Dict[K, V] {
	return &Dict[K, V]{
		m: &sync.Map{},
	}
}
