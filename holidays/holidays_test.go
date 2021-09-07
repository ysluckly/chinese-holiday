package holidays

import (
	"fmt"
	"testing"
	"time"
)

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

func TestGetNthWorkingDay(t *testing.T) {
	d := time.Date(2021, 9, 30, 0, 0, 0, 0, location)
	t.Run("TestGetNthWorkingDay", func(t *testing.T) {
		result, err := GetTNthWorkingDay(d, 3)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(result)
	})
}
