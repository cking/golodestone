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
	CloudyWeather
	// FairWeather Fair weather
	FairWeather
	// FoggyWeather Foggy weather
	FoggyWeather
	// RainyWeather Rain
	RainyWeather
	// WindyWeather Windy weather
	WindyWeather
	// ShowerWeather Rain shower
	ShowerWeather
	// GalesWeather Gale
	GalesWeather
	// ThunderWeather Thunder
	ThunderWeather
	// ThunderstormWeather Thunderstorm
	ThunderstormWeather
	// HeatWaveWeather Heat waves
	HeatWaveWeather
	// GloomyWeather Gloomy weather
	GloomyWeather
	// SnowyWeather Snow
	SnowyWeather
	// BlizzardWeather Blizzards
	BlizzardWeather
	// UmbralWindWeather Umbral winds
	UmbralWindWeather
	// DustStormWeather Dust storms
	DustStormWeather
	// UmbralStaticWeather Umbral statics
	UmbralStaticWeather
)

// Zone Type
type Zone int

const (
	// UpperLaNosceaZone Upper La Noscea
	UpperLaNosceaZone Zone = iota
	// GridaniaZone Gridania
	GridaniaZone
	// CentralShroudZone Central Shroud
	CentralShroudZone
	// CoerthasWesternHighlandsZone Coerthas Western Highlands
	CoerthasWesternHighlandsZone
	// MistZone Mist
	MistZone
	// IshgardZone Ishgard
	IshgardZone
	// CoerthasCentralHighlandsZone Coerthas Western Highlands
	CoerthasCentralHighlandsZone
	// MorDhonaZone Mor Dhona
	MorDhonaZone
	// EasternThanalanZone Eastern Thanalan
	EasternThanalanZone
	// TheDravanianForelandsZone The Dravanian Forelands
	TheDravanianForelandsZone
	// LowerLaNosceaZone Lower La Noscea
	LowerLaNosceaZone
	// EasternLaNosceaZone Eastern La Noscea
	EasternLaNosceaZone
	// NorthShroudZone North Shroud Zone
	NorthShroudZone
	// WesternThanalanZone Western Thanalan
	WesternThanalanZone
	// EastShroudZone East Shroud
	EastShroudZone
	// CentralThanalanZone Central Thanalan
	CentralThanalanZone
	// AzysLlaZone Azys Lla
	AzysLlaZone
	// LimsaLominsaZone Limsa Lominsa
	LimsaLominsaZone
	// WesternLaNosceaZone Western La Noscea
	WesternLaNosceaZone
	// TheDravanianHinterlandsZone The Dravanian Hinterlands
	TheDravanianHinterlandsZone
	// IdyllshireZone Idyllshire
	IdyllshireZone
	// TheLavenderBedsZone The Lavender Beds
	TheLavenderBedsZone
	// UldahZone Ul'dah
	UldahZone
	// SouthShroudZone South Shroud
	SouthShroudZone
	// TheSeaofCloudsZone The Sea of Clouds
	TheSeaofCloudsZone
	// NorthernThanalanZone Northern Thanalan
	NorthernThanalanZone
	// SouthernThanalanZone Southern Thanalan
	SouthernThanalanZone
	// TheChurningMistsZone The Churning Mists
	TheChurningMistsZone
	// TheGobletZone The Goblet
	TheGobletZone
	// OuterLaNosceaZone Outer La Noscea
	OuterLaNosceaZone
	// MiddleLaNosceaZone Middle La Noscea
	MiddleLaNosceaZone
)

var (
	weatherStringMap, stringWeatherMap = weatherTypeMap()
	zoneStringMap, stringZoneMap       = weatherZoneMap()
	weatherMap                         = weatherChanceMap()
)

func weatherTypeMap() (map[Weather]string, map[string]Weather) {
	m := make(map[Weather]string)

	m[ClearWeather] = "Clear weather"
	m[CloudyWeather] = "Cloudy weather"
	m[FairWeather] = "Fair weather"
	m[FoggyWeather] = "Foggy weather"
	m[RainyWeather] = "Rain"
	m[WindyWeather] = "Windy weather"
	m[ShowerWeather] = "Rain shower"
	m[GalesWeather] = "Gale"
	m[ThunderWeather] = "Thunder"
	m[ThunderstormWeather] = "Thunderstorm"
	m[HeatWaveWeather] = "Heat waves"
	m[GloomyWeather] = "Gloomy weather"
	m[SnowyWeather] = "Snow"
	m[BlizzardWeather] = "Blizzards"
	m[UmbralWindWeather] = "Umbral winds"
	m[DustStormWeather] = "Dust storms"
	m[UmbralStaticWeather] = "Umbral statics"

	n := make(map[string]Weather, len(m))
	for k, v := range m {
		n[v] = k
	}

	return m, n
}

