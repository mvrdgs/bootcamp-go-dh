package main

import (
	"github.com/mvrdgs/bootcamp-go-dh/dia2"
	_ "github.com/mvrdgs/bootcamp-go-dh/dia2"
)

func main() {
	aluno := dia2.Aluno{
		Nome:           "Maurício",
		Sobrenome:      "Viegas",
		RG:             44293849234,
		DataDeAdmissao: "16/05/2022",
	}

	aluno.Detalhamento()
}
