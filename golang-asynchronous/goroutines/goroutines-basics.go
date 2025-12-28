package main

import (
	"fmt"
	"time"
)

func fun(value string) {
	for i := 0; i < 3; i++ {
		fmt.Println(value)
		time.Sleep(1 * time.Second)
	}
}

func test() {
	// Direct call
	fun("Direct call")

	// TODO: write goroutines with differents variants for function call
	fgx := fun
	go fgx("goroutine - 2")

	// TODO: goroutine function call
	go fun("goroutine - 1")

	// TODO: goroutine with anonymous value call
	go func() {
		fun("goroutine - 3")
	}()

	// TODO: wait for goroutine to end
	time.Sleep(5 * time.Second)

	fmt.Println("Done...")
}
