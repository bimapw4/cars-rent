package common

import "time"

func GetWeekNumber(startDate, paymentDate time.Time) int {

	start := startDate.Truncate(24 * time.Hour)
	payment := paymentDate.Truncate(24 * time.Hour)

	days := int(payment.Sub(start).Hours() / 24)

	week := (days / 7) + 1

	if week < 1 {
		return 1
	}
	return week
}
