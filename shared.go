package main

import "strings"

func ValUserInput(firstname string, lastname string, mail string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstname) >= 2 && len(lastname) > 2
	isValidEmail := strings.Contains(mail, "@")
	isValidTicket := userTickets > 0 && userTickets <= remainingTickets
	return isValidEmail, isValidName, isValidTicket
}
