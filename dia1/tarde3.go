package dia1

import "fmt"

func PrintMonth(month int) {
	switch {
	case month == 1:
		fmt.Println("Janeiro")
	case month == 2:
		fmt.Println("Fevereiro")
	case month == 3:
		fmt.Println("Março")
	case month == 4:
		fmt.Println("Abril")
	case month == 5:
		fmt.Println("Maio")
	case month == 6:
		fmt.Println("Junho")
	case month == 7:
		fmt.Println("Julho")
	case month == 8:
		fmt.Println("Agosto")
	case month == 9:
		fmt.Println("Setembro")
	case month == 10:
		fmt.Println("Outubro")
	case month == 11:
		fmt.Println("Novembro")
	case month == 12:
		fmt.Println("Dezembro")
	default:
		fmt.Println("Mês inválido")
	}
}
