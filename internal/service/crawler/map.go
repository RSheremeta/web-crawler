package crawler

import "sync"

type LinkMap struct {
	mu      sync.Mutex
	storage map[string]struct{}
}

func newLinkMap() *LinkMap {
	return &LinkMap{
		mu:      sync.Mutex{},
		storage: map[string]struct{}{},
	}
}

func (m *LinkMap) exists(key string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.storage[key]

	return ok
}

func (m *LinkMap) storeIfNotExists(key string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.storage[key]; ok {
		return false
	}

	m.storage[key] = struct{}{}

	return true
}
