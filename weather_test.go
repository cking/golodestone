package golodestone_test

import (
	"time"

	. "github.com/cking/golodestone"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Weather", func() {
	Describe("The weather calculator", func() {
		It("should fail on invalid zone str", func() {
			_, err := GetCurrentWeatherByString("invalid zone")
			Expect(err).To(HaveOccurred())
		})
		It("should fail on invalid zone id", func() {
			_, err := GetCurrentWeather(Zone(-1))
			Expect(err).To(HaveOccurred())
		})
		It("should return Gloomy Weather for MorDhonaZone at 10/25/2016 @ 7:25pm (UTC)", func() {
			weather, err := GetWeather(TimeFrom(time.Unix(1477423542, 0)), MorDhonaZone)
			Expect(err).To(Succeed())
			Expect(weather).To(Equal(GloomyWeather))
		})
		It("should return Gloomy Weather for \"Mor Dhona\" at 10/25/2016 @ 7:25pm (UTC)", func() {
			weather, err := GetWeatherByString(TimeFrom(time.Unix(1477423542, 0)), "Mor Dhona")
			Expect(err).To(Succeed())
			Expect(weather).To(Equal(GloomyWeather))
		})
		It("should return \"Gloomy weather\" for GloomyWeather Weather", func() {
			name, err := WeatherToString(GloomyWeather)
			Expect(err).To(Succeed())
			Expect(name).To(Equal("Gloomy weather"))
		})
		It("should fail on invalid weather id", func() {
			_, err := WeatherToString(Weather(-1))
			Expect(err).To(HaveOccurred())
		})
	})
})
