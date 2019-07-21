package calculate

import (
	"testing"
	"time"
)

func Test_string_to_date_Input_25_11_1995(t *testing.T) {
	expected, _ := time.Parse("02-01-2006", "25-11-1995")
	stringDate := "25-11-1995"
	actual := stringToDateTime(stringDate)
	if actual != expected {
		t.Errorf("Expected %s but it got %s", expected, actual)
	}
}

func Test_string_to_date_Input_18_06_1995(t *testing.T) {
	expected, _ := time.Parse("02-01-2006", "18-06-1995")
	stringDate := "18-06-1995"
	actual := stringToDateTime(stringDate)
	if actual != expected {
		t.Errorf("Expected %s but it got %s", expected, actual)
	}
}

func Test_string_to_date_Input_15_06_1995(t *testing.T) {
	expected, _ := time.Parse("02-01-2006", "15-06-1995")
	stringDate := "15-06-1995"
	actual := stringToDateTime(stringDate)
	if actual != expected {
		t.Errorf("Expected %s but it got %s", expected, actual)
	}
}

func Test_string_to_date_Input_17_10_1995(t *testing.T) {
	expected, _ := time.Parse("02-01-2006", "17-10-1995")
	stringDate := "17-10-1995"
	actual := stringToDateTime(stringDate)
	if actual != expected {
		t.Errorf("Expected %s but it got %s", expected, actual)
	}
}

func Test_string_to_date_Input_03_10_1995(t *testing.T) {
	expected, _ := time.Parse("02-01-2006", "03-10-1995")
	stringDate := "03-10-1995"
	actual := stringToDateTime(stringDate)
	if actual != expected {
		t.Errorf("Expected %s but it got %s", expected, actual)
	}
}

func Test_string_to_date_Input_21_07_2019(t *testing.T) {
	expected, _ := time.Parse("02-01-2006", "21-07-2019")
	stringDate := "21-07-2019"
	actual := stringToDateTime(stringDate)
	if actual != expected {
		t.Errorf("Expected %s but it got %s", expected, actual)
	}
}
