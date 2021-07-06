package session

import (
	"errors"
	"sync"
)

type MemorySession struct {
	SessionId string
	Data      map[string]interface{}
	rwlock    sync.RWMutex
}

func NewMemorySession(id string) Session {
	m := &MemorySession{
		SessionId: id,
		Data:      make(map[string]interface{}, 16),
	}

	return m
}

func (M *MemorySession) Set(key string, value interface{}) (err error) {
	M.rwlock.Lock()
	defer M.rwlock.Unlock()
	M.Data[key] = value
	return
}

func (M *MemorySession) Get(key string) (v interface{}, err error) {
	M.rwlock.RLock()
	defer M.rwlock.RUnlock()
	v, ok := M.Data[key]
	if !ok {
		err = errors.New("key not exists in session")
		return
	}

	return
}

func (M *MemorySession) Del(key string) (err error) {
	M.rwlock.Lock()
	defer M.rwlock.Unlock()
	delete(M.Data, key)
	return
}

func (M *MemorySession) Save() (err error) {
	return
}
