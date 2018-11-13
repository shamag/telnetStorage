package storage

import (
	"sync"
	"sync/atomic"
)

type Storage map[string]string

type SyncStorage struct {
	data Storage
	sync.RWMutex
}

type AtomicStorage struct{
	data atomic.Value
	sync.Mutex
}

func (s *AtomicStorage) Create() *AtomicStorage {
	d := make(Storage)
	var s2 atomic.Value
	s2.Store(d)
	return &AtomicStorage{
		data: s2,
	}
}

func (s *AtomicStorage) Set(key string, value string) {
	s.Lock()

	defer s.Unlock()
	store := s.data.Load().(Storage)
	s2 := make(Storage)
	for k, v := range store {
		s2[k] = v
	}
	s2[key] = value
	s.data.Store(s2)

}

func (s *AtomicStorage) Get(key string) string {
	return s.data.Load().(Storage)[key]
}

func (s *SyncStorage) Create() *SyncStorage {
	store := make(Storage)
	return &SyncStorage{
		data: store,
	}
}

func (s *SyncStorage) Set(key string, value string) {
	s.Lock()
	s.data[key] = value
	s.Unlock()
}

func (s *SyncStorage) Get(key string) string {
	s.RLock()
	defer s.RUnlock()
	return s.data[key]
}
