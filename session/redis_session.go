package session

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/garyburd/redigo/redis"
)

type RedisSession struct {
	SessionId  string
	SessionMap map[string]interface{}
	rwlock     sync.RWMutex
	pool       *redis.Pool
	flag       int
}

const (
	//没有变化
	SessionFlagNone = iota
	//有变化
	SessionFlagModify
)

func NewRedisSession(id string, pool *redis.Pool) (session Session) {
	session = &RedisSession{
		SessionId:  id,
		SessionMap: make(map[string]interface{}, 16),
		pool:       pool,
		flag:       SessionFlagNone,
	}
	return
}

func (R *RedisSession) Set(key string, value interface{}) (err error) {
	R.rwlock.Lock()
	defer R.rwlock.Unlock()
	R.SessionMap[key] = value
	R.flag = SessionFlagModify
	return
}

func (R *RedisSession) Get(key string) (v interface{}, err error) {
	v, err = R.getMemoryMap(key)
	if err != nil {
		err = R.getRedisMap()
		if err != nil {
			return
		}
		v, err = R.getMemoryMap(key)
		if err != nil {
			return
		}
	}
	return
}

func (R *RedisSession) getMemoryMap(key string) (v interface{}, err error) {
	R.rwlock.RLock()
	defer R.rwlock.RUnlock()
	v, ok := R.SessionMap[key]
	if !ok {
		err = errors.New("key not exists in session")
		return
	}
	return
}

func (R *RedisSession) getRedisMap() (err error) {
	R.rwlock.RLock()
	defer R.rwlock.RUnlock()
	conn := R.pool.Get()
	gr, err := conn.Do("Get", R.SessionId)
	if err != nil {
		return
	}
	sgr, err := redis.String(gr, err)
	if err != nil {
		return
	}
	fmt.Println("---redis get-----:", sgr)
	err = json.Unmarshal([]byte(sgr), &R.SessionMap)
	if err != nil {
		return
	}

	return
}

func (R *RedisSession) Del(key string) (err error) {
	R.rwlock.Lock()
	defer R.rwlock.Unlock()
	delete(R.SessionMap, key)
	return
}

func (R *RedisSession) Save() (err error) {
	conn := R.pool.Get()
	sm, err := json.Marshal(R.SessionMap)
	if err != nil {
		return err
	}
	if R.flag == SessionFlagModify {
		_, err = conn.Do("Set", R.SessionId, string(sm))
	}
	R.flag = SessionFlagNone
	if err != nil {
		return err
	}
	return
}
