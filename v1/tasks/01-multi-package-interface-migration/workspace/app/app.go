package app

import (
	"evaltask/memory"
	"evaltask/service"
)

func ReadSetting(cfg map[string]string, key string) string {
	return service.LookupOrDefault(&memory.MemStore{Data: cfg}, key, "default")
}
