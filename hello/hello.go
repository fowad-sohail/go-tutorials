package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Logger setup
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Fowad", "Shahid", "Sohail"}

	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
