package dia3

import (
	"fmt"
	"log"
	"os"
)

// Uma empresa que vende produtos de limpeza precisa de:
// 1. Implementar uma funcionalidade para guardar um arquivo de texto, com a informação
// de produtos comprados, separados por ponto e vírgula (csv).
// 2. Deve possuir o ID do produto, preço e a quantidade.
// 3. Estes valores podem ser hardcodeados ou escritos manualmente em uma variável.

func SaveProduct(id, qtd int, price float64) {
	info := fmt.Sprintf("%d %d %.2f;\n", id, qtd, price)

	f, err := os.OpenFile("./ProductInfo.csv",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(string(info)); err != nil {
		log.Println(err)
	}
}
