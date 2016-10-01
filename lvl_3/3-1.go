// make a for loop that prints 1 - 5

package main

import "fmt"

func main() {
	fmt.Println("For Loop")
	forLoop()
	fmt.Println("While Loop")
	whileLoop()
	fmt.Println("Break Loop")
	breakLoop()
}

func forLoop() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}

func whileLoop() {
	i := 1
	isLessThanFive := true
	for isLessThanFive {
		if i >= 5 {
			isLessThanFive = false
		}
		fmt.Println(i)
		i++
	}
}

func breakLoop() {
	i := 1
	for {
		if i > 5 {
			break
		}
		fmt.Println(i)
		i++
	}
}

func infiniteLoop() {
	// good for network listening processes

	for {
		//someListeningFunctionHere()
	}
}
