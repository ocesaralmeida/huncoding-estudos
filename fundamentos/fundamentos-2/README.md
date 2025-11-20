# Fundamentos Golang II - Estruturas de Controle

## if - else - else if

### if
```go
package main

import "fmt"

func main() {
	if 3 > 2 {
		fmt.Println("true")
	}
}
```

### else
```go
package main

import "fmt"

func main() {
	if 3 < 2 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
```

### else if
```go
package main

import "fmt"

func main() {
	if 3 < 2 {
		fmt.Println("true")
	} else if 4 > 3 {
		fmt.Println("test")
	} else {
		fmt.Println("false")
	}
}
```
---

## if com init

```go
package main

import (
	"fmt"
)

func main() {

	if retorno := test(); retorno == "test" {
		fmt.Print("retorno")
	}
}

func test() string {
	return "test"
}
```
---

## Swtich
```go
package main

import (
	"fmt" 
)

func main() {

	var test string = "test2"

	switch test {

	case "test":
		fmt.Println("CAIU NA PRIMEIRA CONDIÇÃO")

	case "test2":
		fmt.Println("CAIU NA SEGUNDA POSIÇÃO")
	}
}
```

```go
package main

import (
	"fmt"
)

func main() {

	var test string = "test2"

	switch test {

	case "test", "test2", "test999":
		fmt.Println("CAIU NA PRIMEIRA CONDIÇÃO")

	case "test3":
		fmt.Println("CAIU NA SEGUNDA POSIÇÃO")
	}
}
```

---

## Default
```go
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
```
--- 