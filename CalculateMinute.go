package main

import (
	"fmt"
	"time"
)

func main() {
	EndDate := time.Date(2019, 7, 21, 0, 0, 0, 0, time.UTC)

	fmt.Println(calcualteMinutes(time.Date(1995, 11, 25, 0, 0, 0, 0, time.UTC),EndDate))
	fmt.Println(calcualteMinutes(time.Date(1995, 06, 18, 0, 0, 0, 0, time.UTC),EndDate))
	fmt.Println(calcualteMinutes(time.Date(1995, 06, 15, 0, 0, 0, 0, time.UTC),EndDate))
	fmt.Println(calcualteMinutes(time.Date(1995, 10, 17, 0, 0, 0, 0, time.UTC),EndDate))
	fmt.Println(calcualteMinutes(time.Date(1995, 10, 3, 0, 0, 0, 0, time.UTC),EndDate))

}

func calcualteMinutes(StartDate, EndDate time.Time) int {
	data := EndDate.Sub(StartDate).Minutes()
	return int(data) + (1440)
	//return 1
}
