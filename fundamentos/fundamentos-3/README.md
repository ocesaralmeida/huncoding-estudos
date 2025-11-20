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
`Array` tem tamanho fixo

```go
func main() {
    var testArray [4]string = [4]string{"teste1", "teste2", "teste3", "teste4"} 
}

```

`Slice` tem tamanho variável, pode aumentar de acordo com a necesidade

```go
func main() {
    var testSlice []string = []string{"teste", "teste2", "teste3", "teste4"}
}
```

Comparativo dos dois:

```go
func main() {
	var testArray [4]string = [4]string{"teste", "teste", "teste", "teste"}
	fmt.Println("testArray: ", testArray)
	fmt.Println(cap(testArray))
	fmt.Println(len(testArray))

	//testArray = append(testArray, "teste") // não deixa essa atribição acontecer, pois array tem tamanho definido

	var testSlice []string = []string{"teste", "teste", "teste", "teste"}
	fmt.Println("testSlice: ", testSlice)
	fmt.Println(cap(testSlice))
	fmt.Println(len(testSlice))

	testSlice = append(testSlice, "teste") //deixa a tribuição acontecer, pois slice não tem tamanho definido, aumenta a capacidade de acordo com a necessidade
	fmt.Println(cap(testSlice))
	fmt.Println(len(testSlice))

}
```

---

## structs

Semelhante a classe em outras linguagens, são estruturas de dados para agrupar regras de negócios

```go
func main() {

	var user user = user{
		name: "Cesar",
		age:  32,
	}

	fmt.Println(user)
	fmt.Println(user.name)
	fmt.Println(user.age)

}

type user struct {
	name string
	age  int
}

```
--- 