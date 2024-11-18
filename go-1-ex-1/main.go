package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	firstName := "Dylan"
	lastName := "Rodriguez"
	dayOfBirth := 1
	monthOfBirth := 1
	yearOfBirth := 2000
	numberOfSiblings := 3
	heightInMeters := 1.8
	zodiacSign := '\u264b'
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
