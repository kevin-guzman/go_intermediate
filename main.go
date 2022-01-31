package main

import (
	"fmt"
)

func sum(values ...int) int {
	total := 0
	for _, number := range values {
		total += number
	}
	return total
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func getValues(x int) (double int, triple int, cuadruple int) {
	double = 2 * x
	triple = 3 * x
	cuadruple = 4 * x
	return
}

func main() {
	fmt.Println(sum(1, 2, 1, 1))
	printNames("Kevin", "Manolo", "Juanita", "Ashole")
	// Retornos con nombre
	fmt.Println(getValues(2))
}
