package service

import "evaltask/store"

func LookupOrDefault(s store.Store, key, fallback string) string {
	if v := s.Get(key); v != "" {
		return v
	}

	return fallback
}
