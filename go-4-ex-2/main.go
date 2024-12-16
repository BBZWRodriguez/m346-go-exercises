package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt

func computeHypotenuse(a float64, b float64) float64 {

	hypotenuse := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
	return hypotenuse
}

func main() {
	// TODO: call the function computeHypotenuse

	// Test 1: Beispiel mit Kathetenlängen 11.4 und 18.0
	a := 11.4
	b := 18.0
	hypotenuse := computeHypotenuse(a, b)
	fmt.Printf("Test 1 - Die Hypotenuse für a=%.2f und b=%.2f beträgt: %.2f\n", a, b, hypotenuse)

	// Test 2: Beispiel mit Kathetenlängen 3 und 4
	a = 3
	b = 4
	hypotenuse = computeHypotenuse(a, b)
	fmt.Printf("Test 2 - Die Hypotenuse für a=%.2f und b=%.2f beträgt: %.2f\n", a, b, hypotenuse)

	// Test 3: Beispiel mit Kathetenlängen 5 und 12
	a = 5
	b = 12
	hypotenuse = computeHypotenuse(a, b)
	fmt.Printf("Test 3 - Die Hypotenuse für a=%.2f und b=%.2f beträgt: %.2f\n", a, b, hypotenuse)
}
