// create a  program that stores 3 programming languages in an array
// and then prints them back to the console.

package main

import (
	"fmt"
)

func main() {
	staticArrayAssignment()
	appendArrayAssignment()
}

func staticArrayAssignment() {
	var langs [3]string
	langs[0] = "Go"
	langs[1] = "Ruby"
	langs[2] = "Javascript"

	fmt.Println(langs)
}

func appendArrayAssignment() {
	var langs []string
	langs = append(langs, "Go", "LOLcode", "Rawr")

	fmt.Println(langs)
}
