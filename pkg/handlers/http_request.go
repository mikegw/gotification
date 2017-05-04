package handlers

import (
  "net/http"
  "fmt"
  "math/rand"
)

/*--- Public Data Types ---*/

type HTTPRequest func(http.ResponseWriter, *http.Request) (int, error)

func (handler HTTPRequest) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
    requestId := rand.Int()
    logRequest(request, requestId)
    responseCode, err := handler(writer, request)
    if err != nil {
        responseCode = http.StatusInternalServerError
        errorMessage := fmt.Sprintf("{\"message\":\"%s\"}", err.Error())
        http.Error(writer, errorMessage, responseCode)
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
