package balance

import (
	"fmt"

	"github.com/aplescia/atm-demo/lib/audit"
	"github.com/aplescia/atm-demo/lib/auth"
)

var acct_balances = make(map[string]float64)

var total_funds float64

func init() {
	Reset()
}

// Reset initial map. Used in unit tests as well.
func Reset() {
	acct_balances["2859459814"] = 10.24
	acct_balances["1434597300"] = 90000.55
	acct_balances["7089382418"] = 0
	acct_balances["2001377812"] = 60
	total_funds = 90000.55 + 10.24 + 0 + 60
}

// Withdraw some amount from the currently logged in account barring some specified constraints. Returns a log string.
func Withdraw(amount int) string {
	logged_in_acct := auth.GetLoggedInAccount()
	if logged_in_acct == "" {
		return "Authorization required."
	}
	if acct_balances[logged_in_acct] < 0 {
		return "Your account is overdrawn! You may not make withdrawals at this time."
	}
	if amount%20 != 0 {
		return "Not a valid withdrawl amount. Amount must be a multiple of $20."
	}
	var amtFloat = float64(amount)
	//no $20 bills left in the ATM
	if total_funds < 20 {
		return "Unable to process your withdrawal at this time."
	}
	//we are using all of our money. also printing out requested error message.
	if total_funds-amtFloat < 0 {
		fmt.Println("Unable to dispense full amount requested at this time.")
		amtFloat = total_funds
	}
	acct_balances[logged_in_acct] = acct_balances[logged_in_acct] - amtFloat
	total_funds -= amtFloat
	if acct_balances[logged_in_acct] < 0 {
		fmt.Printf("Amount dispensed: %f\n", amtFloat)
		//overdraft
		acct_balances[logged_in_acct] -= float64(5)
		audit.Audit(logged_in_acct, fmt.Sprintf("-%f %f", amtFloat, acct_balances[logged_in_acct]))
		return fmt.Sprintf("You have been charged an overdraft fee of $5. Current balance: %f\n", acct_balances[logged_in_acct])
	} else {
		fmt.Printf("Amount dispensed: %f\n", amtFloat)
		audit.Audit(logged_in_acct, fmt.Sprintf("-%f %f", amtFloat, acct_balances[logged_in_acct]))
		return fmt.Sprintf("Current balance: %f\n", acct_balances[logged_in_acct])
	}
}

// Deposit some amount into the currently logged in account barring some specified constraints. Returns a log string.
func Deposit(amount int) string {
	logged_in_acct := auth.GetLoggedInAccount()
	if logged_in_acct == "" {
		return "Authorization required."
	}
	acct_balances[logged_in_acct] = acct_balances[logged_in_acct] + float64(amount)
	total_funds -= float64(amount)
	audit.Audit(logged_in_acct, fmt.Sprintf("%f %f", float64(amount), acct_balances[logged_in_acct]))
	return fmt.Sprintf("Current balance: %f\n", acct_balances[logged_in_acct])
}

// Get the balance of the currently logged in account.
func Balance() string {
	logged_in_acct := auth.GetLoggedInAccount()
	if logged_in_acct == "" {
		return "Authorization required."
	}
	return fmt.Sprintf("Current balance: %f\n", acct_balances[logged_in_acct])
}

// return the transaction history of the currently logged in account.
func History() []string {
	logged_in_acct := auth.GetLoggedInAccount()
	if logged_in_acct == "" {
		return []string{"Authorization required."}
	}
	return audit.History(logged_in_acct)
}
