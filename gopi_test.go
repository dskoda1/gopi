package gopi

import (
  "testing"
  "net/http"
)

// Handler to demonstrate testing of request handlers
func handler_test(rw http.ResponseWriter, req *http.Request) string {
  rw.WriteHeader(http.StatusOK)
  return "sweet"
}

// Simply call the handler test util func with the arguments desired
// in the form of a HandlerUtil struct.
func Test_handler(t *testing.T) {
  HandlerTestUtil(t, &HandlerTestParams{
    "GET", "/", nil, handler_test, http.StatusOK, "sweet",
  })
}
