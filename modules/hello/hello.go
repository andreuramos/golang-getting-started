package main

import (
	"fmt"
	"log"

	"example/greetings"
)

func main() {
	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Gladys")

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(message)
}