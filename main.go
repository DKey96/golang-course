package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	// Go implies the Types of variables automatically from the value
	// fmt.Printf("The conferenceName is %T, the conferenceTickets is %T and remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to our %v booking application.\n", conferenceName)
	fmt.Printf("Currently there are still available %v tickets our of %v.\n", remainingTickets, conferenceTickets)
	fmt.Println("Book your tickets here to attend.")

	// Definition of the array with 50 elems and empty initial value (it is as in C so it's statically allocated)
	//var bookings = [50]string{}
	//var bookings [50]string

	// To create a dynamic array we use `slice` (it acts as list in python)
	//var bookings[]string
	bookings := []string{}


	// Define var without value
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

	//bookings[conferenceTickets - remainingTickets] = firstName + " " + lastName
	bookings = append(bookings, firstName + " " + lastName)
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for buying a %v tickets.\nConfirmation sent on %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("Remaining tickets: %v\n", remainingTickets)
	fmt.Printf("Already booked: %v", bookings)
}
