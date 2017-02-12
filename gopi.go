package gopi

import (
  "net/http"
  "fmt"
)
type ReqResFunc func(http.ResponseWriter, *http.Request) string
type NetHttpHandleFunc func(w http.ResponseWriter, r *http.Request)

// Base struct
type Router struct {

  routes map[string]ReqResFunc


}

func NewRouter() Router {
  r := Router{}
  r.routes = make(map[string]ReqResFunc)
  return r
}

func (this *Router) HandleRoute(path string, f ReqResFunc) {
  this.routes[path] = f
  http.HandleFunc(path, this.assignRouteToHttp(f))
}

func (this *Router) assignRouteToHttp(f ReqResFunc) NetHttpHandleFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
			s := f(rw, r)
			fmt.Fprintf(rw, s)
	}
}

// Main function for long running servers
func (this *Router) ListenAndServe(addr string) {
	http.ListenAndServe(addr, nil)
}

// Call the appropriate handler function, provided a request.
// Used in testing
func (this *Router) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
  s := this.routes[r.URL.Path](rw, r)
  fmt.Fprintf(rw, s)
}
