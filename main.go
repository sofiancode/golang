package main

import (
	"fmt"
	"strings"
)

const ConferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	// greeting users func
	greetUsers(conferenceName, ConferenceTickets, remainingTickets)

	for {
		firstName, lastName, email, userTickets := getUserInput()
		// validation
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)
			// call function print first names
			firstNames := getFirstNames(bookings)
			fmt.Printf("The fist names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end the program
				fmt.Println("Our conference is booked out! Please come next year!")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("first name of last name is too short")
			}

			if !isValidEmail {
				fmt.Println("your email address does not contain @ sign")
			}

			if !isValidTicketNumber {
				fmt.Println("your number of tickets is invalid")
			}

		}

	}

}

func greetUsers(conferenceName string, ConferenceTickets int, remainingTickets uint) {
	fmt.Printf("welcome to %v booking aplication\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available.\n", ConferenceTickets, remainingTickets)
	fmt.Printf("Get your ticket here to attend\n")
}

func getFirstNames(bookings []string) []string {

	firstNames := []string{}
	for _, booking := range bookings { // (underscore) blank identifier -> to   ignore a variable u don't want to use
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("enter your email adress: ")
	fmt.Scan(&email)
	fmt.Println("enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank u %v %v  for booking %v tickets, u will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
