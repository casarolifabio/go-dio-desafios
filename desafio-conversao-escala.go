package main

import "fmt"

/*
 C = K -273
 100 = K - 273
 K = 100 + 273
 K = 373
*/
const ebulicaoC = 100.0

func main() {
	tempC := ebulicaoC
	tempK := tempC + 273.0

	fmt.Printf("A temperatura de ebulição da água em K = %g , temperatura de ebulição da água em °C =%g.", tempK, tempC)
}
