package golodestone_test

import (
	"time"

	. "github.com/cking/golodestone"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Golodestone", func() {
	Describe("Resettimes", func() {
		Context("with the time set to 2016-10-25 22:50:00 UTC", func() {
			from := time.Unix(1477435800, 0)
			resets := GetResetsFrom(from)
			zone, _ := time.LoadLocation("")
			It("should return the expected daily reset 2016-10-26 15:00:00 UTC", func() {
				Expect(resets.Daily).To(BeTemporally("==", time.Date(2016, 10, 26, 15, 0, 0, 0, zone)))
			})
			It("should return the expected grand company reset 2016-10-26 20:00:00 UTC", func() {
				Expect(resets.GrandCompany).To(BeTemporally("==", time.Date(2016, 10, 26, 20, 0, 0, 0, zone)))
			})
			It("should return the expected weekly reset 2016-11-01 08:00:00 UTC", func() {
				Expect(resets.Weekly).To(BeTemporally("==", time.Date(2016, 11, 01, 8, 0, 0, 0, zone)))
			})
			It("should return the expected scrip reset 2016-10-27 08:00:00 UTC", func() {
				Expect(resets.Scrip).To(BeTemporally("==", time.Date(2016, 10, 27, 8, 0, 0, 0, zone)))
			})
		})
	})
})
