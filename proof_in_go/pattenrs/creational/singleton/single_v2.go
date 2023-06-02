package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once

type SingleV2 struct {
	Value string
}

var singleInstancev2 *SingleV2

func GetInstanceV2() *SingleV2 {
	if singleInstancev2 == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance v2 now.")
				singleInstancev2 = &SingleV2{
					Value: "value",
				}
			})
	} else {
		fmt.Println("Single instance v2 already created.")
	}

	return singleInstancev2
}
