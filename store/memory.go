package store

import "sync"

type InMemoryStore struct {
	storeMap map[string]*Item
	rwMutex sync.RWMutex
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		storeMap: make(map[string]*Item),
	}
}

func (s *InMemoryStore) Get(key string) (*Item, error) {

	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()

	ret, found := s.storeMap[key]

	if !found {
		ret = nil
	}

	return ret, nil

}

func (s *InMemoryStore) Put(key string, item *Item) (created bool, err error) {

	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	_, exists := s.storeMap[key]

	s.storeMap[key] = item

	return !exists, nil

}

func (s *InMemoryStore) Delete(key string) error {

	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	delete(s.storeMap, key)

	return nil

}
