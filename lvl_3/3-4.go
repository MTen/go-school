package main

// Task 1
// Create a for loop that runs while the isAlive variable is false. It helps to
// know that in order to negate a boolean variable in Go, we can prepend it with
// the exclamation point. For example, !true returns false and !false returns true.

// Task 2
// Inside the for loop, invoke the util.PingChatServer() function to ping the chat
// server. This function returns a boolean and when this boolean is true, set
// the isAlive variable to true. This will cause the loop to break.

import (
	"fmt"

	"./pinger/util"
)

func main() {
	isAlive := false
	for !isAlive {
		if util.PingChatServer() {
			isAlive = true
		}
	}

	fmt.Println("Ping is done")
}
