package calculate

import (
	"testing"
	"time"
)

func Test_calculateMinutes_Input_Start_1995_11_25_End_2019_7_21_Should_Be_12441600(t *testing.T){
	expected := 12441600
	StartDate := time.Date(1995, 11, 25, 0, 0, 0, 0, time.UTC)
	EndDate := time.Date(2019, 7, 21, 0, 0, 0, 0, time.UTC)
	actual := calculateMinutes(StartDate,EndDate)
	if actual != expected {
		t.Errorf("expecred %d but it got %d", expected, actual)
	}
}

func Test_calculateMinutes_Input_Start_1995_6_18_End_2019_7_21_Should_Be_12672000(t *testing.T){
	expected := 12672000
	StartDate := time.Date(1995, 6, 18, 0, 0, 0, 0, time.UTC)
	EndDate := time.Date(2019, 7, 21, 0, 0, 0, 0, time.UTC)
	actual := calculateMinutes(StartDate,EndDate)
	if actual != expected {
		t.Errorf("expecred %d but it got %d", expected, actual)
	}
}

func Test_calculateMinutes_Input_Start_1995_6_15_End_2019_7_21_Should_Be_12676320(t *testing.T){
	expected := 12676320
	StartDate := time.Date(1995, 6, 15, 0, 0, 0, 0, time.UTC)
	EndDate := time.Date(2019, 7, 21, 0, 0, 0, 0, time.UTC)
	actual := calculateMinutes(StartDate,EndDate)
	if actual != expected {
		t.Errorf("expecred %d but it got %d", expected, actual)
	}
}

func Test_calculateMinutes_Input_Start_1995_10_17_End_2019_7_21_Should_Be_12497760(t *testing.T){
	expected := 12497760
	StartDate := time.Date(1995, 10, 17, 0, 0, 0, 0, time.UTC)
	EndDate := time.Date(2019, 7, 21, 0, 0, 0, 0, time.UTC)
	actual := calculateMinutes(StartDate,EndDate)
	if actual != expected {
		t.Errorf("expecred %d but it got %d", expected, actual)
	}
}

func Test_calculateMinutes_Input_Start_1995_10_3_End_2019_7_21_Should_Be_12517920(t *testing.T){
	expected := 12517920
	StartDate := time.Date(1995, 10, 3, 0, 0, 0, 0, time.UTC)
	EndDate := time.Date(2019, 7, 21, 0, 0, 0, 0, time.UTC)
	actual := calculateMinutes(StartDate,EndDate)
	if actual != expected {
		t.Errorf("expecred %d but it got %d", expected, actual)
	}
}
