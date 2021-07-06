package session

import (
	"fmt"
	"sync"

	uuid "github.com/satori/go.uuid"
)

type MemorySessionMgr struct {
	SessionMap map[string]Session
	rwlock     sync.RWMutex
}

func NewMemorySessionMgr() *MemorySessionMgr {
	m := &MemorySessionMgr{
		SessionMap: make(map[string]Session, 1024),
	}
	return m
}

func (M *MemorySessionMgr) Init(addr string, option ...string) (err error) {
	return
}

func (M *MemorySessionMgr) CreateSession() (session Session, err error) {
	M.rwlock.Lock()
	defer M.rwlock.Unlock()
	id := uuid.NewV4()
	sessionId := id.String()
	fmt.Printf("uuid : %s\n", sessionId)
	session = NewMemorySession(sessionId)
	M.SessionMap[sessionId] = session
	return
}

func (M *MemorySessionMgr) Get(sessionId string) (session Session, err error) {
	M.rwlock.RLock()
	defer M.rwlock.RUnlock()
	session = M.SessionMap[sessionId]
	return
}

func (M *MemorySessionMgr) GetAll() (sm map[string]Session, err error) {
	M.rwlock.RLock()
	defer M.rwlock.RUnlock()
	sm = M.SessionMap
	return
}
