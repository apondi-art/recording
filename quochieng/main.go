package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"quochieng/Calculate"
)

type Status struct {
	Name      string
	Account   string
	Remaining int
}

func main() {
	argument := os.Args[1:]
	if len(argument) != 2 {
		fmt.Println(`Usage  go run . credit/Debit Amount`)
		return
	}
	for {
		Balance := 1000
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
		p := &Status{
			Name:      Name,
			Account:   Account,
			Remaining: remainingAmount,
		}
		fmt.Println(*p)

		fmt.Printf("New balance : %v\n", remainingAmount)

	}
}
