package gopi

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "net/http"
  "net/http/httptest"
)


func handler_test(req *http.Request) string {
  return "sweet"
}

func Test_First(t *testing.T) {
  // Create a request to pass to our handler. We don't have any query parameters for now, so we'll
  // pass 'nil' as the third parameter.
  req, err := http.NewRequest("GET", "/", nil)
  assert.Nil(t, err)

  // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
  rr := httptest.NewRecorder()
  router:= NewRouter()
  
  router.HandleRoute("/", handler_test)

  router.ServeHTTP(rr, req)
  // // Check the status code is what we expect.
  // if status := rr.Code; status != http.StatusOK {
  //     t.Errorf("handler returned wrong status code: got %v want %v",
  //         status, http.StatusOK)
  // }
  //
  // Check the response body is what we expect.
  expected := `sweet`
  if rr.Body.String() != expected {
      t.Errorf("handler returned unexpected body: got %v want %v",
          rr.Body.String(), expected)
  }
}
