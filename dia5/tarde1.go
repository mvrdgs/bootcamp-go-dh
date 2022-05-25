package dia5

import (
	"fmt"
)

// 1. Em sua função “main”, defina uma variável chamada “salario” e atribua um valor do
// tipo “int”.
// 2. Crie um erro personalizado com uma struct que implemente “Error()” com a
// mensagem “erro: O salário digitado não está dentro do valor mínimo" em que seja
// disparado quando “salário” for menor que 15.000. Caso contrário, imprima no
// console a mensagem “Necessário pagamento de imposto”.

type Erro struct {
	msg string
}

func (e *Erro) Error() string {
	return e.msg
}

func checkTax(salario int) (string, error) {
	if salario < 15000 {
		return "", &Erro{"O salário digitado não está dentro do valor mínimo"}
	}

	return "Necessário pagamento de imposto", nil
}

func Tarde1(salario int) {
	msg, err := checkTax(salario)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(msg)
}
