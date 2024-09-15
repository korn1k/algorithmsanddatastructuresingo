package recursion

import "fmt"

func fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	
	if n == 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func FibonacciRecursionExample() {
	fmt.Println("Fibonacci recursion example")

	num := 10
	fmt.Printf("Fibonacci number at position %d is %d\n", num, fibonacci(num))
}