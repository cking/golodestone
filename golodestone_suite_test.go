package golodestone_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGolodestone(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Golodestone Suite")
}
