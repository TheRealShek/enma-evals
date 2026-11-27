package app

import (
	"evaltask/auth"
	"evaltask/billing"
	"evaltask/core"
)

func Summary(id, plan string) (bool, int) {
	u := core.User{ID: id, Plan: plan}
	return auth.IsPro(u), billing.MonthlyFee(u)
}
