package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

// a struct is a like a class in java but more lightweight
// and does not support inheritance:
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)
		// we are going to make the following break
		// out of the main thread (concurrency)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("First name or Last name you entered is too short.")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered does not contain '@' sign")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets you entered is invalid")
		}
	}
	// this waits until the WaitGroup counter is 0
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
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
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
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

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// we are adding a delay of 10 seconds:
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("############################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("############################")
	wg.Done()
}
