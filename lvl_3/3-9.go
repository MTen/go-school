// In order to add more flexibility to the blockedIPs array, we've decided to change
// its type. Let's finish the manual variable declaration below so that blockedIPs
// is defined as a slice of string types, and not an array.

package main

import (
	"./chat/util"
)

func main() {
	var blockedIPs []string

	blockedIPs = append(blockedIPs, "192.168.0.16")
	blockedIPs = append(blockedIPs, "192.168.0.17")
	blockedIPs = append(blockedIPs, "192.168.0.18")
	blockedIPs = append(blockedIPs, "192.168.0.19")
	blockedIPs = append(blockedIPs, "192.168.0.20")

	addToBlockedList(blockedIPs)
}

func addToBlockedList(ips []string) {
	util.SaveBlockedIPs(ips)
}
