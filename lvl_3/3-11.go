// Use a slice literal to create a slice of type []string with the initial values
// "192.168.0.19" and "192.168.0.20". Then, assign it to the blockedIPs variable
// using type inference with :=.

package main

import (
	"./chat/util"
)

func main() {
	blockedIPs := []string{"192.168.0.19", "192.168.0.20"}
	addToBlockedList(blockedIPs)
}

func addToBlockedList(ips []string) {
	util.SaveBlockedIPs(ips)
}
