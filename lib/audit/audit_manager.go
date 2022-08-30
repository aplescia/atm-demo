package audit

import (
	"fmt"
)

var logs = make(map[string][]string)

// Record a given message for the account_id.
func Audit(account_id string, message string) {
	logs[account_id] = append(logs[account_id], message)
}

// For the given account_id, return a reverse order list of command history.
func History(account_id string) []string {
	var returnSlice []string
	if logs[account_id] == nil {
		fmt.Println("No history found.")
	} else {
		for i := len(logs[account_id]) - 1; i >= 0; i-- {
			returnSlice = append(returnSlice, logs[account_id][i])
		}
	}
	return returnSlice
}
