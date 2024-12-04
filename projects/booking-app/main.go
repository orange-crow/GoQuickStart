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
}
