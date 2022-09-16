package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTicket uint = 50
var bookings = make([]userData, 1)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
	isOptInForNews  bool
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	//for {
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName,
		lastName, email, userTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTickets(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstnames()
		fmt.Printf("The first names of the booking are %v\n", firstNames)

		if remainingTicket == 0 {
			fmt.Println("Sorry, there is no available ticket ")
			//break
		}

	} else {
		if !isValidName {
			fmt.Println("Firstname or lastname is incorrect or too short")
		}
		if !isValidEmail {
			fmt.Println("Enter a valid email address")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}
	}

	//}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking app\n", conferenceName)

	fmt.Printf("we have the total of %v and %v is still available\n",
		conferenceTickets, remainingTicket)
	fmt.Println("Get your tickets here")
}

func getFirstnames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("How many ticket, do you want:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTicket = remainingTicket - userTickets

	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("list of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for buying %v tickets. You will receive a confirmation email at %v\n",
		firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets left for booking\n", remainingTicket)

}

func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTicket, firstName, lastName)
	fmt.Println("################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}
