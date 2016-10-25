package golodestone

import (
	"time"

	"github.com/cking/x/xmath"
)

const eorzeaMultiplier float64 = 3600.0 / 175

// Time Gets the time from Eorzea
func Time() time.Time {
	return TimeFrom(time.Now())
}

// TimeFrom Gets the time from Eorzea for `from` time.Time
func TimeFrom(from time.Time) time.Time {
	fromTicks := from.UTC().Unix()
	eorzeaTicks := xmath.Round(eorzeaMultiplier * float64(fromTicks))
	return time.Unix(eorzeaTicks, 0).UTC()
}
