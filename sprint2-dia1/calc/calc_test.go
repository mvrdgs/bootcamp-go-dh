package calc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// go test ./sprint2-dia1/calc/
func TestSum(t *testing.T) {
	num1 := 3
	num2 := 5
	expectedResult := 8

	result := Sum(num1, num2)

	assert.Equal(t, expectedResult, result, "devem ser iguais")
}

func TestSub(t *testing.T) {
	num1 := 3
	num2 := 5
	expectedResult := -2

	result := Sub(num1, num2)

	assert.Equal(t, expectedResult, result, "devem ser iguais")
}

func TestMult(t *testing.T) {
	num1 := 3
	num2 := 5
	expectedResult := 15

	result := Mult(num1, num2)

	assert.Equal(t, expectedResult, result, "devem ser iguais")
}

func TestDiv(t *testing.T) {
	num1 := 10
	num2 := 5
	expectedResult := 2

	result, _ := Div(num1, num2)

	assert.Equal(t, expectedResult, result, "devem ser iguais")
}

func TestDivFail(t *testing.T) {
	num1 := 10
	num2 := 0

	_, err := Div(num1, num2)

	assert.Errorf(t, err, "deve retornar erro se o denominador for 00")
	assert.Equal(t, err.Error(), "O demoninador não pode ser 0")
}
