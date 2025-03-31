package main

import "fmt"

func main() {
	num := 9
	fmt.Println("Factorial of", num, "using iterative approach", factorial(num))
	fmt.Println("Factorial of", num, "using recursive approach", factorialI(num))
}

func factorial(num int) int {
	total := 1
	for i := range num {
		total *= i + 1
	}
	return total
}

func factorialI(num int) int {
	return num * factorial(num-1)
}
