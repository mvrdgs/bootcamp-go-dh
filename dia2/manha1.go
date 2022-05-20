package dia2

func SalaryTax(salary float64) (tax float64) {
	if salary > 50000 {
		tax = salary * .17
		return tax
	}
	tax = salary * .10
	return tax
}
