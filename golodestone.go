package golodestone

import (
	"math"
	"time"

	"github.com/cking/x/xmath"
)

const eorzeaMultiplier float64 = 3600.0 / 175

// Time Gets the time from Eorzea
func Time() time.Time {
	nowTicks := time.Now().UTC().Unix()
	eorzeaTicks := xmath.Round(eorzeaMultiplier * float64(nowTicks))
	return time.Unix(eorzeaTicks, 0).UTC()
}

const weeklyResetEvery = 7 * 24 * 60 * 60
const weeklyResetOffset = (5*24 + 8) * 60 * 60

// GetWeeklyReset Gets the time of the next weekly reset.
func GetWeeklyReset() time.Time {
	return GetWeeklyResetFrom(time.Now())
}

// GetWeeklyResetFrom Gets the time of the next weekly reset based on `from` time.Time
func GetWeeklyResetFrom(from time.Time) time.Time {
	return calculateResetTime(from, weeklyResetEvery, weeklyResetOffset)
}

const dailyResetEvery = 24 * 60 * 60
const dailyResetOffset = 15 * 60 * 60

// GetDailyReset Gets the time of the next daily reset.
func GetDailyReset() time.Time {
	return GetDailyResetFrom(time.Now())
}

// GetDailyResetFrom Gets the time of the next daily reset based on `from` time.Time
func GetDailyResetFrom(from time.Time) time.Time {
	return calculateResetTime(from, dailyResetEvery, dailyResetOffset)
}

const craftingScriptsResetEvery = 7 * 24 * 60 * 60
const craftingScriptsResetOffset = 8 * 60 * 60

// GetCraftingScripsReset Gets the time of the next crafting scrips reset.
func GetCraftingScripsReset() time.Time {
	return GetCraftingScripsResetFrom(time.Now())
}

// GetCraftingScripsResetFrom Gets the time of the next crafting scrips reset based on `from` time.Time
func GetCraftingScripsResetFrom(from time.Time) time.Time {
	return calculateResetTime(from, craftingScriptsResetEvery, craftingScriptsResetOffset)
}

const grandCompanyResetEvery = 24 * 60 * 60
const grandCompanyResetOffset = 20 * 60 * 60

// GetGrandCompanyReset Gets the time of the next grand company reset.
func GetGrandCompanyReset() time.Time {
	return GetGrandCompanyResetFrom(time.Now())
}

// GetGrandCompanyResetFrom Gets the time of the next grand company reset based on `from` time.Time
func GetGrandCompanyResetFrom(from time.Time) time.Time {
	return calculateResetTime(from, grandCompanyResetEvery, grandCompanyResetOffset)
}

func calculateResetTime(now time.Time, every, offset int64) time.Time {
	end := int64(math.Floor(float64((now.UTC().Unix()-offset)/every))+1)*every + offset
	return time.Unix(end, 0)
}
