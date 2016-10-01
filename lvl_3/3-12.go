// get a list of blocked IPs with getBlockedIPs and then use a for loop to
// print the blockedIPs to the console.

package main

import (
	"fmt"
)

func main() {
	example1()
}

func loopMethodOne() {
	list := getBlockedIPs()
	for i := range list {
		fmt.Println(list[i])
	}
}

func loopMethodTwo() {
	list := getBlockedIPs()
	for _, element := range list {
		fmt.Println(element)
	}
}

func getBlockedIPs() []string {
	return []string{"192.168.0.17", "192.168.0.18", "192.168.0.19", "192.168.0.20"}
}
