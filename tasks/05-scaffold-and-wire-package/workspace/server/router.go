package server

type Router struct {
	limit int
}

func NewRouter(limit int) *Router {
	return &Router{limit: limit}
}

func (r *Router) Handle() int {
	return 200
}
