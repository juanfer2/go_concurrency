package safemap

import (
	"fmt"
	"sync"
)

type SafeMap[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

func New[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		data: make(map[K]V),
	}
}

func (s *SafeMap[K, V]) Add(key K, value V) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
}

func (s *SafeMap[K, V]) Get(key K) (V, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	value, ok := s.data[key]
	if !ok {
		return value, fmt.Errorf("key %v not found", key)
	}

	return value, nil
}

func (s *SafeMap[K, V]) Update(key K, value V) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.data[key]
	if !ok {
		return fmt.Errorf("key %v not found", key)
	}

	s.data[key] = value

	return nil
}

func (s *SafeMap[K, V]) Delete(key K) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.data[key]
	if !ok {
		return fmt.Errorf("key %v not found", key)
	}

	delete(s.data, key)

	return nil
}

func (s *SafeMap[K, V]) HasKey(key K) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.data[key]

	return ok
}
