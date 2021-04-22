package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// oneHundred returns the number one
// hundred without using any numerals.
func oneHundred() int {
	var one int
	one++
	oneHundred, _ := strconv.Atoi(
		fmt.Sprintf(
			"%b",
			one<<one<<one,
		),
	)
	return oneHundred
}

// main loops from one to one hundred
// incrementing a counter and printing
// the new count along the way.
func main() {
	log.SetOutput(os.Stdout)
	var counter int
	for counter < oneHundred() {
		counter++
		log.Printf("%d", counter)
	}
}
