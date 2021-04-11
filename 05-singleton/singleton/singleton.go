package singleton

import "sync"

var once sync.Once

type Counter struct {
	Count int
}

var instance *Counter

func Instance() *Counter {
	once.Do(func() {
		instance = &Counter{Count: 1}
	})
	return instance
}
