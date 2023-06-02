package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type SingleV1 struct {
	Value string
}

var singleInstancev1 *SingleV1

func GetInstanceV1() *SingleV1 {
	if singleInstancev1 == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstancev1 == nil {
			fmt.Println("Creating single instance v1 now.")
			singleInstancev1 = &SingleV1{
				Value: "value",
			}
		} else {
			fmt.Println("Single instance v1 already created.")
		}
	} else {
		fmt.Println("Single instance v1 already created.")
	}

	return singleInstancev1
}
