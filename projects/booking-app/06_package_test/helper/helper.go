package helper

import (
	"fmt"
	"strings"
)

func GreetUser(conferenceName string, conferenceTickets uint, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets, and %v are still avaliabel.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets to attend.\n")
}

func GetFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames

}

func GetUserInput() (string, string, string, uint) {
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
	return firstName, lastName, email, userTickets

}
