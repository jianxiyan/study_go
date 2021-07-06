package main

import (
	"fmt"
	"sync"
	"time"
)

var aOne sync.Once
var bOne sync.Once

// var wy sync.WaitGroup

func write(a chan int) {
	// defer wy.Done()
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		a <- i
	}
	aOne.Do(func() {
		fmt.Println("关闭a")
		close(a)
	})

}

func read(a, b chan int) {
	// defer wy.Done()
	for {
		x, ok := <-a
		if !ok {
			break
		}
		b <- x * x
	}
	bOne.Do(func() {
		fmt.Println("关闭b")
		close(b)
	})
}

func main() {
	a := make(chan int, 10)
	b := make(chan int, 10)
	// wy.Add(2)
	go write(a)
	go write(a)
	go read(a, b)
	go read(a, b)
	fmt.Println("---------------添加结束-----------------")
	// wy.Wait()
	for ret := range b {
		fmt.Println(ret)
	}

}
