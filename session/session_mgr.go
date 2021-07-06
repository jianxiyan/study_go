package session

//session的管理者
type SessionMgr interface {
	//初始化
	Init(addr string, option ...string) (err error)
	CreateSession() (session Session, err error)
	Get(sessionId string) (session Session, err error)
	GetAll() (map[string]Session, error)
}
