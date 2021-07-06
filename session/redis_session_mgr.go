package session

import (
	"fmt"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
)

type RedisSessionMgr struct {
	Addr       string
	Passwd     string
	Pool       *redis.Pool
	rwlock     sync.RWMutex
	SessionMap map[string]Session
}

func NewRedisSessionMgr() SessionMgr {
	sr := &RedisSessionMgr{
		SessionMap: make(map[string]Session, 32),
	}
	return sr
}

func (R *RedisSessionMgr) Init(addr string, option ...string) (err error) {
	if len(option) > 0 {
		R.Passwd = option[0]
	}
	R.Addr = addr
	R.Pool = R.buildRedisPool()
	return
}

func (R *RedisSessionMgr) buildRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   16,   //最初的连接数量
		MaxActive: 1000, //最大连接数量
		// MaxActive:0,    //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			conn, err := redis.Dial("tcp", R.Addr)
			if err != nil {
				return nil, err
			}
			//若有密码
			if _, err = conn.Do("AUTH", R.Passwd); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, nil
		},
		//测试redis链接是否通
		//线上要关掉
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}
}

func (R *RedisSessionMgr) CreateSession() (session Session, err error) {
	R.rwlock.Lock()
	defer R.rwlock.Unlock()
	id := uuid.NewV4()
	sessionId := id.String()
	fmt.Printf("uuid : %s\n", sessionId)
	session = NewRedisSession(sessionId, R.Pool)
	R.SessionMap[sessionId] = session
	return
}

func (R *RedisSessionMgr) Get(sessionId string) (session Session, err error) {
	R.rwlock.RLock()
	defer R.rwlock.RUnlock()
	session = R.SessionMap[sessionId]
	return
}

func (R *RedisSessionMgr) GetAll() (sm map[string]Session, err error) {
	R.rwlock.RLock()
	defer R.rwlock.RUnlock()
	sm = R.SessionMap
	return
}
