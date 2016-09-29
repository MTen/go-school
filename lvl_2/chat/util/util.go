package util

import "fmt"

func GetLocalNetworkIP() string {
	return "3030"
}

func RunGuest(stuff string) {
	fmt.Println(stuff)
}

func RunHost(one string, two string, three string) {
	fmt.Println(one + two + three)
}
