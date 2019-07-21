package calculate

import (
	"fmt"
	"testing"
	"time"
)

func Test_calculate_hour_startDate_15_06_1996_endDate_21_07_1996_Should_be_202488(t *testing.T) {

	expected := 202488

	startDate_temp := "1996-06-15T00:00:00.000Z"
	endDate_temp := "2019-07-21T00:00:00.000Z"

	layout := "2006-01-02T15:04:05.000Z"

	startDate, err1 := time.Parse(layout, startDate_temp)
	endDate, err2 := time.Parse(layout, endDate_temp)

	if err1 != nil {
		fmt.Println(err1)
	}
	if err2 != nil {
		fmt.Println(err2)
	}

	actual := CalcualteHour(startDate, endDate)
	if actual != expected {
		t.Errorf("Expected %d but it got %d", expected, actual)
	}

}
