package dia2

import "fmt"

func calcMedia(notas []int) (int, error) {
	media := 0
	for _, nota := range notas {
		if nota < 0 {
			return 0, fmt.Errorf("%d não é uma nota válida", nota)
		}
		media += nota
	}

	media = media / len(notas)

	return media, nil
}

func Media(notas ...int) {
	media, err := calcMedia(notas)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("A média do aluno é: %d\n", media)
	}
}
