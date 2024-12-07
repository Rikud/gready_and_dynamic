package main

import "fmt"

var memo map[int]int

func fib(n int) int {
	var result int

	// TODO: implement

	return result
}

// Fibonacci with tabulation
func fibTab(n int) int {
	var result int

	// TODO: implement

	return result
}

// Fibonacci with memoization
func fibMemo(n int, memo map[int]int) int {
	// TODO: implement

	return memo[n]
}

// Naive recursive Fibonacci
func fibNaive(n int) int {
	if n <= 1 {
		return n
	}
	return fibNaive(n-1) + fibNaive(n-2)
}

func main() {
	n := 10
	fmt.Printf("Fibonacci number F(%d) = %d\n", n, fib(n))
}
