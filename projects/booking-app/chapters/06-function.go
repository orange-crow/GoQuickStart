package main

import (
	"fmt"
	"strings"
)

const conferenceTickets uint = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings []string //在中括号内指定数字后就表明这个数组是固定size，反之则是可以动态扩展的数组.

func main() {
	greetUser()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidUserTickes := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidUserTickes {
			bookTickets(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
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
			}

			fmt.Println("Your input data is invalid, please try again.")
		}
	}
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets, and %v are still avaliabel.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets to attend.\n")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames

}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTickes := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidUserTickes

}

func getUserInput() (string, string, string, uint) {
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

func bookTickets(userTickets uint, firstName string, lastName string, email string) uint {
	// book ticket in system
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you  %v %v for booking %v tickets, you will recevie an email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	return remainingTickets
}
