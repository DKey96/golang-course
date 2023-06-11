package shared

import "strings"

// To Export the function from a package, use Capital First Letter of the name of the Function
// This also works for variables
var MyGlobalVar = "myglobalvalue"

// GoLang can also return multiple values as python. All of them must be typed. Types must be enclodes in the braces
func ValidateUserInputs(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketsCount := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketsCount
}