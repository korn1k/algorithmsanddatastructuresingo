package recursion

import "fmt"

func factorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * factorial(n-1)
}

func FactorialRecursionExample() {
	fmt.Println("Factorial recursion example")

	num := 5
	fmt.Printf("Factorial of %d is %d\n", num, factorial(num))
}