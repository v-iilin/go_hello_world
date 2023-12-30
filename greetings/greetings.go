package greetings

import (
	"errors"
	"fmt"
	"math/rand"
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

	message := fmt.Sprintf(randomFormat(), name)
	
	return message, nil
}

func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}	
	// return a randomly selected message format by specifying 
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}