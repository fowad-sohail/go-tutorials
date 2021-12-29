package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name cannot be empty")
	}
	// Return a greeting that embeds the name in the message
	message := fmt.Sprintf(randomFormat(), name)

	// To break the test:
	// message := fmt.Sprint(randomFormat())

	return message, nil
}

// Returns a map that associates named people with greeting message
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you %v",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
