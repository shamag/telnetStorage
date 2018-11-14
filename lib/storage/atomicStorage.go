package storage

import (
	"strings"
	"sync"
	"sync/atomic"
)

type AtomicStorage struct {
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

func (s *AtomicStorage) Delete(key string) bool {
	s.Lock()
	defer s.Unlock()
	store := s.data.Load().(Storage)
	if _, exist := store[key]; !exist {
		return false
	}

	s2 := make(Storage)
	for k, v := range store {
		if strings.Compare(k, key) != 0 {
			s2[k] = v
		}
	}
	s.data.Store(s2)
	return true
}

func (s *AtomicStorage) Get(key string) (string, bool) {
	val, exist := s.data.Load().(Storage)[key]
	return val, exist
}

func (s *AtomicStorage) Exist(key string) bool {
	_, exist := s.data.Load().(Storage)[key]
	return exist
}
