package main

import "fmt"

func main() {

	var testValue string = "César"
	copyStringValue(testValue)
	fmt.Println("Fora da função copyStringValue: ", testValue)

	originalStringValue(&testValue)
	fmt.Println("Fora da função originalStringValue: ", testValue)
}

func copyStringValue(stringValue string) {
	stringValue = "TEST"
	fmt.Println("Dentro da função copyStringValue: ", stringValue)
}

func originalStringValue(stringValue *string) {
	*stringValue = "TEST"
	fmt.Println("Dentro da função originalStringValue: ", *stringValue)
}
