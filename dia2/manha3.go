package dia2

import "fmt"

const (
	categoryA = "A"
	categoryB = "B"
	categoryC = "C"
)

func calc(hours int, category string) int {
	switch category {
	case categoryA:
		salary := hours * 3000
		if hours > 160 {
			salary += int(float64(salary) * .5)
		}

		return salary
	case categoryB:
		salary := hours * 1500
		if hours > 160 {
			salary += int(float64(salary) * .2)
		}

		return salary
	default:
		salary := hours * 1000
		return salary
	}
}

func CalculateSalary() {
	fmt.Println(calc(162, categoryC))
	fmt.Println(calc(176, categoryB))
	fmt.Println(calc(172, categoryA))
}
