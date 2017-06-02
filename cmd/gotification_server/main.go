package main

import (
  "net/http"
  "github.com/mikegw/gotification/pkg/handlers"
  "github.com/mikegw/gotification/pkg/notification"
  "github.com/aws/aws-sdk-go/aws/credentials"
  "github.com/aws/aws-sdk-go/aws"
)


func main() {
    persistor := notification.NewSQSPersistor(awsConfiguration())
    http.Handle("/notifications", notificationCreationHandler(&persistor))
    http.ListenAndServe(":4000", nil)
}



func notificationCreationHandler(persistor notification.Persistor) handlers.HTTPRequest {
    creationHandler := func(w http.ResponseWriter, req *http.Request) (int, error) {
        return notification.Create(w, req, persistor)
    }
    return handlers.HTTPRequest(creationHandler)
}

func awsConfiguration() *aws.Config {
    playCredentials := credentials.NewCredentials(
        &credentials.SharedCredentialsProvider{Profile: "play"},
    )
    return &aws.Config{
        Region: aws.String("us-east-1"),
        Credentials: playCredentials,
    }
}
