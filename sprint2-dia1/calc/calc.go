package calc

import "fmt"

func Sum(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mult(a, b int) int {
	return a * b
}

func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("O demoninador não pode ser 0")
	}
	return a / b, nil
}
