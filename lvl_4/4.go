// A couple conditions need to be met in order to allow guest connections to our
// chat app. We’ve coded these conditions as part of the isAllowed() method.
// However, there are a couple of things still missing from the method signature.

// Task 1
// Add the proper receiver to the isAllowed() method. Note: It should define
// the variable g.

// Task 2
// Complete the isAllowed() method signature by indicating it returns a type bool.

// Task 3
// Invoke the isAllowed() method on gConn and assign the return value to a
// variable called isAllowedStatus using type inference with :=.

package main

import (
	"fmt"

	"./chat/util"
)

type guestConnection struct {
	ip       string
	userName string
}

func (g guestConnection) isAllowed() bool {
	return !isIPBlocked(g.ip) && g.userName != "Darth Vader"
}

func main() {
	ip := util.GetGuestIP()
	userName := "Kerry"

	gConn := guestConnection{ip: ip, userName: userName}
	isAllowedStatus := gConn.isAllowed()
	fmt.Println(isAllowedStatus)
}

func isIPBlocked(ip string) bool {
	blockedIPs := getBlockedIPs()
	for _, blockedIP := range blockedIPs {
		if blockedIP == ip {
			return true
		}
	}
	return false
}

func getBlockedIPs() []string {
	return []string{"192.168.0.17", "192.168.0.18", "192.168.0.19", "192.168.0.20"}
}
