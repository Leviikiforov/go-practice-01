package main

import "fmt"

func main() {
	a := 1
	b := 1
	c := 2
	sum := a + b + c
	sr := float64(sum) / 3

	fmt.Printf("Среднее значение времени: %.2f", sr)
}
