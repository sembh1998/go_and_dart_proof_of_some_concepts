package main

import (
	"fmt"

	"sync"

	"github.com/sembh1998/go_and_dart_proof_of_some_concepts/proof_in_go/pattenrs/creational/singleton"
)

var (
	values = make(map[int]*singleton.SingleV1)
	mutex  sync.Mutex
)

func main() {
	fmt.Println("Hello, World!")

	for i := 0; i < 32; i++ {
		go a_ver(i)
	}
	fmt.Scanln()
	see_values()
}

func init() {
	fmt.Println("Main package initialized.")
}

func a_ver(position int) {
	value := singleton.GetInstanceV1()
	value.Value = fmt.Sprintf("position %d", position)

	mutex.Lock()
	values[position] = value
	mutex.Unlock()
}

func see_values() {
	mutex.Lock()
	defer mutex.Unlock()

	for position, value := range values {
		fmt.Printf("Position: %d, Value: %s\n", position, value.Value)
	}
}
