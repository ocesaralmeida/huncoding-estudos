package main

import (
	"fmt"
)

func main() {

	var test string = "testeDefault"

	switch test {

	case "test", "test2", "test999":
		fmt.Println("CAIU NA PRIMEIRA CONDIÇÃO")

	case "test3":
		fmt.Println("CAIU NA SEGUNDA POSIÇÃO")
	default:
		fmt.Println("CAIU NO DEFAULT")
	}
}
