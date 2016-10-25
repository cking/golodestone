package golodestone_test

import (
	. "github.com/cking/golodestone"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Util", func() {
	Describe("The url generator", func() {
		It("should fail on non-lodestone URLs", func() {
			_, err := BuildURL("http://google.de/something")
			Expect(err).To(HaveOccurred())
		})
		It("should create a european lodestone URL", func() {
			url, err := BuildURL("worldstatus")
			Expect(err).To(Succeed())
			Expect(url).To(HavePrefix("http://eu.finalfantasyxiv.com/"))
		})
		It("should prefix the URL path with /lodestone/", func() {
			url, err := BuildURL("worldstatus")
			Expect(err).To(Succeed())
			Expect(url).To(HavePrefix("http://eu.finalfantasyxiv.com/lodestone/"))
		})
		It("should ignore existing /lodestone/ prefixes", func() {
			url, err := BuildURL("/lodestone/worldstatus")
			Expect(err).To(Succeed())
			Expect(url).To(HavePrefix("http://eu.finalfantasyxiv.com/lodestone/"))
		})
		It("should overwrite the host to the european lodestone one", func() {
			url, err := BuildURL("http://na.finalfantasyxiv.com/lodestone/worldstatus")
			Expect(err).To(Succeed())
			Expect(url).To(HavePrefix("http://eu.finalfantasyxiv.com/lodestone/"))
		})
	})
})
