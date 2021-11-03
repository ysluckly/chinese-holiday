package tool

import "time"

func GetBeforeDawnOfDate(date time.Time) time.Time {
	local, _ := time.LoadLocation("Local")
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, local)
}
