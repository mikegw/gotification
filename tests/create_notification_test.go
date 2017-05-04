package gotification_test

import (
  . "github.com/mikegw/gotification/lib/handlers"
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

func createNotification(notificationString string) ([]byte, error) {
    var responseBody []byte

    body := strings.NewReader(notificationString)
    mockRequest := httptest.NewRequest("POST", "http://example.com/foo", body)
    responseRecorder := httptest.NewRecorder()

    err := CreateNotification(responseRecorder, mockRequest)
    if err != nil {
        return responseBody, err
    }

    response := responseRecorder.Result()
    responseBody, _ = ioutil.ReadAll(response.Body)
    return responseBody, nil
}

var _ = Describe("CreateNotification", func() {
  Context("with valid input", func() {
    var responseNotification savedNotification

    BeforeEach(func() {
      responseBody, _ := createNotification("{\"payload\":\"Hi\"}")
      json.Unmarshal(responseBody, &responseNotification)
    })

    It("returns the notification", func() {
      Expect(responseNotification.Payload).To(Equal("Hi"))
    })

    It("adds an ID to the notification", func() {
      Expect(responseNotification.Id).NotTo(BeNil())
    })
  })

  Context("with invalid json", func() {
    var err error

    BeforeEach(func() {
      _, err = createNotification("totally a notification")
    })

    It("returns an error", func() {
      Expect(err.Error()).To(Equal("Invalid JSON"))
    })
  })

  Context("with invalid notification", func() {
    var err error

    BeforeEach(func() {
      _, err = createNotification("{\"some\":\"json\"}")
    })

    It("returns an error", func() {
      if err == nil {
        Fail("No error")
      }
      Expect(err.Error()).To(Equal("Missing JSON parameter: payload"))
    })
  })
})
