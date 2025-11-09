package main

import "fmt"

func main() {

	var user user = user{
		name: "Teste",
		age:  20,
	}

	fmt.Println(user)
	fmt.Println(user.name)
	fmt.Println(user.age)

}

type user struct {
	name string
	age  int
}
