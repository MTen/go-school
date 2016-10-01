// Let’s finish implementing the isIPBlocked() function.
// This function takes the host IP address as its argument and loops
// through a slice of blocked IPs to compare the argument and determine
// whether the host IP should be blocked from accessing our chat app.

// Task 1
// Use range to iterate over the elements from blockedIPs and assign them to a
// variable called blockedIP. Remember, to read elements from a slice, we must
// assign to two values and one of them must be ignored using a special character.

// Task 2
// Inside the for loop, check if the current blockedIP is equal to the IP
// address sent as an argument to the function. If so, then return true.

package main

import (
	"fmt"

	"./chat/util"
)

func main() {
	hostIP := util.GetHostIP()
	fmt.Println(isIPBlocked(hostIP))
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
