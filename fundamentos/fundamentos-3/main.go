package main

import "fmt"

func main() {

	var test []string = []string{"test0", "test1", "test2"}

	for i, value := range test {
		fmt.Println(i, value)
	}
}
