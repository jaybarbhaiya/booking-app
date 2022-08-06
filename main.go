package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	conferenceTickets := 50
	remainingTickets := conferenceTickets

	fmt.Printf("Welcome to the %v ticket booking system \n", conferenceName)
	fmt.Printf("Total ticket: %v and Remaining tickets: %v\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend the conference!")
}
