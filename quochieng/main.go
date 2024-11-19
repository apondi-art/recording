package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"quochieng/quochieng/Calculate"
)

func main() {
	argument := os.Args[1:]
	if len(argument) != 2 {
		fmt.Println(`Usage  go run . credit/Debit Amount`)
		return
	}
	for {
		Balance := 1000
		data := Calculate.ReadFile()
		dataStr := strings.Fields(data)
		if len(dataStr) == 3 {
			remainBalance, _ := strconv.Atoi(dataStr[2])
			Balance = remainBalance
		}

		remainingAmount := 0
		Action := argument[0]
		Amount := argument[1]
		Action = strings.ToLower(Action)
		Name := ""
		Account := ""
		fmt.Println("Enter your Name : ")
		fmt.Scanf("%s", &Name)
		fmt.Println("Enter your Account Number")
		fmt.Scanf("%s", &Account)
		value, err := strconv.Atoi(Amount)
		if err != nil {
			fmt.Printf("%s", err)
			return
		}
		switch Action {
		case "debit":

			remainingAmount = Calculate.Debit(Balance, value)
		case "credit":
			if Balance < value {
				fmt.Println("The balance is less than Amount to be credited")
				return
			}
			remainingAmount = Calculate.Credit(Balance, value)
		}
		p := &Calculate.Status{
			Name:      Name,
			Account:   Account,
			Remaining: remainingAmount,
		}
		details := fmt.Sprintf("%s %s %v\n", p.Name, p.Account, p.Remaining)
		os.WriteFile("text.txt", []byte(details), 0o644)

		fmt.Printf("New balance : %v\n", remainingAmount)
		os.Exit(0)

	}
}
