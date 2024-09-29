package main

import (
	"Bookin-app/helper"
	"fmt"
	"strings"
)

func main() {
	for {

		var conferenceName = "Go confrence"
		const Tickets = 50
		var availableTickets = 50
		//var firstName string
		//var secondName string
		//var email string
		//var requiredTicket int
		var attendees []string

		//fmt.Println("Welcome to", conferenceName)
		//fmt.Println("we have", Tickets, "and available tickets are", availableTickets)
		helper.Firstpage(conferenceName, Tickets, availableTickets)

		firstName, secondName, email, requiredTicket := userinput()

		//fmt.Print("enter your name:\n")
		//fmt.Scan(&firstName)

		//fmt.Print("enter your Second name:\n")
		//fmt.Scan(&secondName)

		//fmt.Print("enter your email:\n")
		//fmt.Scan(&email)

		//fmt.Print("enter number of tickets required: \n")
		//fmt.Scan(&requiredTicket)

		isvalidName := len(firstName) >= 2 && len(secondName) >= 2
		isvalidemail := strings.Contains(email, "@")
		isvalidticket := requiredTicket > 0 && requiredTicket <= availableTickets

		if isvalidName && isvalidemail && isvalidticket {

			//if requiredTicket <= availableTickets {

			availableTickets = availableTickets - requiredTicket

			fmt.Printf("%v booked %v number of tickets \n", firstName, requiredTicket)
			fmt.Printf("%v tickets are available \n", availableTickets)

			attendees = append(attendees, firstName+" "+secondName)
			fmt.Printf("%v \n", attendees)

			//DisplayNames := []string{}
			//for _, attendee := range attendees {
			//	var name = strings.Fields(attendee)
			//	var DisplayName = name[0]
			//	DisplayNames = append(DisplayNames, DisplayName)

			//}
			fnamee := helper.Firstnamedisplay(attendees)
			fmt.Printf("attendes fnames are %v \n", fnamee)

		} else {
			fmt.Printf("available tickets are only %v \n", availableTickets)

		}
	}
}

//func firstpage(confname string, tickets int, availableTickets int) {
//	fmt.Println("Welcome to", confname)
//	fmt.Println("we have", tickets, "and available tickets are", availableTickets)

//}

//func firstnamedisplay(attendees []string) []string {
//	DisplayNames := []string{}
//	for _, attendee := range attendees {
//		var name = strings.Fields(attendee)
//		//var DisplayName = name[0]
//		DisplayNames = append(DisplayNames, name[0])
//	}
//	return DisplayNames
//}

func userinput() (string, string, string, int) {
	var firstName string
	var secondName string
	var email string
	var requiredTicket int

	fmt.Print("enter your name:\n")
	fmt.Scan(&firstName)

	fmt.Print("enter your Second name:\n")
	fmt.Scan(&secondName)

	fmt.Print("enter your email:\n")
	fmt.Scan(&email)

	fmt.Print("enter number of tickets required: \n")
	fmt.Scan(&requiredTicket)

	return firstName, secondName, email, requiredTicket

}
