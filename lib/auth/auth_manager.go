package auth

import (
	"fmt"
	"time"

	"github.com/jellydator/ttlcache/v3"
)

var pins = make(map[string]string)

type account_login struct {
	account_id int
	timestamp  int64
}

// used a lib here to save some dev time
var logged_in = ttlcache.New[string, string](
	ttlcache.WithTTL[string, string](2 * time.Minute),
)

func init() {
	go logged_in.Start()
	pins["2859459814"] = "7386"
	pins["1434597300"] = "4557"
	pins["7089382418"] = "0075"
	pins["2001377812"] = "5950"
}

func Authenticate(account_id string, pin string) bool {
	if pins[account_id] == "" {
		fmt.Printf("User Not Found : %s\n", account_id)
		return false
	}
	if pins[account_id] == pin {
		//set logged in account
		fmt.Println("User Logged In, Session Valid For 2 Minutes...")
		logged_in.Set("logged_in_acct", pin, ttlcache.DefaultTTL)
		return true
	}
	return false
}

func Logout(account_id string) {
	if logged_in.Get("logged_in_acct").Value() == "" {
		fmt.Println("No account is currently authorized.")
		return
	}
	logged_in.Set("logged_in_acct", "", ttlcache.DefaultTTL)
	fmt.Printf("Account %s logged out.\n", account_id)
}

func GetLoggedInAccount() string {
	if logged_in.Get("logged_in_acct") != nil && logged_in.Get("logged_in_acct").Value() != "" {
		return logged_in.Get("logged_in_acct").Value()
	}
	return ""
}
