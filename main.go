package main

import (
	"fmt"
	"strconv"

	"github.com/aplescia/atm-demo/lib/auth"
	"github.com/aplescia/atm-demo/lib/balance"
)

func main() {
	fmt.Println("ATM initiated...")
	for {
		var command string
		var input string
		var inputTwo string //optional
		fmt.Scanf("%s %s %s", &command, &input, &inputTwo)
		if command == "authorize" {
			auth.Authenticate(input, inputTwo)
		} else if command == "withdraw" {
			i, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Error parsing input for withdraw command.")
			} else {
				balance.Withdraw(i)
			}
		}
	}
}
