package main

import "fmt"

func main() {
	addToSlice()
}

func addToSlice() {
	langs := []string{"Go", "Ruby", "Python"}
	fmt.Println(langs)
}

func readSliceByIndexPosition() {
	langs := []string{"Go", "Ruby", "Python"}
	fmt.Println(langs[0])
}

func readSliceWithForLoopAndRange() {
	langs := []string{"Go", "Ruby", "Python"}
	for i := range langs {
		fmt.Println(langs[i])
	}

	// range can return two values, index and value of index
	// since we are not using i - we must assign index to something
	for _, element := range langs {
		fmt.Println(element)
	}

}
