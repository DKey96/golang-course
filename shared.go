package main

import "strings"

// GoLang can also return multiple values as python. All of them must be typed. Types must be enclodes in the braces
func validateUserInputs(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketsCount := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketsCount
} 