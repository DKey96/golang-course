package main

import (
	"booking-app/shared"
	"fmt"
	"strconv"
)

// Defining variables here, they are package wide (global)
var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50
// The number as second input parameter is the initial size of the slice/list. It can be 0 as the slice is dynamic
var bookings = make([]map[string]string, 0) // To create a dynamic array we use `slice` (it acts as list in python)

func main() {
	greetUser()

	// GoLang has only For loops
	// Infinite loop
	for {
		firstName, lastName, email, userTickets := getUserInputs()
		isValidName, isValidEmail, isValidTicketsCount := shared.ValidateUserInputs(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketsCount {
			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
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
			if !isValidName {
				fmt.Println("Please enter a valid first name and last name. (They must be longer than 2 characters).")
			}
			if !isValidEmail {
				fmt.Println("Please enter a valid email address. (should contain @)")
			}

			if !isValidTicketsCount {
				fmt.Printf("Please enter a valid number of tickets. Remaining: %v\n", remainingTickets)
			}
		}
	}
}

func greetUser() {
	fmt.Printf("Welcome to our %v booking application.\n", conferenceName)
	fmt.Printf("Currently there are still available %v tickets our of %v.\n", remainingTickets, conferenceTickets)
	fmt.Println("Book your tickets here to attend.")
}

// Functions with return values must have typed output
func getFirstNames() []string {
	firstNames := []string{}
	for _, bookings := range bookings {
		firstNames = append(firstNames, bookings["firstName"])
	}
	return firstNames
}

func getUserInputs() (string, string, string, uint) {
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

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// Create a map for user data
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	// GoLang cannot have a map (dict) with mixed datatypes as values
	// This function formats the Uint into float so we have to add also the number of decimal numbers
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for buying a %v tickets.\nConfirmation sent on %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("Remaining tickets: %v\n", remainingTickets)
}
