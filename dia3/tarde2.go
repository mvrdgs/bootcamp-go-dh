package dia3

import (
	"encoding/csv"
	"fmt"
	"os"
)

// A mesma empresa precisa ler o arquivo armazenado, para isso exige que:
// - Seja impresso na tela os valores tabelados, com título ( à esquerda para o ID e à
// direita para o Preço e Quantidade), o preço, a quantidade e abaixo do preço o total
// deve ser exibido (somando preço por quantidade)

// Exemplo:
// ID Preco Quantidade
// 111223 30012.00 1
// 444321 1000000.00 4
// 434321 50.50 1

// 4030062.50

type data struct {
	id    string
	preco string
	qty   string
}

func ReadProductFile(fileName string) {
	csvFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, line := range csvLines {
		emp := data{
			id:    line[0],
			preco: line[1],
			qty:   line[2],
		}

		fmt.Println(emp)
	}
}
