package main

import (
	"fmt"

	"github.com/mvrdgs/bootcamp-go-dh/dia4"
)

func main() {
	user := dia4.Usuario{
		Nome:      "user1",
		Sobrenome: "sobreuser1",
	}

	user2 := dia4.Usuario{
		Nome:      "user2",
		Sobrenome: "sobreuser2",
	}

	produto1 := dia4.NovoProduto("MacBook", 20000.00)

	dia4.AdicionarProduto(&user, produto1, 10)
	dia4.AdicionarProduto(&user2, produto1, 50)

	fmt.Println(user)
	fmt.Println(user2)
}
