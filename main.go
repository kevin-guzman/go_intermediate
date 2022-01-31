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

func main() {
	fmt.Println(sum(1, 2, 1, 1))
	printNames("Kevin", "Manolo", "Juanita", "Ashole")
}
