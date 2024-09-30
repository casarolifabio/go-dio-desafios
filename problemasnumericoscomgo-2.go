package main

import "fmt"

func main() {
	numeros := make([]int, 100)

	for i := 1; i <= len(numeros); i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("Pin Pan")
			} else {
				fmt.Println("Pin")
			}
		}
		if i%5 == 0 {
			fmt.Println("Pan")
		}
	}
}
