package dia6

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile() {
	file, err := os.Open("customers.txt")

	if err != nil {
		panic("o arquivo indicado não foi encontrado ou está danificado")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
