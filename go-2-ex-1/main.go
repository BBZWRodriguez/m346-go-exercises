package main

import "fmt"

type FullName struct {
	// TODO: add fields
	firstName string
	lastName  string
}

// TODO: declare a structure for birth date
type birthDate struct {
	dayOfBirth   int
	monthOfBirth int
	yearOfBirth  int
}

type Profile struct {
	// TODO: embed full name and birth date information
	FullName
	birthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		FullName:         FullName{firstName: "Dylan ", lastName: "Rodriguez"},
		birthDate:        birthDate{dayOfBirth: 1, monthOfBirth: 1, yearOfBirth: 2000},
		NumberOfSiblings: 1,        // TODO: adjust
		ZodiacSign:       '\u264b', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
