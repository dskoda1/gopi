package gopi

import (
  "net/http"
  "fmt"
)
type ReqResFunc func(http.ResponseWriter, *http.Request) string
type NetHttpHandleFunc func(w http.ResponseWriter, r *http.Request)

// Base struct
type Router struct {

  // Keep track of all our handler functions
  routes map[string]ReqResFunc

}

// Default Constructor
func NewRouter() Router {
  r := Router{}
  r.routes = make(map[string]ReqResFunc)
  return r
}

// Add a route handler to the router
func (this *Router) HandleRoute(path string, f ReqResFunc) {
  this.routes[path] = f
  http.HandleFunc(path, this.assignRouteToHttp(f))
}

// Internal method that wraps the handler function provided
func (this *Router) assignRouteToHttp(f ReqResFunc) NetHttpHandleFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
			s := f(rw, r)
			fmt.Fprintf(rw, s)
	}
}

// Call the appropriate handler function, provided a request.
// Used for testing
func (this *Router) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
  s := this.routes[r.URL.Path](rw, r)
  fmt.Fprintf(rw, s)
}
