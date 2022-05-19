package dia1

import "fmt"

func wordLength(word string) {
	fmt.Printf("A palavra tem %d letras\n", len(word))
}

func wordLetters(word string) {
	for i := range word {
		fmt.Printf("%c\n", word[i])
	}
}

func WordLetters(word string) {
	wordLength(word)
	wordLetters(word)
}
