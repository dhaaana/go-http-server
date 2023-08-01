package app

import (
	"net/http"
	"strings"

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

func (r *Router) Put(path string, handler http.HandlerFunc) {
	r.addRoute(http.MethodPut, path, handler)
}

func (r *Router) Delete(path string, handler http.HandlerFunc) {
	r.addRoute(http.MethodDelete, path, handler)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Handle OPTIONS request for CORS preflight
	if req.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.WriteHeader(http.StatusOK)
		return
	}

	if handlers, ok := r.routes[req.Method]; ok {
		for path, handler := range handlers {
			if matchPath(path, req.URL.Path) {
				handler(w, req)
				return
			}
		}
	}
	http.NotFound(w, req)
}

func (r *Router) addRoute(method, path string, handler http.HandlerFunc) {
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]http.HandlerFunc)
	}

	r.routes[method][path] = middleware.Logging(middleware.Cors(handler))
}

func matchPath(routePath, reqPath string) bool {
	routeParts := strings.Split(routePath, "/")
	reqParts := strings.Split(reqPath, "/")

	if len(routeParts) != len(reqParts) {
		return false
	}

	for i := 0; i < len(routeParts); i++ {
		if strings.HasPrefix(routeParts[i], ":") {
			// Named parameter, skip
			continue
		}

		if routeParts[i] != reqParts[i] {
			return false
		}
	}

	return true
}
