package main

import "fmt"

func main() {
	// var conferenceName = "Go Conference"
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcode to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets, and %v are still avaliabel.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets to attend.\n")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// asking your info
	fmt.Println("Enter your first name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter your lastr name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter your number of tickets: ")
	fmt.Scanln(&userTickets)

	// book ticket in system
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you  %v %v for booking %v tickets, you will recevie an email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
