package main

import (
	"context"
	"fmt"

	"github.com/addhyanmalhotra/GoChat/client"
	"github.com/addhyanmalhotra/GoChat/server"
)

func main() {

	// Variable to store Choice
	var choice int

	// Temporary Variables to store user input
	var password, address, clientName string

	// instantiate context
	ctx, cancel := context.WithCancel(context.Background())

	// Display User Menu
	fmt.Println(
		"Simple Chat Application\n",
		" 1. Create a server\n",
		" 2. Create a Client")

	// Input promt
	fmt.Print(">>")

	// Get user input
	fmt.Scanf("%d\n", &choice)

	if choice == 1 {

		// Cancel on Termination
		defer cancel()

		// Input socket 'domanin:port' as a string
		fmt.Printf("Enter the Server Port      : ")
		fmt.Scanf("%s\n", &address)

		// Input password for the server
		fmt.Printf("Enter the Server Password  : ")
		fmt.Scanf("%s\n", &password)

		// Create Server using constructor
		s := server.Server(password, address)

		// Chan to for IPC
		done := make(chan bool)

		// Run the new server
		go s.Run(ctx, done)

		// Wait till server terminates execution
		select {
		case <-done:
		}

	} else if choice == 2 {

		fmt.Printf("Enter the Client username (leave empty for default username) : ")
		fmt.Scanf("%s\n", &clientName)

		// Input socket 'domanin:port' as a string
		fmt.Printf("Enter the Server Port      : ")
		fmt.Scanf("%s\n", &address)

		// Input password for the server
		fmt.Printf("Enter the Server Password  : ")
		fmt.Scanf("%s\n", &password)

		term := make(chan bool)

		// use Default constructor to create a new client
		newClient := client.Client(password, address, clientName)

		// Run as a go process
		go newClient.Run(ctx, term)
		select {
		case <-ctx.Done():
			break
		case <-term:
			break
		}
	}
}
