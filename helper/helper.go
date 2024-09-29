package helper

import (
	"fmt"
	"strings"
)

func Firstnamedisplay(attendees []string) []string {
	DisplayNames := []string{}
	for _, attendee := range attendees {
		var name = strings.Fields(attendee)
		//var DisplayName = name[0]
		DisplayNames = append(DisplayNames, name[0])
	}
	return DisplayNames
}

func Firstpage(confname string, tickets int, availableTickets int) {
	fmt.Println("Welcome to", confname)
	fmt.Println("we have", tickets, "and available tickets are", availableTickets)

}
