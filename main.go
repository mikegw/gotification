package main

import (
  "net/http"
  . "github.com/mikegw/gotification/lib/handlers"
  . "github.com/mikegw/gotification/lib"

)

func main() {
    http.Handle("/notifications", RequestHandler(CreateNotification))
    http.ListenAndServe(":4000", nil)
}
