package main

import (
	"fmt"
	"time"

	"github.com/ysluckly/chinese-holiday/holidays"
)

func main() {
	// 距离当前日期的第N工作日
	result, _ := holidays.GetSNthWorkingDay("2022-10-06 05:00:00", 3)
	// 是否是节假日
	isHoliday, _ := holidays.IsHoliday(time.Now())
	// 是否是工作日
	isWorkDay, _ := holidays.IsWorkingDay(time.Now().Add(6 * 24 * time.Hour))

	fmt.Println(result, isHoliday, isWorkDay)
	return
}
