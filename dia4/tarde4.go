package dia4

import (
	"fmt"
	"math/rand"
	"time"
)

// Uma empresa de sistemas precisa analisar que algoritmos de ordenamento utilizar para seus
// serviços.
// Para eles é necessário instanciar 3 arrays com valores aleatórios não ordenados
// - uma matriz de inteiros com 100 valores
// - uma matriz de inteiros com 1000 valores
// - uma matriz de inteiros com 10.000 valores

// Para instanciar as variáveis, utilize o rand:
// package main
// import (
// "math/rand"
// )

// func main() {
// variavel := rand.Perm(100)
// variave2 := rand.Perm(1000)
// variave3 := rand.Perm(10000)
// }
// Cada um deve ser ordenado por:
// - Inserção
// - grupo
// - seleção

// Uma rotina para cada execução de classificação
// Tenho que esperar terminar a ordenação de 100 números para continuar com a ordenação de
// 1000 e depois a ordenação de 10000.
// Por fim, devo medir o tempo de cada um e mostrar o resultado na tela, para saber qual
// ordenação ficou melhor para cada arranjo.

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
			j = j - 1
		}
	}

	return arr
}

func selectionSort(arr []int) []int {
	length := len(arr)

	for i := 0; i < length; i++ {
		minIndex := i
		for j := i; j < length; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	return arr
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	first := mergeSort(arr[:len(arr)/2])
	second := mergeSort(arr[len(arr)/2:])
	return merge(first, second)
}

func merge(a, b []int) []int {
	final := []int{}
	right, left := 0, 0

	for right < len(a) && left < len(b) {
		if a[right] < b[left] {
			final = append(final, a[right])
			right += 1
		} else {
			final = append(final, b[left])
			left += 1
		}
	}

	for right < len(a) {
		final = append(final, a[right])
		right += 1
	}

	for left < len(b) {
		final = append(final, b[left])
		left += 1
	}

	return final
}

func compare(variavel []int) {
	ch := make(chan string)
	go func() {
		start := time.Now()
		insertionSort(variavel)
		ch <- fmt.Sprint("insertion in ", time.Since(start))
	}()
	go func() {
		start := time.Now()
		selectionSort(variavel)
		ch <- fmt.Sprint("selection in ", time.Since(start))
	}()
	go func() {
		start := time.Now()
		mergeSort(variavel)
		ch <- fmt.Sprint("merge in ", time.Since(start))
	}()

	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
}

func Tarde4() {
	variavel := rand.Perm(100)
	variavel2 := rand.Perm(1000)
	variavel3 := rand.Perm(10000)

	fmt.Println("100 numeros")
	compare(variavel)
	fmt.Println("\n1000 numeros")
	compare(variavel2)
	fmt.Println("\n10000 numeros")
	compare(variavel3)
}
