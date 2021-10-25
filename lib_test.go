package lib

import (
	"fmt"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestMin(t *testing.T) {
	var numbers [10]float64 = [10]float64{2, 3, 5, 7, 11, 13, -1, 9, -5, 0}
	var count int = 2
	min, err := Min(numbers[:], 2)
	if err == nil && len(min) != count && min[0] != -5 && min[1] != -1 {
		t.Fatalf(`min = %v, want {-5, -1}`, min)
	}

	fmt.Printf("%v", numbers)
	fmt.Printf("%v", min)
}

func TestMax(t *testing.T) {
	var numbers [10]float64 = [10]float64{2, 3, 5, 7, 11, 13, -1, 9, -5, 0}
	var count int = 2
	max, err := Max(numbers[:], count)

	if err == nil && len(max) != count && max[0] != 11 && max[1] != 13 {
		t.Fatalf(`min = %v, want {11, 13}`, max)
	}

	fmt.Printf("%v", numbers)
	fmt.Printf("%v", max)
}

func TestAverage(t *testing.T) {
	var numbers [10]float64 = [10]float64{2, 3, 5, 7, 11, 13, -1, 9, -5, 0}
	expected := float64(4.4)
	actual := Average(numbers[:])

	if actual != expected {
		t.Fatalf(`avg = %v, want %v`, actual, expected)
	}

	fmt.Printf("%v", numbers)
	fmt.Printf("%v", actual)
}

func TestMedian(t *testing.T) {
	var numbers [10]float64 = [10]float64{2, 3, 5, 7, 11, 13, -1, 9, -5, 0}
	expected := float64(4)
	actual := Median(numbers[:])

	if actual != expected {
		t.Fatalf(`median = %v, want %v`, actual, expected)
	}

	var numbers2 [9]float64 = [9]float64{3, 5, 7, 11, 13, -1, 9, -5, 0}
	expected = float64(5)
	actual = Median(numbers2[:])

	fmt.Printf("%v", numbers2)
	fmt.Printf("%v", actual)

	if actual != expected {
		t.Fatalf(`median = %v, want %v`, actual, expected)
	}
}

func TestPercentile(t *testing.T) {
	var numbers [5]float64 = [5]float64{15, 20, 35, 40, 50}
	q := 30
	expected := float64(20)
	actual := Percentile(numbers[:], q)

	if actual != expected {
		t.Fatalf(`percentile = %v, want %v`, actual, expected)
	}

	q = 50
	expected = float64(35)
	actual = Percentile(numbers[:], q)

	if actual != expected {
		t.Fatalf(`percentile = %v, want %v`, actual, expected)
	}
}
