package main

import "fmt"

// TODO: declare a type for Student (with first and last name)

type Student struct {
	FirstName string
	LastName  string
}

// TODO: declare a type for Class (consisting of multiple students)

type Class struct {
	Name     string
	Students []Student
}

func main() {

	// Beispiel-Schüler für Klasse 1
	class1 := Class{
		Name: "Class 1",
		Students: []Student{
			{"Alice", "Smith"},
			{"Bob", "Johnson"},
			{"Charlie", "Brown"},
		},
	}

	// Beispiel-Schüler für Klasse 2
	class2 := Class{
		Name: "Class 2",
		Students: []Student{
			{"Diana", "Davis"},
			{"Eve", "Miller"},
			{"Frank", "Wilson"},
		},
	}

	// Module mit Klassen
	modules := map[string][]Class{
		"346": {class1, class2},
		"247": {class1},
		"158": {class2},
	}

	// Daten ausgeben
	fmt.Println("Klassenverwaltung:")
	fmt.Println("Klassen und Schüler:")
	fmt.Println(class1)
	fmt.Println(class2)

	fmt.Println("\nModule und zugehörige Klassen:")
	for module, classes := range modules {
		fmt.Printf("Modul %s:\n", module)
		for _, class := range classes {
			fmt.Printf("- %s\n", class.Name)
		}
	}
	// TODO: declare a map of modules being attended by multiple classes
	// TODO: output everything using fmt.Println()
}
