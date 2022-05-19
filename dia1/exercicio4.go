package dia1

import "fmt"

func TiposDeDados() {
	var sobrenome string = "Silva"
	var idade int = 25
	boolean := false
	var salario float32 = 4585.90
	var nome string = "Fellipe"

	fmt.Printf("Nome completo: %s %s\nIdade: %d\nSalário: %.2f\nBoolean aleatório: %t\n", nome, sobrenome, idade, salario, boolean)
}
