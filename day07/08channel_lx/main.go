package main

import (
	"fmt"
	"math/rand"
	"time"
)

type job struct {
	value int64
}

type sumJob struct {
	job *job
	sum int64
}

func write(w chan<- *job) {
	for {
		x := rand.Int63()
		time.Sleep(100 * time.Millisecond)
		j := &job{
			value: x,
		}
		w <- j
	}
}

func read(w <-chan *job, r chan<- *sumJob) {
	for {
		j := <-w
		var sum int64
		v := j.value
		for v > 0 {
			sum += v % 10
			v = v / 10
		}
		r <- &sumJob{
			job: j,
			sum: sum,
		}
	}
}

func main() {
	w := make(chan *job, 10)
	r := make(chan *sumJob, 10)
	go write(w)

	for i := 0; i < 3; i++ {
		go read(w, r)
	}

	for sj := range r {
		fmt.Println(sj.job.value, sj.sum)
	}

}
