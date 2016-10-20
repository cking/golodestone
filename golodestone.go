package golodestone

import (
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
