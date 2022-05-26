package dia6

import (
	"fmt"
	"github.com/google/uuid"
)

func GenerateId() string {
	id := uuid.New()
	return fmt.Sprint(id)
}

func Exercicio2() {
	id := GenerateId()

	if id == "" {
		panic("id não pode ser vazio")
	}

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	ReadFile()
}
