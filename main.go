package main

import (
	"fmt"
	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
)

func main() {
	message := "Hello, Cowsay in Go!"

	// Create a new cow instance with default settings
	cow, err := cowsay.New()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Generate the Cowsay message
	say, err := cow.Say(message)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(say)
}

