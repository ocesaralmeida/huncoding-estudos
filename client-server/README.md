# Usando goroutines com server/clients


## para um client
- E se quisermos que o código receba várias requisições de rede dentro de um server?
- o exemplo abaixo mostra um server ouvindo a porta 8000, se der erro ele dá um fatal nesse erro
- enquanto ele receber conexão, ele vai abrir uma nova variável de conexão e vai ficar printando "response from server"

```go
package server

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal(err)
		}

		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, "response from server\n")
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
```

- Enquanto o server está de pé, o código do client roda
- no exemplo abaixo, o client vai fazr uma conexão TCP no localhost:8000(a mesma conexão aberta no lado do server)
- ele vai copiar pra dentro do writer(nesse caso, o terminal) o que for retornado pra ele(nesse caso, o "response from server" do server)

```go
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8000")
	fmt.Println("retorno da conexão: \n", conn)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	mustCopy(os.Stdout, conn)

}

func mustCopy(ds io.Writer, src io.Reader) {
	if _, err := io.Copy(ds, src); err != nil {
		log.Fatal(err)
	}
}
```

esse é o retorno do server e do client:

```
retorno da conexão: 
 &{{0xc000150000}}
response from server
response from server
response from server
response from server
response from server
response from server
response from server
response from server
```

## para varios clients
- E se quisermos receber vários clients conectados?
- ao tentar rodar o client em mais um terminal, vão vai retornar nada nos outros terminais, só no primeiro 
- o código clientestá aberto apenas para receber uma conexão, ele abre a conexão e o código fica travado cuidando dessa conexão
- ele não fica concorrente cuidando da conexão e esperando outras conexões, por isso ao parar um dos clients, o outro começa a printar, ele cuida de uma por vez
- para fazer todas as conexões aceitarem as requisiões ao mesmo tempo, precisamos utilizar goroutines

## gorutines para multiplas conexões

- no exemplo anterior do server, faziamos o `handleConn(conn)`apenas uma vez, pois ele fica respondendo só essa conexão
- adicionando o `handleConn(conn)` como uma goroutine, cada nova conexão se torna uma goroutine, assim permitindo responder vários clients diferentes
- veja o exemplo abaixo, utilizando o `handleConn(conn)` como goroutine:

```go
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, "response from server\n")
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
```

- dessa maneira, cada conexão client é executada de maneira concorrente