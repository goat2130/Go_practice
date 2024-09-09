package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	// a slice of names.
	names := []string{"Glady", "Samantha", "Darrin"}

	// If an error was returned, print it to the console and
	// exit the program.
	// message, err := greetings.Hello("Gladys")
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// message := greetings.Hello("Gladys")
	fmt.Println(messages)
}
