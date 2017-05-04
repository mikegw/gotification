package lib

import (
  "net/http"
  "fmt"
  "math/rand"
  . "github.com/mikegw/gotification/lib/handlers"
)

/*--- Public Data Types ---*/

type RequestHandler func(http.ResponseWriter, *http.Request) error

func (handler RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    requestId := rand.Int()
    logRequest(r, requestId)
    var responseCode int
    if err := handler(w, r); err == nil {
      responseCode = http.StatusOK
    } else {
      responseCode = handleError(err, w)
    }
    logResponseCode(responseCode, requestId)
}



/*--- Private Functions ---*/

func logRequest(req *http.Request, requestId int) {
  fmt.Printf("[%d] %s %s\n", requestId, req.Method, req.URL.Path)
}

func logResponseCode(statusCode, requestId int) {
  statusText := http.StatusText(statusCode)
  fmt.Printf("[%d] %d %s\n", requestId, statusCode, statusText)
}

func handleError(err error, writer http.ResponseWriter) (code int) {
  if _, badRequest := err.(BadRequest); badRequest {
    code = http.StatusBadRequest
    } else {
      code = http.StatusInternalServerError
    }
    http.Error(writer, err.Error(), int(code))
    return
  }
