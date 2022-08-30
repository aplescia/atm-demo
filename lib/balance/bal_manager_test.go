package balance

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aplescia/atm-demo/lib/auth"
)

var validLogins = map[string]string{
	"2859459814": "7386",
	"1434597300": "4557",
	"7089382418": "0075",
	"2001377812": "5950",
}

func TestBalance(t *testing.T) {
	var noLoginBalance = Balance()
	if noLoginBalance != "Authorization required." {
		t.Errorf("Could not handle login and balance retrieval")
	}
	for k, v := range validLogins {
		auth.Authenticate(k, v)
		var balance = Balance()
		if balance != fmt.Sprintf("Current balance: %f\n", acct_balances[k]) {
			t.Errorf("Could not get correct balance")
		}
	}
}

func TestWithdraw(t *testing.T) {
	var noLoginBalance = Balance()
	if noLoginBalance != "Authorization required." {
		t.Errorf("Could not handle login and withdraw")
	}
	for k, v := range validLogins {
		auth.Authenticate(k, v)
		var balance = Withdraw(30)
		if balance != "Not a valid withdrawl amount. Amount must be a multiple of $20." {
			t.Errorf("Could not handle non 20 modulo withdrawal error")
		}
	}
	auth.Authenticate("2859459814", "7386")
	Withdraw(1_000_000)
	fmt.Printf("Total Funds : %f\n", total_funds)
	var noMoney = Withdraw(1_000_000)
	if noMoney != "Your account is overdrawn! You may not make withdrawals at this time." {
		t.Errorf("Unable to handle maxed out acct")
	}
	auth.Authenticate("1434597300", "4557")
	noMoney = Withdraw(20)
	if noMoney != "Unable to process your withdrawal at this time." {
		t.Errorf("Unable to handle maxed out ATM")
	}

}

func TestDeposit(t *testing.T) {
	var noLoginDeposit = Deposit(100)
	if noLoginDeposit != "Authorization required." {
		t.Errorf("Could not handle login and deposit")
	}
	auth.Authenticate("2859459814", "7386")
	var expected = acct_balances["2859459814"] + float64(30)
	var deposit = Deposit(30)

	if deposit != fmt.Sprintf("Current balance: %f\n", expected) {
		t.Errorf("Incorrect math")
	}

}

func TestHistory(t *testing.T) {
	auth.Authenticate("2859459814", "7386")
	Deposit(30)
	Withdraw(20)
	Withdraw(40)
	var hist = History()
	if len(hist) < 3 {
		t.Errorf("Not enough history messages")
	}
	//doing it this way as timestamp in logs make direct comparison tricky
	if !strings.Contains(hist[2], "30") {
		t.Errorf("Not in reverse order")
	}
}
