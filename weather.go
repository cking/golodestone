package golodestone

import (
	"fmt"
	"time"
)

type weatherChanceFunction func(int) Weather
type weatherForecastChance struct {
	chance  int
	weather Weather
}

// Weather Type of weather
type Weather int

const (
	// ClearWeather Clear weather
	ClearWeather Weather = iota
	// CloudyWeather Cloudy weather
	CloudyWeather Weather = iota
	// FairWeather Fair weather
	FairWeather Weather = iota
	// FoggyWeather Foggy weather
	FoggyWeather Weather = iota
	// RainyWeather Rain
	RainyWeather Weather = iota
	// WindyWeather Windy weather
	WindyWeather Weather = iota
	// ShowerWeather Rain shower
	ShowerWeather Weather = iota
	// GalesWeather Gale
	GalesWeather Weather = iota
	// ThunderWeather Thunder
	ThunderWeather Weather = iota
	// ThunderstormWeather Thunderstorm
	ThunderstormWeather Weather = iota
	// HeatWaveWeather Heat waves
	HeatWaveWeather Weather = iota
	// GloomyWeather Gloomy weather
	GloomyWeather Weather = iota
	// SnowyWeather Snow
	SnowyWeather Weather = iota
	// BlizzardWeather Blizzards
	BlizzardWeather Weather = iota
	// UmbralWindWeather Umbral winds
	UmbralWindWeather Weather = iota
	// DustStormWeather Dust storms
	DustStormWeather Weather = iota
	// UmbralStaticWeather Umbral statics
	UmbralStaticWeather Weather = iota
)

var (
	zoneMap    = weatherZoneMap()
	weatherMap = weatherChanceMap()
)

func weatherZoneMap() map[string]string {
	m := make(map[string]string)
	m["uldah"] = "ul'dah"
	m["azys la"] = "azys lla"
	return m
}

