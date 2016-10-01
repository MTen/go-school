// Task 1
// There’s a way to rewrite the previous solution using a for loop with no
// component at all. Complete the code with the missing keyword so that the
// for loop stops once util.PingChatServer() returns true.

package main

import (
	"fmt"

	"./pinger/util"
)

func main() {
	for {
		if util.PingChatServer() {
			break
		}
	}
	fmt.Println("Ping is done")
}
