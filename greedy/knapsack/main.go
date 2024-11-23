package main

import (
	"fmt"
)

// Item represents an item with a value and weight
type Item struct {
	value, weight int
	ratio         float64
}

// FractionalKnapsack solves the problem using the Greedy Algorithm
func FractionalKnapsack(items []Item, capacity int) float64 {
	totalValue := 0.0

	// TODO: implement

	return totalValue
}

func main() {
	items := []Item{
		{value: 60, weight: 10},
		{value: 100, weight: 20},
		{value: 120, weight: 30},
	}
	capacity := 50

	result := FractionalKnapsack(items, capacity)
	fmt.Printf("Maximum value in Fractional Knapsack: %.2f\n", result)
}
