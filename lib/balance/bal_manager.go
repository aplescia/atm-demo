package balance

import (
	"fmt"

	"github.com/aplescia/atm-demo/lib/auth"
)

var acct_balances = make(map[string]float64)

func init() {
	acct_balances["2859459814"] = 10.24
	acct_balances["1434597300"] = 90000.55
	acct_balances["7089382418"] = 0
	acct_balances["2001377812"] = 60
}

func Withdraw(amount int) {
	logged_in_acct := auth.GetLoggedInAccount()
	if logged_in_acct == "" {
		fmt.Println("Authorization required.")
		return
	}
	if amount%20 != 0 {
		fmt.Println("Not a valid withdrawl amount. Amount must be a multiple of $20.")
		return
	}
	acct_balances[logged_in_acct] = acct_balances[logged_in_acct] - float64(amount)
	if acct_balances[logged_in_acct] < 0 {
		fmt.Printf("Amount dispensed: %d\n", amount)
		fmt.Printf("You have been charged an overdraft fee of $5. Current balance: %f\n", acct_balances[logged_in_acct])
	}

}
