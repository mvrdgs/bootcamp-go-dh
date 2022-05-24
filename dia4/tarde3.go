package dia4

import "fmt"

// Precisamos de 3 estruturas:
// - Produtos: nome, preço, quantidade.
// - Serviços: nome, preço, minutos trabalhados.
// - Manutenção: nome, preço.
// Precisamos de 3 funções:
// - Somar Produtos: recebe um array de produto e devolve o preço total (preço *
// quantidade).
// - Somar Serviços: recebe uma array de serviço e devolve o preço total (preço * média
// hora trabalhada, se ele não vier trabalhar por 30 minutos, ele será cobrado como se
// tivesse trabalhado meia hora).
// - Somar Manutenção: recebe um array de manutenção e devolve o preço total.

// Os 3 devem ser executados concomitantemente e ao final o valor final deve ser mostrado na
// tela (somando o total dos 3).

type Servico struct {
	Nome               string
	Preco              float64
	MinutosTrabalhados float64
}

type Manutencao struct {
	Nome  string
	Preco float64
}

func SomarProdutos(p []Produto, c chan<- float64) {
	total := 0.0

	for _, produto := range p {
		total += produto.Preco * float64(produto.Quantidade)
	}

	c <- total
}

func SomarServicos(s []Servico, c chan<- float64) {
	total := 0.0

	for _, servico := range s {
		total += servico.Preco * servico.MinutosTrabalhados
	}

	c <- total
}

func SomarManutencao(m []Manutencao, c chan<- float64) {
	total := 0.0

	for _, manutencao := range m {
		total += manutencao.Preco
	}

	c <- total
}

func Tarde3() {
	produtos := make([]Produto, 0)
	produtos = append(produtos, Produto{"MacBook", 20000.00, 5}, Produto{"Mouse", 100.00, 10})

	servicos := make([]Servico, 0)
	servicos = append(servicos, Servico{"Aula", 10.00, 100}, Servico{"Prova", 20.00, 50.00})

	manutencoes := make([]Manutencao, 0)

	manutencoes = append(manutencoes, Manutencao{"Conserto", 500.00}, Manutencao{"Conserto", 1500})

	c := make(chan float64, 3)

	go SomarProdutos(produtos, c)
	go SomarServicos(servicos, c)
	go SomarManutencao(manutencoes, c)

	total := 0.0

	for i := 0; i < 3; i++ {
		total += <-c
	}

	fmt.Println(total)
}
