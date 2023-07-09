package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mp := sync.Map{}
	go func() {
		for {
			mp.Store("hello", "world")
		}
	}()
	go func() {
		for {
			t, _ := mp.Load("hello")
			fmt.Println(t)
		}
	}()
	select {}
}

func test() {
	r, w := 100000, 100000
	fmt.Println("读次数:", r, "写次数:", w)
	syncMapTest(r, w)
}

func syncMapTest(readNum, writeNum int) {
	mp := sync.Map{}
	start := time.Now()
	readEnd, writeEnd := make(chan bool), make(chan bool)
	go func() {
		for i := 0; i < writeNum; i++ {
			mp.Store("hello", "world")
		}
		writeEnd <- true
	}()
	go func() {
		for i := 0; i < readNum; i++ {
			_, _ = mp.Load("hello")
		}
		readEnd <- true
	}()
	for i := 0; i < 2; i++ {
		select {
		case <-readEnd:
			fmt.Println("syncMap读已完成，时间花费：", time.Since(start))
		case <-writeEnd:
			fmt.Println("syncMap写已完成，时间花费：", time.Since(start))
		}
	}
}
