package handler

import (
	"evaltask/cache"
	"evaltask/repo"
	"evaltask/service"
)

func HandleRequest(id, name string) (string, bool) {
	svc := service.New(repo.New(cache.New()))
	return svc.RegisterAndFetch(id, name)
}
