# Fundamentos Golang I - Variáveis

## tipos de variáveis

```go
func main() {
	testString := "25.33"  //string
	fmt.Printf("%T", testString)

	testInt := 30 //int
	fmt.Printf("%T", testInt)

	testFloat := 23.11 //float64
	fmt.Printf("%T", testFloat)

	testMap := map[string]string{"teste": "pos1"}//map[string]string
	fmt.Printf("%T", testMap)
}
```
---

## variáveis públicas e privadas

```go
func test() {} //privada

func Test() {} //publixa, pode ser acessada em outros pacotes e modulos
```
---

## Tipos e tipo "any"(interface{})
Tipo `interface{}` aceita qualquer tipo, é como o `any` do javascript 

```go
func main() {
	var test interface{}
 
    //atribuindo tipos diferentes a mesma variável
	test = "teste"
	test = 20
	test = 20.5

	var testJson = map[interface{}]interface{}{
		"test12345": "Teste",
		"test2":     10,
		4.3:         true,
	}

	test = testJson

    //exibindo o tipo de variável
	fmt.Printf("%T", test)
}
```

---

## Arrays e Slices
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