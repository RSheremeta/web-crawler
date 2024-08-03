package crawler

import "sync"

type LinkMap struct {
	storage sync.Map
}

func newLinkMap() *LinkMap {
	return &LinkMap{}
}

func (m *LinkMap) storeIfNotExists(key string) bool {
	_, loaded := m.storage.LoadOrStore(key, struct{}{})
	return !loaded
}

func (m *LinkMap) len() int {
	count := 0
	m.storage.Range(func(_, _ interface{}) bool {
		count++
		return true
	})
	return count
}
