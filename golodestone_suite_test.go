package golodestone_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGolodestone(t *testing.T) {
	RegisterFailHandler(Fail)
	additionalReporters := make([]Reporter, 1)
	if os.Getenv("CI_REPORT") == "" {
		additionalReporters = append(additionalReporters, reporters.NewJUnitReporter(os.Getenv("CI_REPORT")))
	}
	RunSpecsWithDefaultAndCustomReporters(t, "Golodestone Suite", additionalReporters)
}
