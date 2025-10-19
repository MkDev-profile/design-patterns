// thread-safe singleton pattern

package main

import "sync"

// type for "singleton-instance"
type Singleton struct {
    data string // some field
}

var (
    instance *Singleton // private var for "singleton-instance"
    once     sync.Once // private "mutex-object"
)

// public method "GetInstance"
func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{
            data: "initial data",
        }
    })

    return instance
}

/*

example use-cases:
- ConfigInstance
- CacheInstance
- LoggerInstance
- DbInstance

p.s. in golang-source-code:
type Once struct {
	done atomic.Bool
	m    Mutex
}

*/













