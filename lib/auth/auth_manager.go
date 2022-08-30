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

func Authenticate(account_id string, pin string) string {
	if pins[account_id] == "" {
		return "Authorization failed."
	}
	if pins[account_id] == pin {
		//set logged in account
		logged_in.Set("logged_in_acct", account_id, ttlcache.DefaultTTL)
		return fmt.Sprintf("%s successfully authorized.", account_id)
	}
	return "Authorization failed."
}

func Logout() string {
	if GetLoggedInAccount() == "" {
		return "No account is currently authorized."
	}
	var logged_in_acct = GetLoggedInAccount()
	logged_in.Set("logged_in_acct", "", ttlcache.DefaultTTL)
	return fmt.Sprintf("Account %s logged out.\n", logged_in_acct)
}

func GetLoggedInAccount() string {
	if logged_in.Get("logged_in_acct") != nil && logged_in.Get("logged_in_acct").Value() != "" {
		return logged_in.Get("logged_in_acct").Value()
	}
	return ""
}
