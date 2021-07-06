package session

import (
	"fmt"
	"testing"
)

func TestInitMemory(t *testing.T) {
	sm, err := Init("redis", "127.0.0.1:6379")
	if err != nil {
		fmt.Printf("init err: %s\n", err)
	}
	ms, err := sm.CreateSession()

	if err != nil {
		fmt.Printf("CreateSession err: %s\n", err)
	}
	ms.Set("appkey", "111")
	err = ms.Save()
	if err != nil {
		fmt.Printf("Save err: %s\n", err)
	}
	getAll, _ := sm.GetAll()
	for k, v := range getAll {
		sv, _ := v.Get("appkey")
		fmt.Printf("k: %s, v: %v\n", k, sv)
	}

}
