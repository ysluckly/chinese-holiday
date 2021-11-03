package holidays

import (
	"fmt"
	"testing"
	"time"
)

// TestIsHoliday
func TestIsHoliday(t *testing.T) {
	d := time.Date(2019, 10, 1, 0, 0, 0, 0, location)
	result, err := IsHoliday(d)
	if err != nil {
		t.Error(err)
	}

	if !result {
		t.Fail()
	}
}

// TestIsWorkingday
func TestIsWorkingday(t *testing.T) {
	d := time.Date(2019, 9, 30, 0, 0, 0, 0, location)
	result, err := IsWorkingDay(d)
	if err != nil {
		t.Error(err)
	}

	if !result {
		t.Fail()
	}
}

// TestGetNthWorkingDay
func TestGetNthWorkingDay(t *testing.T) {
	d := time.Date(2022, 10, 9, 5, 0, 0, 0, location)
	t.Run("TestGetNthWorkingDay", func(t *testing.T) {
		result, err := GetTNthWorkingDay(d, 3)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(result)
	})
}

// GetSNthWorkingDay
func TestGetSNthWorkingDay(t *testing.T) {
	d := "2022-01-31 12:12:12"
	t.Run("GetSNthWorkingDay", func(t *testing.T) {
		result, err := GetSNthWorkingDay(d, 3)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(result)
	})
}
