package calculate

import (
	"testing"
	"time"
)

func Test_calculateSeconds_Input_startDate_1995_11_19_endDate_2019_7_21_Should_Be_746496000(t *testing.T) {

	expected := 746496000
	startDate := time.Date(1995, 11, 19, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2019, 7, 21, 0, 0, 0, 0, time.UTC)

	actual := calculateSeconds(startDate, endDate)
	if actual != expected {
		t.Errorf("Expected %d but it got %d", expected, actual)
	}

}
