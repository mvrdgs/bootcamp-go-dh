package dia2

import "fmt"

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

// animalDog, msg := Animal(dog)
// animalCat, msg := Animal(cat)

// var amount float64
// amount+= animaldog(5)
// amount+= animalCat(8)

func animalCat(qty int) float64 {
	return float64(qty) * 5
}

func animalDog(qty int) float64 {
	return float64(qty) * 10
}

func animalHamster(qty int) float64 {
	return float64(qty) * .250
}

func animalTarantula(qty int) float64 {
	return float64(qty) * .150
}

func Animal(animal string) (func(int) float64, error) {
	switch animal {
	case dog:
		return animalCat, nil
	case cat:
		return animalDog, nil
	case hamster:
		return animalHamster, nil
	case tarantula:
		return animalTarantula, nil
	default:
		return nil, fmt.Errorf("%s não é um animal cadastrado", animal)
	}
}

func CalcFood() {
	var amount float64
	amount += animalDog(5)
	amount += animalCat(8)
	amount += animalHamster(5)
	amount += animalTarantula(6)

	fmt.Printf("A quantidade de ração necessária é: %.2f\n", amount)
}
