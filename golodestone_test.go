package golodestone_test

import (
	"time"

	. "github.com/cking/golodestone"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Golodestone", func() {
	Describe("Timer Resets", func() {
		It("should succeed when Weekly Reset is equal to 1477987200 for 1477435800", func() {
			reset := GetWeeklyResetFrom(time.Unix(1477435800, 0))
			Expect(reset.Unix()).To(Equal(int64(1477987200)))
		})
		It("should succeed when Daily Reset is equal to 1477494000 for 1477435800", func() {
			reset := GetDailyResetFrom(time.Unix(1477435800, 0))
			Expect(reset.Unix()).To(Equal(int64(1477494000)))
		})
		It("should succeed when Crafting Scrips Reset is equal to 1477555200 for 1477435800", func() {
			reset := GetCraftingScripsResetFrom(time.Unix(1477435800, 0))
			Expect(reset.Unix()).To(Equal(int64(1477555200)))
		})
		It("should succeed when Grand Company Reset is equal to 1477512000 for 1477435800", func() {
			reset := GetGrandCompanyResetFrom(time.Unix(1477435800, 0))
			Expect(reset.Unix()).To(Equal(int64(1477512000)))
		})
	})
})
