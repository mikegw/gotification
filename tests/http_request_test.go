package gotification_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
    "net/http"
    "strings"
    "io/ioutil"
)


// This test demonstrates how http.Request bodys are used.
var _ = Describe("http.Request", func() {
  It("has a body", func() {
    expectedStringBody := "whooaa"
    body := strings.NewReader(expectedStringBody)

    request, err := http.NewRequest("GET", "http://test.ing", body)
    actualBody, err := ioutil.ReadAll(request.Body)

    if err != nil {
      Fail(err.Error())
    }

    Expect(string(actualBody)).To(Equal(expectedStringBody))
  })
})
