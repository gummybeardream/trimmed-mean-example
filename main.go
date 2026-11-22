package main

import (
	"fmt"
	"log"

	trimmedmean "github.com/gummybeardream/trimmed-mean"
)

const trimProportion = 0.05

func main() {
	integers, err := readIntegers("data/integers.csv")
	if err != nil {
		log.Fatalf("read integer data: %v", err)
	}

	floats, err := readFloats("data/floats.csv")
	if err != nil {
		log.Fatalf("read floating-point data: %v", err)
	}

	integerMean, err := trimmedmean.Mean(integers, trimProportion)
	if err != nil {
		log.Fatalf("calculate integer trimmed mean: %v", err)
	}

	floatMean, err := trimmedmean.Mean(floats, trimProportion)
	if err != nil {
		log.Fatalf("calculate floating-point trimmed mean: %v", err)
	}

	fmt.Println("Go trimmed mean validation")
	fmt.Printf(
		"Integer trimmed mean (5%%): %.10f\n",
		integerMean,
	)
	fmt.Printf(
		"Floating-point trimmed mean (5%%): %.10f\n",
		floatMean,
	)
}
