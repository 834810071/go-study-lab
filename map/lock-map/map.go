package main

import (
	"fmt"
	"sync"
)

type MyMap struct {
	sync.RWMutex
	mp map[string]string
}

func (m *MyMap) Get(key string) (value string, ok bool) {
	m.RLock()
	defer m.RUnlock()
	value, ok = m.mp[key]
	return
}

func (m *MyMap) Set(key, value string) {
	m.Lock()
	defer m.Unlock()
	m.mp[key] = value
}

func main() {
	mp := MyMap{
		RWMutex: sync.RWMutex{},
		mp:      make(map[string]string),
	}
	go func() {
		for {
			mp.Set("hello", "world")
		}
	}()
	go func() {
		for {
			t, _ := mp.Get("hello")
			fmt.Println(t)
		}
	}()
	select {}
}
