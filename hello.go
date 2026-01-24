package main

import (
	"fmt"
	"strconv"
	"strings"
)

const conferenceTickets = 50

var motto = "Go conference"
var remainingTickets uint = 50

func main() {
	var bookings = make([]map[string]string, 0)

	greetUsers(motto, conferenceTickets, remainingTickets)

	for {
		// get user input
		firstname, lastname, mail, userTickets := getUserInput()

		// validation of user input
		isValidEmail, isValidName, isValidTicket := valUserInput(firstname, lastname, mail, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicket {
			fmt.Printf("Thank you %v %v for booking %v tickets with us. You will receive a confirmation email at %v\n", firstname, lastname, userTickets, mail)

			// booking the conference
			bookings = bookingProcess(remainingTickets, userTickets, firstname, lastname, mail, bookings, motto)

			// call func that prints firstnames
			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			fmt.Println()

			//	noTickets := remainingTickets == 0 // you can set your condition as a variable so  you just set it inside the code, incase you would be using it several times.
			if remainingTickets == 0 {
				fmt.Printf("%v is booked out. Come back next year", motto)
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your firstname or lastname is too short")
			}
			if !isValidEmail {
				fmt.Println("Invlaid E-mail Address")
			}
			if !isValidTicket {
				fmt.Println("Invalid ticket number")
			}

		}

	}

}

func greetUsers(motto string, conferenceTickets uint, remainingTickets uint) {
	fmt.Printf("Welcome to %v\n", motto)
	fmt.Printf("We have a total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
}

func valUserInput(firstname string, lastname string, mail string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstname) >= 2 && len(lastname) > 2
	isValidEmail := strings.Contains(mail, "@")
	isValidTicket := userTickets > 0 && userTickets <= remainingTickets
	return isValidEmail, isValidName, isValidTicket
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings { // that is an underscore or a blank identifier, it tells GO that there is a variable there and we want to ignnore it.
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstname string
	var userTickets uint
	var lastname string
	var mail string

	fmt.Println("Kindly enter your firstname: ")
	fmt.Scan(&firstname)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastname)

	fmt.Println("Enter your e-mail address: ")
	fmt.Scan(&mail)

	fmt.Println("How many tickets would you like to book?")
	fmt.Scan(&userTickets)

	return firstname, lastname, mail, userTickets
}

func bookingProcess(remainingTickets uint, userTickets uint, firstname string, lastname string, mail string, bookings []string,  motto string) []string,0 {
	remainingTickets = remainingTickets - userTickets

	// creating a map for a user
	var userData = make(map[string]string, 0)
	userData["firstName"] = firstname
	userData["lastName"] = lastname
	userData["E-mail"] = mail
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, firstname+" "+lastname) // adding the values to the slice
	// NB: extracting values from a slice is the same way as an array ("varname[index]")
	fmt.Printf("\n There are %v tickets remaining for %v\n Grab yours now\n", remainingTickets, motto)
	fmt.Printf("These are all our bookings %v\n", bookings)
	return bookings
}
