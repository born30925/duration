package calculate

import "time"

func Test_calculateSeconds_Input_startDate_25111995_endDate_21072019_Should_Be_746496000(t *test) {

	expected := 746496000
	startDate := time.Date()
	endDate := time.Date()

	actual := calculateSeconds(startDate, endDate)
	if actual != expected {
		t.Errorf("Expected %d but it got %d", expected, actual)
	}

}
