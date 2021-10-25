package lib

import (
	"errors"
	"math"
	"sort"
)

func sortUs(numbers []float64) {
	sort.Float64s(numbers)
}

func Min(numbers []float64, count int) ([]float64, error) {
	sortUs(numbers)

	if count > len(numbers) {
		return numbers, errors.New("the number of elements requested is too large")
	}

	return numbers[0:count], nil
}

func Max(numbers []float64, count int) ([]float64, error) {
	sortUs(numbers)

	if count > len(numbers) {
		return numbers, errors.New("the number of elements requested is too large")
	}

	return numbers[len(numbers)-count:], nil
}

// average that avoid number overflowing during the calculations
// https://www.heikohoffmann.de/htmlthesis/node134.html
func Average(numbers []float64) float64 {
	sortUs(numbers)

	var average float64 = 0
	var denominator float64 = 1

	for _, value := range numbers {
		average += ((value - average) / denominator)
		denominator++
	}

	return average
}

func Median(numbers []float64) float64 {
	sortUs(numbers)

	var medianIndex int = len(numbers) / 2
	if (len(numbers) % 2) == 1 {
		return numbers[medianIndex]
	} else {
		return numbers[medianIndex]/2 + numbers[medianIndex-1]/2
	}
}

func Percentile(numbers []float64, q int) float64 {
	sortUs(numbers)

	var ordinalRank int = int(math.Ceil(float64(q*len(numbers)) / 100))
	var index int = ordinalRank - 1
	return numbers[index]
}
