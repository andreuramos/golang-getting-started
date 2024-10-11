package main

import "fmt"

func main() {
    // Initialize a map for the integer values
    ints := map[string]int64{
        "first":  34,
        "second": 12,
    }

    // Initialize a map for the float values
    floats := map[string]float64{
        "first":  35.98,
        "second": 26.99,
    }

    fmt.Printf("Non-Generic Sums: %v and %v\n",
        SumInts(ints),
        SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		genericSum(ints),
		genericSum(floats))
}

func genericSum[K comparable, V int64 | float64](values map[K]V) V {
	var sum V
	for _, value := range(values) {
		sum += value
	}
	return sum
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, value := range(m) {
		s += value
	}
	return s
}

func SumFloats(values map[string]float64) float64 {
	var sum float64
	for _, value := range(values) {
		sum += value
	}
	return sum
}