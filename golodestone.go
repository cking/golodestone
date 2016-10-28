package golodestone

import (
	"time"

	"github.com/cking/x/xmath"
)

const eorzeaMultiplier float64 = 3600.0 / 175
const day = 60 * 60 * 24

type resetDefinition struct {
	offset   time.Duration
	stepSize int64
}

// Resets holds all reset times
type Resets struct {
	Daily        time.Time
	GrandCompany time.Time
	Weekly       time.Time
	Scrip        time.Time
}

var dailyReset = resetDefinition{time.Hour * 15, day}
var grandCompanyReset = resetDefinition{time.Hour * 20, day}
var weeklyReset = resetDefinition{time.Hour * (5*24 + 8), day * 7}
var scripReset = resetDefinition{time.Hour * 8, day * 7}

// Time Gets the time from Eorzea
func Time() time.Time {
	nowTicks := time.Now().UTC().Unix()
	eorzeaTicks := xmath.Round(eorzeaMultiplier * float64(nowTicks))
	return time.Unix(eorzeaTicks, 0).UTC()
}

// GetResets returns all resets for the current time
func GetResets() *Resets {
	return GetResetsFrom(time.Now())
}

// GetResetsFrom returns all resets for the given time
func GetResetsFrom(from time.Time) *Resets {
	return &Resets{calculateResetTime(from, dailyReset), calculateResetTime(from, grandCompanyReset), calculateResetTime(from, weeklyReset), calculateResetTime(from, scripReset)}
}

func calculateResetTime(now time.Time, reset resetDefinition) time.Time {
	base := now.UTC().Add(-reset.offset).Unix()
	nextStep := base/reset.stepSize + 1
	resetTime := nextStep * reset.stepSize
	return time.Unix(resetTime, 0).Add(reset.offset)
}
