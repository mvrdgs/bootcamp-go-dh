package dia1

import "fmt"

func Employees() {
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Println(employees["Benjamin"])

	counter := 0

	for _, age := range employees {
		if age >= 21 {
			counter += 1
		}
	}

	fmt.Printf("%d funcionários possuem 21 anos ou mais\n", counter)

	employees["Federico"] = 25

	delete(employees, "Pedro")

	fmt.Println(employees)
}
