package reqvars

import (
	"net/http"
	"sync"
)

type Storage struct {
	lock sync.RWMutex
	vars map[*http.Request]RequestStorage
}

type RequestStorage struct {
	data map[string]interface{}
}

func New() *Storage {
	return &Storage{
		vars: map[*http.Request]RequestStorage{},
	}
}

func (s *Storage) Open(r *http.Request) RequestStorage {
	s.lock.Lock()
	s.vars[r] = RequestStorage{
		data: make(map[string]interface{}),
	}
	rs := s.vars[r]
	s.lock.Unlock()
	return rs
}

func (s *Storage) Close(r *http.Request) {
	s.lock.Lock()
	delete(s.vars, r)
	s.lock.Unlock()
}

func (s RequestStorage) Set(key string, value interface{}) {
	s.data[key] = value
}

func (s RequestStorage) Get(key string) (value interface{}) {
	value = s.data[key]
	return
}

func (s RequestStorage) TryGet(key string) (value interface{}, ok bool) {
	value, ok = s.data[key]
	return
}
