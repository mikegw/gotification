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
  ID string
}

type errorResponse struct {
  Message string
}

type mockPersistor struct {
  Persisted bool
}

func (m *mockPersistor) Persist(notification.Model) (string, error) {
    m.Persisted = true
    return "some id", nil
}

type creationResult struct {
    ResponseBody []byte
    ResponseCode int
    Persisted bool
    CreationError error
}

func createNotification(notificationString string) (creationResult) {
    /*--- Setup ---*/
    body := strings.NewReader(notificationString)
    mockRequest := httptest.NewRequest("POST", "http://example.com/foo", body)
    responseRecorder := httptest.NewRecorder()
    persistor := mockPersistor{}


    /*--- Run Code ---*/
    _, err := notification.Create(responseRecorder, mockRequest, &persistor)


    /*--- Build Result ---*/
    result := creationResult{
        Persisted: persistor.Persisted,
        CreationError: err,
    }
    if err == nil {
        response := responseRecorder.Result()
        result.ResponseBody, _ = ioutil.ReadAll(response.Body)
        result.ResponseCode = response.StatusCode
    }
    return result
}

var _ = Describe("CreateNotification", func() {
  Context("with valid input", func() {
    var result creationResult
    var responseNotification savedNotification

    BeforeEach(func() {
      result = createNotification("{\"payload\":\"Hi\"}")
      json.Unmarshal(result.ResponseBody, &responseNotification)
    })

    It("persists the data", func() {
      Expect(result.Persisted).To(BeTrue())
    })

    It("returns a 200 response", func() {
      Expect(result.ResponseCode).To(Equal(200))
    })

    It("returns the notification", func() {
      Expect(responseNotification.Payload).To(Equal("Hi"))
    })

    It("adds an ID to the notification", func() {
      Expect(responseNotification.ID).NotTo(BeNil())
    })
  })

  Context("with invalid json", func() {
    var result creationResult
    var jsonResponse errorResponse

    BeforeEach(func() {
      result = createNotification("totally a notification")
      json.Unmarshal(result.ResponseBody, &jsonResponse)
    })

    It("returns an error", func() {
      Expect(jsonResponse.Message).To(Equal("Invalid JSON"))
    })
  })

  Context("with invalid notification", func() {
    var result creationResult
    var jsonResponse errorResponse

    BeforeEach(func() {
      result = createNotification("{\"key\":\"value\"}")
      json.Unmarshal(result.ResponseBody, &jsonResponse)
    })

    It("returns an error", func() {
      Expect(jsonResponse.Message).To(Equal("Missing JSON parameter: payload"))
    })
  })
})
