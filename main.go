package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const ConferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking aplication\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available.\n", ConferenceTickets, remainingTickets)
	fmt.Printf("Get your ticket here to attend\n")

	for {

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

		// validation

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank u %v %v  for booking %v tickets, u will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}

			for _, booking := range bookings { // (underscore) blank identifier -> to ignore a variable u don't want to use
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The fist names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end the program
				fmt.Println("Our conference is booked out! Please come next year!")
				break
			}

		} else {
			fmt.Println("Your input data is invalid! Try again!")

		}

	}

}
