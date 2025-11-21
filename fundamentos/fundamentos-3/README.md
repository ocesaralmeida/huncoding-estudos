# Fundamentos Golang 3 - Esttutura de repetição e funções

## for
Tipo `interface{}` aceita qualquer tipo, é como o `any` do javascript 

```go
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
```

---

## while
Não existe a estrutura de repetição `while` em golang, o que existe é uma adaptação do `for` 

```go
package main

import "fmt"

func main() {
	var test int = 0

	for test <= 10 {
		fmt.Println("VALOR: ", test)
		test++
	}
}
```

---

## do  while
Executa uma vez para depois verificar se a condição é verdadeira
```go
package main

import "fmt"

func main() {
	var anyExpression bool = false

	for ok := true; ok; ok = anyExpression {
		fmt.Println("PASSOU AQUI")
	}
}

```

---

## for com range
```go
package main

import "fmt"

func main() {

	var test []string = []string{"test0", "test1", "test2"}

	for i, value := range test {
		fmt.Println(i, value)
	}
}

```

---

## funções com retorno nomeado e retorno múltiplo
Retorno múltiplo:

```go
package main

import "fmt"

func main() {

	value, err := test()

	if err != nil {
		fmt.Println("ERRO")
	} else {
		fmt.Print("TEST", value)
	}

}

func test() (string, error) {
	return "", nil
}
```

Retorno nomeado:

```go
package main

import "fmt"

func main() {

	value, _ := test()

	fmt.Print(value)

}

func test() (retornoString string, retornoErro error) {
	retornoString = "test"
	retornoErro = nil

	return
}
```

---

## passando funções por parameetro e recebendo funções como retorno

```go
package main

import "fmt"

func main() {

	funcaoTest := func(test string, testInt int) {

		fmt.Println(test, testInt)
	}

	test(funcaoTest)
}

func test(value func(string, int)) {

	value("otavio", 20)
}
```
--- 

## recebendo uma função como parâmetro

```go
package main

import "fmt"

func main() {

	funcao := test()

	funcao("otavio", 20)
}

func test() func(string, int) {

	funcaoTest := func(valorString string, valorInt int) {
		fmt.Println(valorString, valorInt)
	}

	return funcaoTest
}
```

---

## função anônima

```go
package main

import "fmt"

func main() {

	test()
}

func test() {

	func(valorString string, valorInt int) {
		fmt.Println(valorString, valorInt)
	}("otavio", 20)

}
```

---

## recebendo N parâmetros em uma função

```go
package main

import "fmt"

func main() {

	test()
}

func test() {

	func(valorString string, valorInt int) {
		fmt.Println(valorString, valorInt)
	}("otavio", 20)

}
```