func weatherChanceMap() map[string]weatherChanceFunction {
	m := make(map[string]weatherChanceFunction)
	m["limsa lominsa"] = createWeatherForecast(
		RainyWeather,
		weatherForecastChance{20, CloudyWeather},
		weatherForecastChance{50, ClearWeather},
		weatherForecastChance{80, FairWeather},
		weatherForecastChance{90, FoggyWeather},
	)

	m["middle la noscea"] = createWeatherForecast(
		RainyWeather,
		weatherForecastChance{20, CloudyWeather},
		weatherForecastChance{50, ClearWeather},
		weatherForecastChance{70, FairWeather},
		weatherForecastChance{80, WindyWeather},
		weatherForecastChance{90, FoggyWeather},
	)

	m["lower la noscea"] = createWeatherForecast(
		RainyWeather,
		weatherForecastChance{20, CloudyWeather},
		weatherForecastChance{50, ClearWeather},
		weatherForecastChance{70, FairWeather},
		weatherForecastChance{80, WindyWeather},
		weatherForecastChance{90, FoggyWeather},
	)

	m["eastern la noscea"] = createWeatherForecast(
		ShowerWeather,
		weatherForecastChance{5, FoggyWeather},
		weatherForecastChance{50, ClearWeather},
		weatherForecastChance{80, FairWeather},
		weatherForecastChance{90, CloudyWeather},
		weatherForecastChance{95, RainyWeather},
	)

	m["western la noscea"] = createWeatherForecast(
		GalesWeather,
		weatherForecastChance{10, FoggyWeather},
		weatherForecastChance{40, ClearWeather},
		weatherForecastChance{60, FairWeather},
		weatherForecastChance{80, CloudyWeather},
		weatherForecastChance{90, WindyWeather},
	)

	m["upper la noscea"] = createWeatherForecast(
		ThunderstormWeather,
		weatherForecastChance{30, ClearWeather},
		weatherForecastChance{50, FairWeather},
		weatherForecastChance{70, CloudyWeather},
		weatherForecastChance{80, FoggyWeather},
		weatherForecastChance{90, ThunderWeather},
	)

	m["outer la noscea"] = createWeatherForecast(
		RainyWeather,
		weatherForecastChance{30, ClearWeather},
		weatherForecastChance{50, FairWeather},
		weatherForecastChance{70, CloudyWeather},
		weatherForecastChance{85, FoggyWeather},
	)

	m["mist"] = createWeatherForecast(
		RainyWeather,
		weatherForecastChance{20, CloudyWeather},
		weatherForecastChance{50, ClearWeather},
		weatherForecastChance{80, FairWeather},
		weatherForecastChance{90, FoggyWeather},
	)

	m["gridania"] = createWeatherForecast(
		FairWeather,
		weatherForecastChance{20, RainyWeather},
		weatherForecastChance{30, FoggyWeather},
		weatherForecastChance{40, CloudyWeather},
		weatherForecastChance{55, FairWeather},
		weatherForecastChance{85, ClearWeather},
	)

	m["central shroud"] = createWeatherForecast(
		FairWeather,
		weatherForecastChance{5, ThunderWeather},
		weatherForecastChance{20, RainyWeather},
		weatherForecastChance{30, FoggyWeather},
		weatherForecastChance{40, CloudyWeather},
		weatherForecastChance{55, FairWeather},
		weatherForecastChance{85, ClearWeather},
	)

	m["east shroud"] = createWeatherForecast(
		FairWeather,
		weatherForecastChance{5, ThunderWeather},
		weatherForecastChance{20, RainyWeather},
		weatherForecastChance{30, FoggyWeather},
		weatherForecastChance{40, CloudyWeather},
		weatherForecastChance{55, FairWeather},
		weatherForecastChance{85, ClearWeather},
	)

	m["south shroud"] = createWeatherForecast(
		ClearWeather,
		weatherForecastChance{5, FoggyWeather},
		weatherForecastChance{10, ThunderstormWeather},
		weatherForecastChance{25, ThunderWeather},
		weatherForecastChance{30, FoggyWeather},
		weatherForecastChance{40, CloudyWeather},
		weatherForecastChance{70, FairWeather},
	)

	m["north shroud"] = createWeatherForecast(
		ClearWeather,
		weatherForecastChance{5, FoggyWeather},
		weatherForecastChance{10, ShowerWeather},
		weatherForecastChance{25, RainyWeather},
		weatherForecastChance{30, FoggyWeather},
		weatherForecastChance{40, CloudyWeather},
		weatherForecastChance{70, FairWeather},
	)

	m["the lavender beds"] = createWeatherForecast(
		FairWeather,
		weatherForecastChance{5, CloudyWeather},
		weatherForecastChance{20, RainyWeather},
		weatherForecastChance{30, FoggyWeather},
		weatherForecastChance{40, CloudyWeather},
		weatherForecastChance{55, FairWeather},
		weatherForecastChance{85, ClearWeather},
	)

	m["ul'dah"] = createWeatherForecast(
		RainyWeather,
		weatherForecastChance{40, ClearWeather},
		weatherForecastChance{60, FairWeather},
		weatherForecastChance{85, CloudyWeather},
		weatherForecastChance{95, FoggyWeather},
	)

	m["western thanalan"] = createWeatherForecast(
		RainyWeather,
		weatherForecastChance{40, ClearWeather},
		weatherForecastChance{60, FairWeather},
		weatherForecastChance{85, CloudyWeather},
		weatherForecastChance{95, FoggyWeather},
	)

	m["central thanalan"] = createWeatherForecast(
		RainyWeather,
		weatherForecastChance{15, DustStormWeather},
		weatherForecastChance{55, ClearWeather},
		weatherForecastChance{75, FairWeather},
		weatherForecastChance{85, CloudyWeather},
		weatherForecastChance{95, FoggyWeather},
	)

	m["eastern thanalan"] = createWeatherForecast(
		ShowerWeather,
		weatherForecastChance{40, ClearWeather},
		weatherForecastChance{60, FairWeather},
		weatherForecastChance{70, CloudyWeather},
		weatherForecastChance{80, FoggyWeather},
		weatherForecastChance{85, RainyWeather},
	)

	m["southern thanalan"] = createWeatherForecast(
		FoggyWeather,
		weatherForecastChance{20, HeatWaveWeather},
		weatherForecastChance{60, ClearWeather},
		weatherForecastChance{80, FairWeather},
		weatherForecastChance{90, CloudyWeather},
	)

	m["northern thanalan"] = createWeatherForecast(
		FoggyWeather,
		weatherForecastChance{5, ClearWeather},
		weatherForecastChance{20, FairWeather},
		weatherForecastChance{50, CloudyWeather},
	)

	m["the goblet"] = createWeatherForecast(
		RainyWeather,
		weatherForecastChance{40, ClearWeather},
		weatherForecastChance{60, FairWeather},
		weatherForecastChance{85, CloudyWeather},
		weatherForecastChance{95, FoggyWeather},
	)

	m["mor dhona"] = createWeatherForecast(
		FairWeather,
		weatherForecastChance{15, CloudyWeather},
		weatherForecastChance{30, FoggyWeather},
		weatherForecastChance{60, GloomyWeather},
		weatherForecastChance{75, ClearWeather},
	)

	m["ishgard"] = createWeatherForecast(
		FoggyWeather,
		weatherForecastChance{60, SnowyWeather},
		weatherForecastChance{70, FairWeather},
		weatherForecastChance{75, ClearWeather},
		weatherForecastChance{90, CloudyWeather},
	)

	m["coerthas central highlands"] = createWeatherForecast(
		FoggyWeather,
		weatherForecastChance{20, BlizzardWeather},
		weatherForecastChance{60, SnowyWeather},
		weatherForecastChance{70, FairWeather},
		weatherForecastChance{75, ClearWeather},
		weatherForecastChance{90, CloudyWeather},
	)

	m["coerthas western highlands"] = createWeatherForecast(
		FoggyWeather,
		weatherForecastChance{20, BlizzardWeather},
		weatherForecastChance{60, SnowyWeather},
		weatherForecastChance{70, FairWeather},
		weatherForecastChance{75, ClearWeather},
		weatherForecastChance{90, CloudyWeather},
	)

	m["the sea of clouds"] = createWeatherForecast(
		UmbralWindWeather,
		weatherForecastChance{30, ClearWeather},
		weatherForecastChance{60, FairWeather},
		weatherForecastChance{70, CloudyWeather},
		weatherForecastChance{80, FoggyWeather},
		weatherForecastChance{90, WindyWeather},
	)

	m["azys lla"] = createWeatherForecast(
		ThunderWeather,
		weatherForecastChance{35, FairWeather},
		weatherForecastChance{70, CloudyWeather},
	)

	m["the dravanian forelands"] = createWeatherForecast(
		FairWeather,
		weatherForecastChance{10, CloudyWeather},
		weatherForecastChance{20, FoggyWeather},
		weatherForecastChance{30, ThunderWeather},
		weatherForecastChance{40, DustStormWeather},
		weatherForecastChance{70, ClearWeather},
	)

	m["the dravanian hinterlands"] = createWeatherForecast(
		FairWeather,
		weatherForecastChance{10, CloudyWeather},
		weatherForecastChance{20, FoggyWeather},
		weatherForecastChance{30, RainyWeather},
		weatherForecastChance{40, ShowerWeather},
		weatherForecastChance{70, ClearWeather},
	)

	m["the churning mists"] = createWeatherForecast(
		FairWeather,
		weatherForecastChance{10, CloudyWeather},
		weatherForecastChance{20, GalesWeather},
		weatherForecastChance{40, UmbralStaticWeather},
		weatherForecastChance{70, ClearWeather},
	)

	m["idyllshire"] = createWeatherForecast(
		FairWeather,
		weatherForecastChance{10, CloudyWeather},
		weatherForecastChance{20, FoggyWeather},
		weatherForecastChance{30, RainyWeather},
		weatherForecastChance{40, ShowerWeather},
		weatherForecastChance{70, ClearWeather},
	)
	return m
}

