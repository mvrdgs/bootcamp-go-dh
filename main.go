package main

import (
	"fmt"

	"github.com/mvrdgs/bootcamp-go-dh/dia4"
)

func main() {
	user := dia4.User{
		Nome:      "teste",
		Sobrenome: "silva",
		Idade:     19,
		Email:     "teste@teste.com",
		Senha:     "12345",
	}

	fmt.Println(user)

	user.MudarNome("testado")
	user.MudarSobrenome("silavado")
	user.MudarIdade(200)
	user.MudarEmail("testado@testado.com")
	user.MudarSenha("senhatestada123")

	fmt.Println(user)
}
