package common

import (
	"math"
	"time"
)

func CalculateRentalCost(pickupDate, dropoffDate time.Time, dayRate, monthRate float64) float64 {
	duration := dropoffDate.Sub(pickupDate).Hours() / 24
	totalDays := int(math.Ceil(duration))

	months := totalDays / 30
	remainingDays := totalDays % 30

	return float64(months)*monthRate + float64(remainingDays)*dayRate
}
