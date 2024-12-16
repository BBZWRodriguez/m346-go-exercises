//`go-4-ex-1/main.go`: Eine Note `N` von 1 (schlechteste) bis 6 (beste) kann
//linear anhand der erreichten Punktzahl `E` und maximalen Punktezahl `M`
//berechnet werden:

package main

import "fmt"

// TODO: implement the function computeGrade
func computeGrade(gotPoints, maxPoints float64) float64 {

	grade := (gotPoints/maxPoints)*5 + 1

	return grade
}

func main() {
	// TODO: call the function computeGrade
	gotPoints := 17.5
	maxPoints := 28.0

	grade := computeGrade(gotPoints, maxPoints)
	fmt.Printf("Die Note für %.2f/%.2f Punkte ist: %.2f\n", gotPoints, maxPoints, grade)

}
