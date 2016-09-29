package main

import (
	"fmt"
	"os"
)

func main() {
	staticType()
	dynamicType()
}

func staticType() {
	// static variables need var {{variable name}} {{data type}}
	var args []string

	args = os.Args
	if len(args) > 1 {
		fmt.Println(args[1], "static")
	} else {
		fmt.Println("Hi, I'm static assignment")
	}
}
func dynamicType() {
	// dynamic assignment is prefered in go but can lead to bugs in code
	dynamicArg := os.Args
	if len(dynamicArg) > 1 {
		fmt.Println(dynamicArg[1], "I'm dynamic assignment")
	} else {
		fmt.Println("Hi I'm dynamic")
	}
}
