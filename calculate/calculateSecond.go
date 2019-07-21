package calculate

import "time"

func calculateSeconds(startDate, endDate time.Time) int {
	result := endDate.Sub(startDate).Seconds()
	return int(result) + (1440 * 60)
}
