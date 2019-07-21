package calculate

import "time"

func calculateMinutes(StartDate, EndDate time.Time) int {
	data := EndDate.Sub(StartDate).Minutes()
	return int(data) + (1440)
}
