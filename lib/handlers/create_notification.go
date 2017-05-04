package handlers

import (
    "fmt"
    "io"
    "io/ioutil"
    "encoding/json"
    "net/http"
)

/*--- Public Data Types ---*/

type Notification struct {
    Payload string `json:"payload"`
}


type SavedNotification struct {
    Payload string `json:"payload"`
    Id string `json:"id"`
}

// Implements the 'error' interface
type BadRequest struct {
    Message string
}

func (err BadRequest) Error() string {
    return err.Message
}



/*--- Public Functions ---*/

func CreateNotification(w http.ResponseWriter, request *http.Request) error {
    notification, err := fetchNotification(request)
    if err != nil {
        return err
    }

    savedNotification, err := save(notification)
    if err != nil {
        return err
    }

    notificationResponse, err := buildResponse(savedNotification)
    if err != nil {
        return err
    }

    writeResponse(notificationResponse, w)
    return nil
}



/*--- Private Functions ---*/

func fetchNotification(request *http.Request) (Notification, error) {
    var notification Notification

    rawBody, err := readBody(request)
    if err != nil {
        return notification, err
    }

    err = json.Unmarshal(rawBody, &notification)
    if err != nil {
        return notification, BadRequest{"Invalid JSON"}
    }

    err = validateNotification(notification)
    if err != nil {
        return notification, err
    }

    return notification, nil
}

func save(notification Notification) (SavedNotification, error) {
    var savedNotification SavedNotification
    savedNotification.Payload = notification.Payload
    savedNotification.Id = "some id"
    return savedNotification, nil
}

func buildResponse(notification SavedNotification) (string, error) {
    var response string
    responseBytes, err := json.Marshal(notification)
    if err != nil {
        return response, err
    }
    response = string(responseBytes)
    return response, nil
}

func readBody(request *http.Request) ([]byte, error) {
    return ioutil.ReadAll(io.LimitReader(request.Body, 256))
}

func writeResponse(notificationResponse string, writer http.ResponseWriter) {
    fmt.Fprintf(writer, notificationResponse)
}

func validateNotification(notification Notification) error {
    if notification.Payload == "" {
        return BadRequest{"Missing JSON parameter: payload"}
    }
    return nil
}
