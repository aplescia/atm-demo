package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aplescia/atm-demo/lib/auth"
	"github.com/aplescia/atm-demo/lib/balance"
)

func main() {
	fmt.Println("ATM initiated...")
	for {
		var command string
		var input string    //optional
		var inputTwo string //optional
		fmt.Scanf("%s %s %s", &command, &input, &inputTwo)
		//only check for valid commands
		if command == "exit" {
			os.Exit(3)
		} else if command == "authorize" {
			var result = auth.Authenticate(input, inputTwo)
			fmt.Println(result)
		} else if command == "withdraw" {
			i, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Error parsing input for withdraw command.")
			} else {
				var result = balance.Withdraw(i)
				fmt.Println(result)
			}
		} else if command == "deposit" {
			i, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Error parsing input for deposit command.")
			} else {
				var result = balance.Deposit(i)
				fmt.Println(result)
			}
		} else if command == "history" {
			var history = balance.History()
			for i := range history {
				log.Println(history[i])
			}
		} else if command == "balance" {
			var result = balance.Balance()
			fmt.Println(result)
		} else if command == "logout" {
			var result = auth.Logout()
			fmt.Println(result)
		}
	}
}
