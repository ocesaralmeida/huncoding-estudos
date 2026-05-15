## Goroutines

- Uma goroutine é um encademanento leve gerenciado pelo runtime do Go.
- para iniciar um novo goroutine execuntando, adicione a palavra-chave `go` abtes da chamada da função `go add(a, b)`


### Implementação na prática goroutines

#### 1. antes de implementar
```go
package main

import (
	"fmt"
	"time"
)

func fun(value string) {
	for i := 0; i < 3; i++ {
		fmt.Println(value)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("Direct call")

	// TODO: write goroutines with differents variants for function call

	// TODO: goroutine function call

	// TODO: goroutine with anonymous value call

	// TODO: wait for goroutine to end

	fmt.Println("Done...")
}
```

- a implementação acima gera esse output:

```
Direct call
Direct call
Direct call
Done...
```

- agora utilizando `goroutine`:

```go
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

func main() {
	// Direct call
	fun("Direct call")

	// TODO: write goroutines with differents variants for function call

	// TODO: goroutine function call
	go fun("goroutine - 1")

	// TODO: goroutine with anonymous value call

	// TODO: wait for goroutine to end

	fmt.Println("Done...")
}
```

- vai gerar o mesmo outpud a última vez:

```
Direct call
Direct call
Direct call
Done...
```

##### por que isso ocorre? por que usando go fun("goroutine - 1") não gerou um output diferente?
- uma goroutine executa concorrentemente a outra, uma goroutine não conhece a outra(pelo menos nesse cenário)
- uma goroutine não vai esperar a outra executar, nem mesmo se for a gorotuine main
- se você executa uma goroutine a partir da goroutine main, ela só vai esperar essa goroutine terminar se fizer algo para isso acontecer(watiGroup, channel)
- dentro da função `fun()` ele demora 1milisegundo paraexecutar toda vez que ele printa
- a função `main()` chama a função `fun()`, printa a mensagem, e quando chamar a goroutine, não tem mais nada para fazer
- a main finaliza a execução antes da goroutine executar, a `main()` não espera a `goroutine`
- isso acontece pois o processo acaba antes da goroutine conseguir printar o valor
- colocando um `time.sleep()`, podemos fazer o programa "esperar" a execução da goroutine:

```go
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

func main() {
	// Direct call
	fun("Direct call")

	// TODO: write goroutines with differents variants for function call

	// TODO: goroutine function call
	go fun("goroutine - 1")

	// TODO: goroutine with anonymous value call

	// TODO: wait for goroutine to end
	time.Sleep(5 * time.Second)

	fmt.Println("Done...")
}
```

- a implementação com `time.Sleep()` gera esse output:

```
Direct call
Direct call
Direct call
goroutine - 1
goroutine - 1
goroutine - 1
Done...
```

- ao chamar a função fun em outra goroutine, o output vai sair "bagunçado", pois uma goroutine não tem conhecimento de outra goroutine.
- lembrando que uma goroutine executa independente de outras goroutines
- veja a implementação abaixo:

```go
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

func main() {
	// Direct call
	fun("Direct call")

	// TODO: write goroutines with differents variants for function call
	fgx := fun
	go fgx("goroutine - 2")

	// TODO: goroutine function call
	go fun("goroutine - 1")

	// TODO: goroutine with anonymous value call

	// TODO: wait for goroutine to end
	time.Sleep(5 * time.Second)

	fmt.Println("Done...")
}
```

- ao executar o programa duas vezes, ele mudou a ordem dos outputs:

```
execução 1:
Direct call
Direct call
Direct call
goroutine - 1
goroutine - 2
goroutine - 2
goroutine - 1
goroutine - 1
goroutine - 2
Done...
-------------
execução 2:
Direct call
Direct call
Direct call
goroutine - 1
goroutine - 2
goroutine - 1
goroutine - 2
goroutine - 2
goroutine - 1
Done...
-------------
```

- fazendo a mesma implementação com funções anônimas, o resultado é o mesmo, output "bagunçados":

```go
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

func main() {
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
```

```
execução 1:
Direct call
Direct call
Direct call
goroutine - 3
goroutine - 2
goroutine - 1
goroutine - 1
goroutine - 2
goroutine - 3
goroutine - 2
goroutine - 3
goroutine - 1
Done...
-------------
execução 2:
Direct call
Direct call
Direct call
goroutine - 3
goroutine - 2
goroutine - 3
goroutine - 1
goroutine - 3
goroutine - 1
goroutine - 2
goroutine - 2
goroutine - 1
Done...
```

#### Como faríamos caso seja necessário que todas as goroutines printassem o mesmo valor, e se esse valor for alterado, as goroutines devem printar também  o valor alterado?

## Usando ponteiros com goroutines
- podemos mudar o valor de uma váriável durante a execução de uma goroutine com pointeiros, veja o exemplo abaixo:

```go
package main

import (
	"fmt"
	"time"
)

func funPointer(value *string) {
	for {
		fmt.Println(*value)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var test string = "test"
	var pointTest *string = &test

	go funPointer(pointTest)

	time.Sleep(time.Second)

	*pointTest = "HunCoding"

	time.Sleep(3 * time.Second)
}
```

esse exemplo produziu esse output:

```
test
test
HunCoding
HunCoding
HunCoding
HunCoding
HunCoding
HunCoding
```