package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets

	fmt.Printf("Welcome to the %v ticket booking system \n", conferenceName)
	fmt.Printf("Total ticket: %v and Remaining tickets: %v\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend the conference!")

	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Printf("Enter first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter email: ")
	fmt.Scan(&email)

	fmt.Printf("Number of ticket: ")
	fmt.Scan(&userTicket)
	remainingTickets -= userTicket

	fmt.Printf("Thank you %v %v for bookin %v tickets. You will recevice and email confirmation at %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for the %v", remainingTickets, conferenceName)
}
