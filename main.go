package main

import (
	"fmt"
	"sync"
	"time"
)

var lock = &sync.Mutex{}

type single struct{}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}
	return singleInstance
}

func main() {

	for i := 0; i < 10; i++ {
		go getInstance()
	}

	time.Sleep(3 * time.Second)
}
