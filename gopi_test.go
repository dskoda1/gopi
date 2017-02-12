package gopi

import (
  "testing"
  "net/http"
  "net/http/httptest"

  "github.com/stretchr/testify/assert"
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

func TestAssignRouteToHttp(t *testing.T) {
  router := NewRouter()
  handler_func_wrapper := router.assignRouteToHttp(handler_test)

  // Create a recorder to observe changes to later.
  rr := httptest.NewRecorder()
  // Create the request provided fields from the HandlerUtil.
  req, err := http.NewRequest("GET", "/", nil)
  assert.Nil(t, err)

  handler_func_wrapper(rr, req)

  assert.Equal(t, http.StatusOK, rr.Code)
  assert.Equal(t, "sweet", rr.Body.String())

}
