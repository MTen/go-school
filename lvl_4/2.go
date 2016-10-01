// For auditing purposes, we’ll need to track our guest's IP addresses and their
// user name. We are currently just using variables to store these values, but
// in order to keep our code organized and encapsulated, we want to use a struct.

// Task 1
// Define a new struct named guestConnection.

// Task 2
// Add two properties to the new struct: one named ip and one named userName.
// Make both have type string.

// Task 3
// From inside the main() function, create a new guestConnection struct and
// assign it to a variable named gConn. Then assign the values from the ip and
// userName variables to the corresponding properties on the struct.

package main

import (
	"fmt"

	"./chat/util"
)

type guestConnection struct {
	ip       string
	userName string
}

func main() {
	ip := util.GetGuestIP()
	userName := "Kerry"
	gConn := guestConnection{ip: ip, userName: userName}

	fmt.Println(gConn)
}
