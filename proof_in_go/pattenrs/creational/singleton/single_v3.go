package singleton

import (
	"fmt"
)

type SingleV3 struct {
	Value string
}

var singleInstancev3 *SingleV3

func GetInstanceV3() *SingleV3 {
	if singleInstancev3 == nil {
		fmt.Println("Creating single instance v3 now.")
		singleInstancev3 = &SingleV3{
			Value: "value",
		}
	} else {
		fmt.Println("Single instance v3 already created.")
	}

	return singleInstancev3
}
