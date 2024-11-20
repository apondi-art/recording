package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"quochieng/quochieng/Calculate"
)

func main() {
	Calculate.Read()
	defer Calculate.Save()
	user := &Calculate.Status{}
	argument := os.Args[1:]
	if len(argument) != 2 {
		fmt.Println(`Usage  go run . credit/Debit Amount`)
		return
	}
	for {
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

		key := fmt.Sprintf("%s:%s", Name, Account)

		if _, ok := Calculate.Record[key]; !ok {
			user.Name = Name
			user.Account = Account
			user.Balance = 1000
			

		} else {
			user = Calculate.Record[key]
		}
		switch Action {
		case "debit":

			user.Debit(value)
		case "credit":
			if user.Balance < value {
				fmt.Println("The balance is less than Amount to be credited")
				return
			}
			user.Credit(value)
		}
		fmt.Printf("balance is  %v", user.Balance)
		fmt.Println("RECORD: ", Calculate.Record)
		os.Exit(0)

	}
}
