package app

import (
	"net/http"

	"github.com/dhaaana/go-http-server/middleware"
)

type Router struct {
	routes map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) Get(path string, handler http.HandlerFunc) {
	r.addRoute(http.MethodGet, path, handler)
}

func (r *Router) Post(path string, handler http.HandlerFunc) {
	r.addRoute(http.MethodPost, path, handler)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handlers, ok := r.routes[req.Method]; ok {
		if handler, ok := handlers[req.URL.Path]; ok {
			handler(w, req)
			return
		}
	}
	http.NotFound(w, req)
}

func (r *Router) addRoute(method, path string, handler http.HandlerFunc) {
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]http.HandlerFunc)
	}

	r.routes[method][path] = middleware.Logging(handler)
}
