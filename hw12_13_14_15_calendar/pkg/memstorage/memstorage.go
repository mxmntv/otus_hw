package memstorage

import "sync"

type MemStorage struct {
	Storage map[string]interface{}
	sync.RWMutex
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		Storage: make(map[string]interface{}),
	}
}
