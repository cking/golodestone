package golodestone_test

import (
	. "github.com/cking/golodestone"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Util", func() {
	Describe("The url generator", func() {
		It("should fail on malformed urls", func() {
			_, err := BuildURL("not%avalid#url#to parse")
			Expect(err).To(HaveOccurred())
		})
		It("should fail on non-lodestone URLs", func() {
			_, err := BuildURL("http://google.de/something")
			Expect(err).To(HaveOccurred())
		})
		It("should fail to generate a valid URL, even if a valid hostname is given", func() {
			_, err := BuildURL("http://na.finalfantasyxiv.com/worldstatus")
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

	Describe("The scraper", func() {
		It("should fail on non-lodestone links", func() {
			_, err := QueryLodestone("http://example.com")
			Expect(err).To(HaveOccurred())
		})

		PIt("should fail on unfetchable requests")
		PIt("should fail on non-html pages")

		It("should fail on unexpected page layout", func() {
			_, err := QueryLodestone("not-existing-url-lalala")
			Expect(err).To(HaveOccurred())
		})
		It("should return a valid html-node", func() {
			_, err := QueryLodestone("worldstatus")
			Expect(err).To(Succeed())
		})
	})
})
