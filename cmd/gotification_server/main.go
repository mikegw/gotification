package main

import (
  "net/http"
  "github.com/mikegw/gotification/pkg/handlers"
  "github.com/mikegw/gotification/pkg/notification"

)

func main() {
    http.Handle("/notifications", handlers.HTTPRequest(notification.Create))
    http.ListenAndServe(":4000", nil)
}
