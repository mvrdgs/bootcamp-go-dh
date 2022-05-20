package dia2

import (
	"fmt"
	"sort"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func minFunc(valores ...int) int {
	sort.Ints(valores)
	return valores[0]
}

func maxFunc(valores ...int) int {
	sort.Ints(valores)
	return valores[len(valores)-1]
}

func averageFunc(valores ...int) int {
	soma := 0
	for _, valor := range valores {
		soma += valor
	}

	return soma / len(valores)
}

func operation(operacao string) (func(...int) int, error) {
	switch operacao {
	case minimum:
		return minFunc, nil
	case maximum:
		return maxFunc, nil
	case average:
		return averageFunc, nil
	default:
		return nil, fmt.Errorf("Operação %s é inválida", operacao)
	}
}

func Calc() {
	minFunc, err := operation(minimum)
	averageFunc, err2 := operation(average)
	maxFunc, err3 := operation(maximum)

	if err != nil {
		fmt.Println(err)
	}

	if err2 != nil {
		fmt.Println(err)
	}

	if err3 != nil {
		fmt.Println(err)
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println(minValue)
	fmt.Println(averageValue)
	fmt.Println(maxValue)
}
