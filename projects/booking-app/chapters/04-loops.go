package main

import (
	"fmt"
	"strings"
)

func main() {
	// var conferenceName = "Go Conference"
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string //在中括号内指定数字后就表明这个数组是固定size，反之则是可以动态扩展的数组.
	fmt.Printf("Welcode to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets, and %v are still avaliabel.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets to attend.\n")

	for {
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
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("The whole slice: %v\n", bookings)
		fmt.Printf("The first value: %v\n", bookings[0])
		fmt.Printf("Slice Type: %T\n", bookings)
		fmt.Printf("Slice length: %v\n", len(bookings))

		fmt.Printf("Thank you  %v %v for booking %v tickets, you will recevie an email at %v \n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("There are all our bookings: %v\n", firstNames)

	}
}
