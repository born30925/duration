package calculate

import (
	"testing"
	"time"
)

func Test_calculateSeconds_Input_startDate_1995_11_25_endDate_2019_7_21_Should_Be_746496000(t *testing.T) {

	expected := 746496000
	startDate := time.Date(1995, 11, 25, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2019, 7, 21, 0, 0, 0, 0, time.UTC)

	actual := calculateSeconds(startDate, endDate)
	if actual != expected {
		t.Errorf("Expected %d but it got %d", expected, actual)
	}

}

func Test_calculateSeconds_Input_startDate_1995_06_18_endDate_2019_7_21_Should_Be_760320000(t *testing.T) {

	expected := 760320000
	startDate := time.Date(1995, 6, 18, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2019, 7, 21, 0, 0, 0, 0, time.UTC)

	actual := calculateSeconds(startDate, endDate)
	if actual != expected {
		t.Errorf("Expected %d but it got %d", expected, actual)
	}

}

func Test_calculateSeconds_Input_startDate_1995_06_15_endDate_2019_7_21_Should_Be_760579200(t *testing.T) {

	expected := 760579200
	startDate := time.Date(1995, 6, 15, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2019, 7, 21, 0, 0, 0, 0, time.UTC)

	actual := calculateSeconds(startDate, endDate)
	if actual != expected {
		t.Errorf("Expected %d but it got %d", expected, actual)
	}

}

func Test_calculateSeconds_Input_startDate_1995_10_17_endDate_2019_7_21_Should_Be_749865600(t *testing.T) {

	expected := 749865600
	startDate := time.Date(1995, 10, 17, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2019, 7, 21, 0, 0, 0, 0, time.UTC)

	actual := calculateSeconds(startDate, endDate)
	if actual != expected {
		t.Errorf("Expected %d but it got %d", expected, actual)
	}

}

func Test_calculateSeconds_Input_startDate_1995_10_03_endDate_2019_7_21_Should_Be_751075200(t *testing.T) {

	expected := 751075200
	startDate := time.Date(1995, 10, 3, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2019, 7, 21, 0, 0, 0, 0, time.UTC)

	actual := calculateSeconds(startDate, endDate)
	if actual != expected {
		t.Errorf("Expected %d but it got %d", expected, actual)
	}

}