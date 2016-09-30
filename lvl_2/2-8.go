package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {
	time := time.Now().Hour()
	greeting, err := getGreeting(time)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	println(greeting)
}

func getGreeting(hour int) (string, error) {
	var message string

	if hour < 24 {
		err := errors.New("Too early for greetings!")
		return message, err
	} else if hour < 24 {
		message = "Hello there!"
	}
	return message, nil
}
