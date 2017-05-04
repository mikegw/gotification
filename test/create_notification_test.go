package gotification_test

import (
  "github.com/mikegw/gotification/pkg/notification"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "net/http/httptest"
  "strings"
  "io/ioutil"
  "encoding/json"
)
type savedNotification struct {
  Payload string
  Id string
}

type errorResponse struct {
  Message string
}

func createNotification(notificationString string) ([]byte, int, error) {
    var responseBody []byte
    var responseCode int

    body := strings.NewReader(notificationString)
    mockRequest := httptest.NewRequest("POST", "http://example.com/foo", body)
    responseRecorder := httptest.NewRecorder()

    _, err := notification.Create(responseRecorder, mockRequest)
    if err != nil {
        return responseBody, 0, err
    }

    response := responseRecorder.Result()

    responseBody, _ = ioutil.ReadAll(response.Body)
    responseCode = response.StatusCode
    return responseBody, responseCode, nil
}

var _ = Describe("CreateNotification", func() {
  Context("with valid input", func() {
    var responseNotification savedNotification
    var responseCode int

    BeforeEach(func() {
      responseBody, code, _ := createNotification("{\"payload\":\"Hi\"}")
      responseCode = code
      json.Unmarshal(responseBody, &responseNotification)
    })

    It("returns a 200 response", func() {
      Expect(responseCode).To(Equal(200))
    })

    It("returns the notification", func() {
      Expect(responseNotification.Payload).To(Equal("Hi"))
    })

    It("adds an ID to the notification", func() {
      Expect(responseNotification.Id).NotTo(BeNil())
    })
  })

  Context("with invalid json", func() {
    var response errorResponse
    var responseCode int

    BeforeEach(func() {
      responseBody, code, _ := createNotification("totally a notification")
      responseCode = code
      json.Unmarshal(responseBody, &response)
    })

    It("returns an error", func() {
      Expect(response.Message).To(Equal("Invalid JSON"))
    })
  })

  Context("with invalid notification", func() {
    var response errorResponse
    var responseCode int

    BeforeEach(func() {
      responseBody, code, _ := createNotification("{\"key\":\"value\"}")
      responseCode = code
      json.Unmarshal(responseBody, &response)
    })

    It("returns an error", func() {
      Expect(response.Message).To(Equal("Missing JSON parameter: payload"))
    })
  })
})
