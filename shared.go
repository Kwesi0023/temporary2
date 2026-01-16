package main

import "strings"

func ValUserInput(firstname string, lastname string, mail string, userTickets uint, remainingTickets uint) (bool, bool, bool) { // check if both names are at least two characters long
	isValidName := len(firstname) >= 2 && len(lastname) > 2
	isValidEmail := strings.Contains(mail, "@")
	isValidTicket := userTickets > 0 && userTickets <= remainingTickets
	// MAtch the order : Name, Email, Ticket
	return isValidEmail, isValidName, isValidTicket
}
