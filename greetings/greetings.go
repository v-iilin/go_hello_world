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

// Hellos returns a map that associates each of the named people
// with a greeting message
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// in the map, associate the retieved message with the name
		messages[name] = message
	}
	return messages, nil
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