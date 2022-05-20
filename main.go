package main

import (
	"github.com/mvrdgs/bootcamp-go-dh/dia1"
	"github.com/mvrdgs/bootcamp-go-dh/dia2"
)

func main() {
	dia1.Imprime()
	// dia1.Clima(16, 1013.25, 84)
	// dia1.TiposDeDados()
	// dia1.WordLetters("go meli go")
	// dia1.Loan(23, true, 2, 50000)
	// dia1.Loan(15, false, 2, 100000)
	// dia1.Loan(23, true, 2, 100000)
	// dia1.PrintMonth(9)
	// dia1.PrintMonth(13)
	// dia1.Employees()
	// fmt.Printf("%.2f\n", dia2.SalaryTax(50000))
	// fmt.Printf("%.2f\n", dia2.SalaryTax(150000))
	dia2.Media(7, 5, 6)
	dia2.Media(7, 5, 6, -2)
}
