package dia5

import (
	"errors"
	"fmt"
)

// Faça a mesma coisa que no exercício anterior, mas reformule o código para que, em vez de
// “Error()”, seja implementado “errors.New()”.

func checkTax2(salario int) (string, error) {
	if salario < 15000 {
		return "", errors.New("O salário digitado não está dentro do valor mínimo")
	}

	return "Necessário pagamento de imposto", nil
}

func Tarde2(salario int) {
	msg, err := checkTax2(salario)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(msg)
}
