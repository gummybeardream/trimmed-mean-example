package main

import (
	"fmt"
	"log"

	trimmedmean "github.com/gummybeardream/trimmed-mean"
)

func main() {
	integers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	integerMean, err := trimmedmean.Mean(integers, 0.10)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Integer trimmed mean: %.2f\n", integerMean)

	floats := []float64{1.5, 2.5, 3.5, 4.5, 100.5}

	floatMean, err := trimmedmean.Mean(floats, 0.20)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Floating-point trimmed mean: %.2f\n", floatMean)
}
