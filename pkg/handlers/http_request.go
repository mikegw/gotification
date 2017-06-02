package handlers

import (
  "net/http"
  "fmt"
  "math/rand"
)

/*--- Public Data Types ---*/

type HTTPRequest func(http.ResponseWriter, *http.Request) (int, error)

func (handler HTTPRequest) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
    requestID := rand.Int()
    logRequest(request, requestID)
    responseCode, err := handler(writer, request)
    if err != nil {
        responseCode = http.StatusInternalServerError
        errorMessage := fmt.Sprintf("{\"message\":\"%s\"}", err.Error())
        fmt.Printf(errorMessage)
        http.Error(writer, errorMessage, responseCode)
    }
    logResponseCode(responseCode, requestID)
}



/*--- Private Functions ---*/

func logRequest(req *http.Request, requestID int) {
    fmt.Printf("[%d] %s %s\n", requestID, req.Method, req.URL.Path)
}

func logResponseCode(statusCode, requestID int) {
    statusText := http.StatusText(statusCode)
    fmt.Printf("[%d] %d %s\n", requestID, statusCode, statusText)
}
