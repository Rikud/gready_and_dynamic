package main

import "fmt"

// Knapsack01 solves the 0/1 Knapsack problem using Dynamic Programming
func Knapsack01(values []int, weights []int, capacity int) int {
	var result int

	// TODO: implement

	return result
}

func main() {
	values := []int{60, 120, 100}
	weights := []int{10, 30, 20}
	capacity := 50

	result := Knapsack01(values, weights, capacity)
	fmt.Printf("Maximum value in 0/1 Knapsack: %d\n", result)
}