func weatherChance(timestamp time.Time) int {
	inc := (timestamp.Hour() + 8) % 8

	// Take Eorzea days since unix epoch
	days := int64((timestamp.Unix() / 4200) << 32)

	var calcBase = int64(inc) + days*100
	var step1 = int64((calcBase << 11) ^ calcBase)
	var step2 = int64((step1 >> 8) ^ step1)
	return int(step2 % 100)
}

// GetCurrentWeather Get the current weather for the specified zone
func GetCurrentWeather(zone string) (Weather, error) {
	return GetWeather(Time(), zone)
}

// GetWeather Get the weather at `timestamp` for `zone`
func GetWeather(timestamp time.Time, zone string) (Weather, error) {
	_, ok := zoneMap[zone]
	for ok {
		zone = zoneMap[zone]
		_, ok = zoneMap[zone]
	}

	if _, ok = weatherMap[zone]; !ok {
		return 0, fmt.Errorf("zone `%v` not found", zone)
	}

	chance := weatherChance(timestamp)
	return weatherMap[zone](chance), nil
}

func createWeatherForecast(defaultWeather Weather, weathers ...weatherForecastChance) func(int) Weather {
	return func(chance int) Weather {
		for _, forecastChance := range weathers {
			if chance < forecastChance.chance {
				return forecastChance.weather
			}
		}

		return defaultWeather
	}
}
