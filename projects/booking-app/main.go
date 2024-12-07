package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

const conferenceTickets uint = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]User, 0)

type User struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	helper.GreetUser(conferenceName, conferenceTickets, remainingTickets)

	for {
		firstName, lastName, email, userTickets := helper.GetUserInput()
		isValidName, isValidEmail, isValidUserTickes := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidUserTickes {
			bookTickets(userTickets, firstName, lastName, email)

			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are : %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short.")
			}

			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign.")
			}

			if !isValidUserTickes {
				fmt.Println("number of you entered is invalid.")
				fmt.Printf("Remaining tickets : %v, but you entered %v\n", remainingTickets, userTickets)
			}

			fmt.Println("Your input data is invalid, please try again.")
			continue
		}
	}
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTickes := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidUserTickes

}

func getFirstNames(bookings []User) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames

}

func bookTickets(userTickets uint, firstName string, lastName string, email string) uint {
	// book ticket in system
	remainingTickets = remainingTickets - userTickets
	// create user struct
	var user = User{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, user)

	fmt.Printf("Thank you  %v %v for booking %v tickets, you will recevie an email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	return remainingTickets
}
