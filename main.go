package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{} // To create a dynamic array we use `slice` (it acts as list in python)

	greetUser(conferenceName, conferenceTickets, remainingTickets)

	// GoLang has only For loops
	// Infinite loop
	for {
		firstName, lastName, email, userTickets :=	getUserInputs()
		isValidName, isValidEmail, isValidTicketsCount := validateUserInputs(firstName, lastName, email, userTickets, remainingTickets)
		
		if isValidName && isValidEmail && isValidTicketsCount {
			bookings := bookTicket(remainingTickets, userTickets, firstName, lastName, email, bookings)
			firstNames := getFirstNames(bookings)
			fmt.Printf("First names of bookings are: %v\n", firstNames)

			// noTicketsRemaining := remainingTickets == 0
			// if noRemainingTickets {
			if remainingTickets == 0 {
				fmt.Println("The conference is booked out. Please come next year.")
				break
			} else {
				continue
			}
		} else {
			if !isValidName{
				fmt.Println("Please enter a valid first name and last name. (They must be longer than 2 characters).")
			}
			if !isValidEmail{
				fmt.Println("Please enter a valid email address. (should contain @)")
			}
			
			if !isValidTicketsCount{ 
				fmt.Printf("Please enter a valid number of tickets. Remaining: %v\n", remainingTickets)
			}
		}
	}
}

func greetUser(conferenceName string, remainingTickets int, conferenceTickets uint){
	fmt.Printf("Welcome to our %v booking application.\n", conferenceName)
	fmt.Printf("Currently there are still available %v tickets our of %v.\n", remainingTickets, conferenceTickets)
	fmt.Println("Book your tickets here to attend.")
}

// Functions with return values must have typed output
func getFirstNames(bookings []string) []string{
	firstNames := []string{}
	for _, name := range bookings{
		var names = strings.Fields(name)[0]
		firstNames = append(firstNames, names)
	}
	return firstNames
}

// GoLang can also return multiple values as python. All of them must be typed. Types must be enclodes in the braces
func validateUserInputs(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketsCount := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketsCount
}

func getUserInputs() (string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Bear in mind that there has to be a pointer! 
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)	
	fmt.Println("Enter your email address")
	fmt.Scan(&email)	
	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)	
	
	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, firstName string, lastName string, email string, bookings []string) []string{
	bookings = append(bookings, firstName + " " + lastName)
	remainingTickets = remainingTickets - userTickets
		
	fmt.Printf("Thank you %v %v for buying a %v tickets.\nConfirmation sent on %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("Remaining tickets: %v\n", remainingTickets)
	return bookings
}