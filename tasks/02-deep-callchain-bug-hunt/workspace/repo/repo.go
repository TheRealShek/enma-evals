package repo

import "evaltask/cache"

type Repo struct {
	c *cache.Mem
}

func New(c *cache.Mem) *Repo {
	return &Repo{c: c}
}

func (r *Repo) SaveUser(id, name string) {
	r.c.Set(id, name)
}

func (r *Repo) FindUser(id string) (string, bool) {
	return r.c.Get(id)
}
