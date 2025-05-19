package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

func main() {
	sum := Add(3, 4)
	product := Multiply(3, 4)
	fmt.Printf("3 + 4 = %d\n", sum)
	fmt.Printf("3 * 4 = %d\n", product)
}
