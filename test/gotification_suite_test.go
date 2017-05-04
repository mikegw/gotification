package gotification_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGotification(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gotification Suite")
}
