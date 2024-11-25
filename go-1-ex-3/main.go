package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var eyes = rand.Intn(5) + 1
	var when = time.Now()

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stdout, "The dice shows", eyes, "eyes")

	// TODO: use fmt.Fprintln instead!
	fmt.Fprintln(os.Stdout, "The dice was rolled at", when)
	//fmt.Println("the dice was rolled at", when)

	// TODO: how to write the output into eyes.txt and dice.log?
	eyesFile, err := os.Create("eyes.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error, cannot create eyes.txt:", err)
		return
	}
	defer eyesFile.Close()

	logFile, err := os.Create("dice.log")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error, cannot create dice.log:", err)
		return
	}
	defer logFile.Close()

	fmt.Fprintln(eyesFile, "The dice shows", eyes, "eyes")
	fmt.Fprintln(logFile, "The dice shows", eyes, "eyes")
	fmt.Fprintln(logFile, "The dice was rolled at", when)

}
