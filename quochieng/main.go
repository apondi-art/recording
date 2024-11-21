package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"quochieng/quochieng/Calculate"
)

func main() {
	Calculate.Read()       // Load data from JSON into the Record map
	defer Calculate.Save() // Ensure data is saved back to JSON on program exit

	// Validate command-line arguments
	argument := os.Args[1:]
	if len(argument) != 2 {
		fmt.Println(`Usage: go run . credit/Debit Amount`)
		return
	}

	Action := strings.ToLower(argument[0]) // Extract and normalize action
	Amount := argument[1]

	// Convert Amount to an integer
	value, err := strconv.Atoi(Amount)
	if err != nil {
		fmt.Printf("Invalid amount: %s\n", err)
		return
	}

	// Loop to process transactions
	for {
		var Name, Account string
		fmt.Print("Enter your Name: ")
		fmt.Scanf("%s", &Name)
		fmt.Print("Enter your Account Number: ")
		fmt.Scanf("%s", &Account)

		// Generate the map key
		key := fmt.Sprintf("%s:%s", Name, Account)
		fmt.Println("Processing Key:", key)

		// Check if the user exists in the record
		var user *Calculate.Status
		if _, exist := Calculate.Record[key]; exist {
			user = Calculate.Record[key] // Use the existing user
		} else {
			// Create a new user with default balance if not found
			user = &Calculate.Status{
				Name:    Name,
				Account: Account,
				Balance: 1000,
				
			}
			Calculate.Record[key] = user
		}

		// Perform the action
		switch Action {
		case "debit":
			if user.Balance < value {
				fmt.Println("The balance is less than the amount to be debited")
				continue // Prompt user again without exiting
			}
			user.Debit(value)
		case "credit":
			user.Credit(value)
		default:
			fmt.Println("Invalid action. Only 'credit' or 'debit' are allowed.")
			continue // Prompt user again without exiting
		}

		// Update the Record map

		// Print the updated balance and record
		fmt.Printf("Updated balance: %v\n", user.Balance)

		break // Exit the loop after processing one transaction
	}
}
