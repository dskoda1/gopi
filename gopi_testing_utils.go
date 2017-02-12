package gopi

import (
  "io"
  "testing"
  "net/http"
  "net/http/httptest"

  "github.com/stretchr/testify/assert"
)

// Util struct to encapsulate arguments for checking
// on typical requests and responses to an api
type HandlerTestParams struct {
  ReqType string
  Path string
  QueryString io.Reader
  Handler ReqResFunc
  ExpectedStatus int
  ExpectedBody string
}
// Util function to spin up a Gopi router, and test behaviour
// of a handler function provided a HandlerUtil struct
func HandlerTestUtil(t *testing.T, htp *HandlerTestParams) {

  // Create the request provided fields from the HandlerUtil.
  req, err := http.NewRequest(htp.ReqType, htp.Path, htp.QueryString)
  assert.Nil(t, err)

  // Create a recorder to observe changes to later.
  rr := httptest.NewRecorder()

  router := NewRouter()

  // Add the handler to the router, and call it with our recorder.
  router.HandleRoute(htp.Path, htp.Handler)
  router.ServeHTTP(rr, req)

  // Check the response status and body.
  assert.Equal(t, htp.ExpectedStatus, rr.Code)
  assert.Equal(t, htp.ExpectedBody, rr.Body.String())

}
