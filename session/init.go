package session

import "errors"

func Init(provider string, addr string, option ...string) (sm SessionMgr, err error) {
	switch provider {
	case "memory":
		sm = NewMemorySessionMgr()
	case "redis":
		sm = NewRedisSessionMgr()
	default:
		err = errors.New("不支持")
	}
	sm.Init(addr, option...)
	return
}
