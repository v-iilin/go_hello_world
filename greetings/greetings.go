package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// Return a greeting that embeds the name in a message.
	
	// if no name was given, return an error witha message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// if a name was received, return a value that embeds the name

	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	
	return message, nil
}