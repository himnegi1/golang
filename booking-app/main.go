package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const totalConferenceTickets = 50 
	var remainingConferenceTickets = 50
	fmt.Printf("welcome to %v booking app\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are avaialble\n", totalConferenceTickets, remainingConferenceTickets)
	fmt.Println("Get your ticket to attend")

	var userName string
	var useTickets int


	userName = "Himanshu"
	useTickets = 30
	fmt.Printf("User: %v used %v tickets\n", userName, useTickets)

}