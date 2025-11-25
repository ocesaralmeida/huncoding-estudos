# Ponteiros

- go trabalha com o conceito de ponteiros, ou seja, referencia de memória
- isso é interessante pois em determinados momentos, eu posso querer acessar o valore real da variável, em outros, posso querer apenas a "cópia" desse valor. Ou seja, posso modificar esse valor sem afetar a referência

```go
package main

import "fmt"

// trabalhando com ponteiros
func main() {
	var y int = 10
	var x *int = &y

	fmt.Println("endereço de y, armazenado em x: ", x) //0xc00000a0f8
	fmt.Println("valor de y, olhando para x: ", *x)    // 10
	fmt.Println("endereço de x: ", &x)                 //0xc000060060

	y = 15

	fmt.Println("novo valor de y, olhando para y: ", y)
	fmt.Println("novo valor de y, olhando para x: ", *x)
	fmt.Println("endereço de y, armazenado em x: ", x) //0xc00000a0f8
	fmt.Println("endereço de x: ", &x)                 //0xc000060060
}
```

```go
package main

import "fmt"

// trabalhando sem ponteiros
func main() {
	var y int = 10
	var x int = y

	fmt.Println("valor de y: ", y) // 10
	fmt.Println("valor de x: ", x) // 10

	y = 15

	fmt.Println("novo valor de y: ", y) // 15
	fmt.Println("valor de x: ", x)      // 10
}
```

```go
package main

import "fmt"

// exemplo com funções
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
```