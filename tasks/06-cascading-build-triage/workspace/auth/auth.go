package auth

import "evaltask/core"

func IsPro(u core.User) bool {
	return u.Tier == "pro"
}
