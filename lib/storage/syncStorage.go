package storage

import (
	"sync"
)

type SyncStorage struct {
	data Storage
	sync.RWMutex
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

func (s *SyncStorage) Delete(key string) bool {
	s.Lock()
	defer s.Unlock()
	if _, exist := s.data[key]; !exist {
		return false
	}
	delete(s.data, key)
	return true
}

func (s *SyncStorage) Get(key string) (string, bool) {
	s.RLock()
	defer s.RUnlock()
	val, exist := s.data[key]
	return val, exist
}

func (s *SyncStorage) Exist(key string) bool {
	s.RLock()
	defer s.RUnlock()
	_, exist := s.data[key]
	return exist
}
