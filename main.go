package main

import (
	"fmt"

	"github.com/mvrdgs/bootcamp-go-dh/dia2"
)

func main() {
	loja := dia2.NovaLoja()

	mouse := dia2.NovoProduto("Mouse", "Pequeno", 100.0)
	macBook := dia2.NovoProduto("MacBook Pro", "Medio", 20000.00)
	frigobar := dia2.NovoProduto("Frigobar", "Grande", 3000.00)

	fmt.Println("Custo mouse: ", mouse.CalcularCusto())
	fmt.Println("Custo Mac Book: ", macBook.CalcularCusto())
	fmt.Println("Custo Frigobar: ", frigobar.CalcularCusto())

	loja.Adicionar(mouse)
	loja.Adicionar(macBook)
	loja.Adicionar(frigobar)

	fmt.Println("Total: ", loja.Total())
}
