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

func (m *LinkMap) store(key string) {
	m.mu.Lock()

	m.storage[key] = struct{}{}

	m.mu.Unlock()
}

func (m *LinkMap) storeBatch(keys []string) {
	for i := range keys {
		m.store(keys[i])
	}
}

func (m *LinkMap) flush() {
	m.storage = nil
}
