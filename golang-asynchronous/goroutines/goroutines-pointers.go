package main

import (
	"fmt"
	"time"
)

func funPointer(value *string) {
	for {
		fmt.Println(*value)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {

	var test string = "test"
	var pointerTest *string = &test

	go funPointer(pointerTest)

	time.Sleep(time.Second)
}