func weatherZoneMap() (map[Zone]string, map[string]Zone) {
	m := make(map[Zone]string)

	m[UpperLaNosceaZone] = "Upper La Noscea"
	m[GridaniaZone] = "Gridania"
	m[CentralShroudZone] = "Central Shroud"
	m[CoerthasWesternHighlandsZone] = "Coerthas Western Highlands"
	m[MistZone] = "Mist"
	m[IshgardZone] = "Ishgard"
	m[CoerthasCentralHighlandsZone] = "Coerthas Central Highlands"
	m[MorDhonaZone] = "Mor Dhona"
	m[EasternThanalanZone] = "Eastern Thanalan"
	m[TheDravanianForelandsZone] = "The Dravanian Forelands"
	m[LowerLaNosceaZone] = "Lower La Noscea"
	m[EasternLaNosceaZone] = "Eastern La Noscea"
	m[NorthShroudZone] = "North Shroud"
	m[WesternThanalanZone] = "Western Thanalan"
	m[EastShroudZone] = "East Shroud"
	m[CentralThanalanZone] = "Central Thanalan"
	m[AzysLlaZone] = "Azys Lla"
	m[LimsaLominsaZone] = "Limsa Lominsa"
	m[WesternLaNosceaZone] = "Western La Noscea"
	m[TheDravanianHinterlandsZone] = "The Dravanian Hinterlands"
	m[IdyllshireZone] = "Idyllshire"
	m[TheLavenderBedsZone] = "The Lavender Beds"
	m[UldahZone] = "Ul'dah"
	m[SouthShroudZone] = "South Shroud"
	m[TheSeaofCloudsZone] = "The Sea of Clouds"
	m[NorthernThanalanZone] = "Northern Thanalan"
	m[SouthernThanalanZone] = "Southern Thanalan"
	m[TheChurningMistsZone] = "The Churning Mists"
	m[TheGobletZone] = "The Goblet"
	m[OuterLaNosceaZone] = "Outer La Noscea"
	m[MiddleLaNosceaZone] = "Middle La Noscea"

	n := make(map[string]Zone, len(m))
	for k, v := range m {
		n[v] = k
	}

	n["uldah"] = UldahZone
	n["cch"] = CoerthasCentralHighlandsZone

	return m, n
}

