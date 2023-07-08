package main

import (
	"flag"
	"sync"
	"the-lock/atom"
)

// Locker ...
type Locker interface {
	Lock()
	Unlock()
}

// Waiter ...
type Waiter interface {
	Add(n int)
	Wait()
}

func main() {
	var p int
	flag.IntVar(&p, "pause", 0, "pause loop times")
	var c int
	flag.IntVar(&c, "concurrency", 2, "goroutine count")
	var t int
	flag.IntVar(&t, "times", 100000000, "for loop times")
	flag.Parse()
	var l Locker = &atom.Spin{0, int32(p)}
	var w Waiter = new(sync.WaitGroup)
	var n int
	for i := 0; i < c; i++ { // 两个协程
		w.Add(1)
		go routine(&n, l, w, t, 1-i%2*2)
	}
	w.Wait()
	println(n)
}

func routine(v *int, l Locker, w Waiter, c, d int) {
	defer w.Add(-1)
	for t := 0; t < c; t++ { // 循环1亿次
		func() {
			l.Lock()
			defer l.Unlock()
			*v += d
		}()
	}
}
