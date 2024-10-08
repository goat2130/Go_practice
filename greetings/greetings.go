package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	// If no name was given, return an error message
	if name == "" {
		return "", errors.New("empty name")
	}

	// Create a message using a random format.
	message := fmt.Sprint(randomFormat())

	// message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people with a greeting message.
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages
	messages := make(map[string]string)

	// Loop through the received slice of names, callig
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// In the map, associate the retrieved message with the name.
		messages[name] = message
	}
	return messages, nil
}

// random Format a returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}

	/**
	* Return a randomly selected message format by specifying
	* random index for the slece of formats
	* The code rand.Intn(len(formats)) generates a random integer using
	* the length of the formats slice as the upper limit of the range, and returns that integer.
	 */
	return formats[rand.Intn(len(formats))]
}
