package dia5

import "fmt"

// Repita o processo anterior, mas agora implementando "fmt.Errorf()", para que a mensagem de
// erro receba como parâmetro o valor de "salario", indicando que não atinge o mínimo
// tributável (a mensagem exibida pelo console deve dizer : "erro: o mínimo tributável é 15.000 e
// o salário informado é: [salario]”, onde [salario] é o valor do tipo int passado pelo parâmetro).

func checkTax3(salario int) (string, error) {
	if salario < 15000 {
		return "", fmt.Errorf("erro: o mínimo tributável é 15.000 e o salário informado é %d", salario)
	}

	return "Necessário pagamento de imposto", nil
}

func Tarde3(salario int) {
	msg, err := checkTax3(salario)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(msg)
}
