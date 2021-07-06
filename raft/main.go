package main

import "sync"

//1.实现3个选举代码
//2.实现分布式选举代码，加入RPC调用

const raftCount = 3

type Leader struct {
	//任期
	Term int
	//编号
	LeaderId int
}

//声明raft
type Raft struct {
	//锁
	mu sync.Mutex
	//节点标号
	me int
	//当前任期
	currentTerm int
	//为哪个节点投票
	votedFor int
	//状态 0 follower 1 candidate 2 leader
	state int
	//发送最后一条数据的时间
	lastMessageTime int64
	//设置当前的领导
	currentLeader int
	//节点发信息的通道
	message chan bool
	//选举的通道
	electCh chan bool
	//心跳的通道
	heartBeat chan bool
	//返回心跳信号的通道
	heartbeatRe chan bool
	//超时时间
	timeout int
}
