package gotification_test

import (
  "github.com/mikegw/gotification/pkg/notification"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("SQS", func(){
  Describe("SQSPersistor", func(){
    It("sends the persisted model to SQS", func(){
      mockSender := notification.MockMessageSender{}
      persistor := notification.SQSPersistorImpl{&mockSender}
      model := notification.Model{
        Payload: "{\"message\":\"hi\"}",
      }
      persistor.Persist(model)
      Expect(mockSender.InputBody()).To(Equal("{\"message\":\"hi\"}"))
    })
  })
})
