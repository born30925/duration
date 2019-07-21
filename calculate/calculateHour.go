package calculate

import "time"

func CalcualteHour(StartDate, EndDate time.Time) int {
	data := EndDate.Sub(StartDate).Hours()
	return int(data) + 24
}
