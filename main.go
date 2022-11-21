package main

import (
	"fmt"
	"golang/helper"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	// greeting users func
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		// validation
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			// call function print first names
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

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

func greetUsers() {
	fmt.Printf("welcome to %v booking aplication\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your ticket here to attend\n")
}

func getFirstNames() []string {

	firstNames := []string{}
	for _, booking := range bookings { // (underscore) blank identifier -> to   ignore a variable u don't want to use
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
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

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	// create a map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank u %v %v  for booking %v tickets, u will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
