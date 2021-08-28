package app

import (
	"net/http"
)

type Router struct {
	http.Handler
	Mapping map[string]func(w http.ResponseWriter, r *http.Request)
}

func NewRouter(mapping map[string]func(w http.ResponseWriter, r *http.Request)) *Router {
	return &Router{Mapping: mapping}
}

func (router Router) Handle(w http.ResponseWriter, r *http.Request) {
	if router.Mapping[r.Method] != nil {
		router.Mapping[r.Method](w, r)
		return
	}
	http.Error(w, "Not found", http.StatusNotFound)
}

func AddRoute(pattern string, mapping map[string]func(w http.ResponseWriter, r *http.Request)) {
	handler := NewRouter(mapping).Handle
	finalHandler := http.HandlerFunc(handler)
	http.Handle(pattern, JwtAuthentication(finalHandler))
}
