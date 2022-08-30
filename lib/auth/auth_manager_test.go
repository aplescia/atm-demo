package auth

import (
	"fmt"
	"testing"
)

var validLogins = map[string]string{
	"2859459814": "7386",
	"1434597300": "4557",
	"7089382418": "0075",
	"2001377812": "5950",
}

func TestLogin(t *testing.T) {
	invalidLogin := Authenticate("dummy", "val")
	if invalidLogin != "Authorization failed." {
		t.Errorf("Failed to prevent login of bad id")
	}
	for k, v := range validLogins {
		if Authenticate(k, v) != fmt.Sprintf("%s successfully authorized.", k) {
			t.Errorf("Failed to allow valid login")
		}
	}
}

func TestLogout(t *testing.T) {
	var out = Logout()
	if out != "No account is currently authorized." {
		t.Errorf("Not able to handle logout error")
	}
	for k, v := range validLogins {
		Authenticate(k, v)
		out = Logout()
		if out != fmt.Sprintf("Account %s logged out.\n", k) {
			t.Errorf("Not able to logout logged in acct")
		}
	}
}
