package dia1

import "fmt"

// Um banco deseja conceder empréstimos a seus clientes, mas nem todos podem acessá-los.
// Para isso, possui certas regras para saber a qual cliente pode ser concedido. Apenas
// concede empréstimos a clientes com mais de 22 anos, empregados e com mais de um ano
// de atividade. Dentro dos empréstimos que concede, não cobra juros para quem tem um
// salário superior a US$ 100.000.
// É necessário fazer uma aplicação que possua essas variáveis e que imprima uma mensagem
// de acordo com cada caso.
// Dica: seu código deve ser capaz de imprimir pelo menos 3 mensagens diferentes.

func Loan(age int, employed bool, acitivity int, salary float32) {
	if age >= 22 && acitivity >= 1 && employed {
		if salary >= 100000 {
			fmt.Println("Cliente isento de juros")
		} else {
			fmt.Println("Cliente elegível")
		}
	} else {
		fmt.Println("Cliente não elegível")
	}
}
