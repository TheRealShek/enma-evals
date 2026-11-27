package service

import "evaltask/repo"

type Service struct {
	r *repo.Repo
}

func New(r *repo.Repo) *Service {
	return &Service{r: r}
}

func (s *Service) RegisterAndFetch(id, name string) (string, bool) {
	s.r.SaveUser(id, name)
	return s.r.FindUser(id)
}
