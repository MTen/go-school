// A new function named authorizeAdmin() has been added to our chat program. This
// function checks whether a new guest connection should receive administrator
// rights. If so, then the function assigns true to the isAdmin property of the
// struct passed as an argument. The new function is only partially implemented.
// Let’s finish it!

// Task 1
// Create a new guestConnection struct. Add the special operator that returns a
// pointer to the new struct.

// Task 2
// Complete the signature of the authorizeAdmin() function so that it accepts a
// pointer to the guestConnection struct as its single argument. Note: It
// should define the variable c.

package main

import (
	"fmt"

	"./chat/util"
)

type guestConnection struct {
	ip       string
	userName string
	isAdmin  bool
}

func (g guestConnection) isAllowed() bool {
	return !isIPBlocked(g.ip) && g.userName != "Darth Vader"
}

func main() {
	ip := util.GetGuestIP()
	userName := "Obi-Wan"

	gConn := &guestConnection{ip: ip, userName: userName}
	fmt.Println("Before auth", gConn)
	authorizeAdmin(gConn)
	fmt.Println("After auth", gConn)
}

func authorizeAdmin(c *guestConnection) {
	if c.isAllowed() && c.ip == "192.168.0.10" {
		c.isAdmin = true
	}
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
