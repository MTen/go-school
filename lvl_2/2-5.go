package main

import (
	"fmt"
	"time"
)

func main() {
	hourOfDay := time.Now().Hour()
	greeting := getGreeting(hourOfDay)
	fmt.Println(greeting)
}

func getGreeting(hour int) string {
	if hour < 12 {
		return "Good Morning"
	} else if hour < 16 {
		return "Good Afternoon"
	} else {
		return "Good Evening"
	}
}