func weatherChanceMap() map[Zone]weatherChanceFunction {
	m := make(map[Zone]weatherChanceFunction)
	m[LimsaLominsaZone] = createWeatherForecast(
		RainyWeather,
		weatherForecastChance{20, CloudyWeather},
		weatherForecastChance{50, ClearWeather},
		weatherForecastChance{80, FairWeather},
		weatherForecastChance{90, FoggyWeather},
	)

	m[MiddleLaNosceaZone] = createWeatherForecast(
		RainyWeather,
		weatherForecastChance{20, CloudyWeather},
		weatherForecastChance{50, ClearWeather},
		weatherForecastChance{70, FairWeather},
		weatherForecastChance{80, WindyWeather},
		weatherForecastChance{90, FoggyWeather},
	)

	m[LowerLaNosceaZone] = createWeatherForecast(
		RainyWeather,
		weatherForecastChance{20, CloudyWeather},
		weatherForecastChance{50, ClearWeather},
		weatherForecastChance{70, FairWeather},
		weatherForecastChance{80, WindyWeather},
		weatherForecastChance{90, FoggyWeather},
	)

	m[EasternLaNosceaZone] = createWeatherForecast(
		ShowerWeather,
		weatherForecastChance{5, FoggyWeather},
		weatherForecastChance{50, ClearWeather},
		weatherForecastChance{80, FairWeather},
		weatherForecastChance{90, CloudyWeather},
		weatherForecastChance{95, RainyWeather},
	)

	m[WesternLaNosceaZone] = createWeatherForecast(
		GalesWeather,
		weatherForecastChance{10, FoggyWeather},
		weatherForecastChance{40, ClearWeather},
		weatherForecastChance{60, FairWeather},
		weatherForecastChance{80, CloudyWeather},
		weatherForecastChance{90, WindyWeather},
	)

	m[UpperLaNosceaZone] = createWeatherForecast(
		ThunderstormWeather,
		weatherForecastChance{30, ClearWeather},
		weatherForecastChance{50, FairWeather},
		weatherForecastChance{70, CloudyWeather},
		weatherForecastChance{80, FoggyWeather},
		weatherForecastChance{90, ThunderWeather},
	)

	m[OuterLaNosceaZone] = createWeatherForecast(
		RainyWeather,
		weatherForecastChance{30, ClearWeather},
		weatherForecastChance{50, FairWeather},
		weatherForecastChance{70, CloudyWeather},
		weatherForecastChance{85, FoggyWeather},
	)

	m[MistZone] = createWeatherForecast(
		RainyWeather,
		weatherForecastChance{20, CloudyWeather},
		weatherForecastChance{50, ClearWeather},
		weatherForecastChance{80, FairWeather},
		weatherForecastChance{90, FoggyWeather},
	)

	m[GridaniaZone] = createWeatherForecast(
		FairWeather,
		weatherForecastChance{20, RainyWeather},
		weatherForecastChance{30, FoggyWeather},
		weatherForecastChance{40, CloudyWeather},
		weatherForecastChance{55, FairWeather},
		weatherForecastChance{85, ClearWeather},
	)

	m[CentralShroudZone] = createWeatherForecast(
		FairWeather,
		weatherForecastChance{5, ThunderWeather},
		weatherForecastChance{20, RainyWeather},
		weatherForecastChance{30, FoggyWeather},
		weatherForecastChance{40, CloudyWeather},
		weatherForecastChance{55, FairWeather},
		weatherForecastChance{85, ClearWeather},
	)

	m[EastShroudZone] = createWeatherForecast(
		FairWeather,
		weatherForecastChance{5, ThunderWeather},
		weatherForecastChance{20, RainyWeather},
		weatherForecastChance{30, FoggyWeather},
		weatherForecastChance{40, CloudyWeather},
		weatherForecastChance{55, FairWeather},
		weatherForecastChance{85, ClearWeather},
	)

	m[SouthShroudZone] = createWeatherForecast(
		ClearWeather,
		weatherForecastChance{5, FoggyWeather},
		weatherForecastChance{10, ThunderstormWeather},
		weatherForecastChance{25, ThunderWeather},
		weatherForecastChance{30, FoggyWeather},
		weatherForecastChance{40, CloudyWeather},
		weatherForecastChance{70, FairWeather},
	)

	m[NorthShroudZone] = createWeatherForecast(
		ClearWeather,
		weatherForecastChance{5, FoggyWeather},
		weatherForecastChance{10, ShowerWeather},
		weatherForecastChance{25, RainyWeather},
		weatherForecastChance{30, FoggyWeather},
		weatherForecastChance{40, CloudyWeather},
		weatherForecastChance{70, FairWeather},
	)

	m[TheLavenderBedsZone] = createWeatherForecast(
		FairWeather,
		weatherForecastChance{5, CloudyWeather},
		weatherForecastChance{20, RainyWeather},
		weatherForecastChance{30, FoggyWeather},
		weatherForecastChance{40, CloudyWeather},
		weatherForecastChance{55, FairWeather},
		weatherForecastChance{85, ClearWeather},
	)

	m[UldahZone] = createWeatherForecast(
		RainyWeather,
		weatherForecastChance{40, ClearWeather},
		weatherForecastChance{60, FairWeather},
		weatherForecastChance{85, CloudyWeather},
		weatherForecastChance{95, FoggyWeather},
	)

	m[WesternThanalanZone] = createWeatherForecast(
		RainyWeather,
		weatherForecastChance{40, ClearWeather},
		weatherForecastChance{60, FairWeather},
		weatherForecastChance{85, CloudyWeather},
		weatherForecastChance{95, FoggyWeather},
	)

	m[CentralThanalanZone] = createWeatherForecast(
		RainyWeather,
		weatherForecastChance{15, DustStormWeather},
		weatherForecastChance{55, ClearWeather},
		weatherForecastChance{75, FairWeather},
		weatherForecastChance{85, CloudyWeather},
		weatherForecastChance{95, FoggyWeather},
	)

	m[EasternThanalanZone] = createWeatherForecast(
		ShowerWeather,
		weatherForecastChance{40, ClearWeather},
		weatherForecastChance{60, FairWeather},
		weatherForecastChance{70, CloudyWeather},
		weatherForecastChance{80, FoggyWeather},
		weatherForecastChance{85, RainyWeather},
	)

	m[SouthernThanalanZone] = createWeatherForecast(
		FoggyWeather,
		weatherForecastChance{20, HeatWaveWeather},
		weatherForecastChance{60, ClearWeather},
		weatherForecastChance{80, FairWeather},
		weatherForecastChance{90, CloudyWeather},
	)

	m[NorthernThanalanZone] = createWeatherForecast(
		FoggyWeather,
		weatherForecastChance{5, ClearWeather},
		weatherForecastChance{20, FairWeather},
		weatherForecastChance{50, CloudyWeather},
	)

	m[TheGobletZone] = createWeatherForecast(
		RainyWeather,
		weatherForecastChance{40, ClearWeather},
		weatherForecastChance{60, FairWeather},
		weatherForecastChance{85, CloudyWeather},
		weatherForecastChance{95, FoggyWeather},
	)

	m[MorDhonaZone] = createWeatherForecast(
		FairWeather,
		weatherForecastChance{15, CloudyWeather},
		weatherForecastChance{30, FoggyWeather},
		weatherForecastChance{60, GloomyWeather},
		weatherForecastChance{75, ClearWeather},
	)

	m[IshgardZone] = createWeatherForecast(
		FoggyWeather,
		weatherForecastChance{60, SnowyWeather},
		weatherForecastChance{70, FairWeather},
		weatherForecastChance{75, ClearWeather},
		weatherForecastChance{90, CloudyWeather},
	)

	m[CoerthasCentralHighlandsZone] = createWeatherForecast(
		FoggyWeather,
		weatherForecastChance{20, BlizzardWeather},
		weatherForecastChance{60, SnowyWeather},
		weatherForecastChance{70, FairWeather},
		weatherForecastChance{75, ClearWeather},
		weatherForecastChance{90, CloudyWeather},
	)

	m[CoerthasWesternHighlandsZone] = createWeatherForecast(
		FoggyWeather,
		weatherForecastChance{20, BlizzardWeather},
		weatherForecastChance{60, SnowyWeather},
		weatherForecastChance{70, FairWeather},
		weatherForecastChance{75, ClearWeather},
		weatherForecastChance{90, CloudyWeather},
	)

	m[TheSeaofCloudsZone] = createWeatherForecast(
		UmbralWindWeather,
		weatherForecastChance{30, ClearWeather},
		weatherForecastChance{60, FairWeather},
		weatherForecastChance{70, CloudyWeather},
		weatherForecastChance{80, FoggyWeather},
		weatherForecastChance{90, WindyWeather},
	)

	m[AzysLlaZone] = createWeatherForecast(
		ThunderWeather,
		weatherForecastChance{35, FairWeather},
		weatherForecastChance{70, CloudyWeather},
	)

	m[TheDravanianForelandsZone] = createWeatherForecast(
		FairWeather,
		weatherForecastChance{10, CloudyWeather},
		weatherForecastChance{20, FoggyWeather},
		weatherForecastChance{30, ThunderWeather},
		weatherForecastChance{40, DustStormWeather},
		weatherForecastChance{70, ClearWeather},
	)

	m[TheDravanianHinterlandsZone] = createWeatherForecast(
		FairWeather,
		weatherForecastChance{10, CloudyWeather},
		weatherForecastChance{20, FoggyWeather},
		weatherForecastChance{30, RainyWeather},
		weatherForecastChance{40, ShowerWeather},
		weatherForecastChance{70, ClearWeather},
	)

	m[TheChurningMistsZone] = createWeatherForecast(
		FairWeather,
		weatherForecastChance{10, CloudyWeather},
		weatherForecastChance{20, GalesWeather},
		weatherForecastChance{40, UmbralStaticWeather},
		weatherForecastChance{70, ClearWeather},
	)

	m[IdyllshireZone] = createWeatherForecast(
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

// GetCurrentWeatherByString Get the current weather for the specified `zone` where `zone` is string
func GetCurrentWeatherByString(zone string) (Weather, error) {
	id, found := stringZoneMap[zone]
	if !found {
		return 0, fmt.Errorf("zone `%v` not found", zone)
	}

	return GetWeather(Time(), id)
}

// GetCurrentWeather Get the current weather for the specified zone
func GetCurrentWeather(zone Zone) (Weather, error) {
	_, found := zoneStringMap[zone]
	if !found {
		return 0, fmt.Errorf("zone `%v` not found", zone)
	}

	return GetWeather(Time(), zone)
}

// GetWeatherByString Get the weather at `timestamp` for `zone` where `zone` is string
func GetWeatherByString(timestamp time.Time, zone string) (Weather, error) {
	id, found := stringZoneMap[zone]
	if !found {
		return 0, fmt.Errorf("zone `%v` not found", zone)
	}

	return GetWeather(timestamp, id)
}

// GetWeather Get the weather at `timestamp` for Zone
func GetWeather(timestamp time.Time, zone Zone) (Weather, error) {
	if _, ok := weatherMap[zone]; !ok {
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
