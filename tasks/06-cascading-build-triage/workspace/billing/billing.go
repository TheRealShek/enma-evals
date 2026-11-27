package billing

import "evaltask/core"

func MonthlyFee(u core.User) int {
	if u.Plan == "pro" {
		return 20
	}

	return 0
}
