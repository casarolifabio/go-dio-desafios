package main

import "fmt"

func main() {
	// Criação de um array de tamanho fixo [100] para armazenar números de 1 a 100
	numeros := make([]int, 100)

	for i := 1; i <= len(numeros); i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}

	}
}
