package main

import (
	"fmt"

	"github.com/ysliving/chinese-holidays/holidays"
)

// 测试
func main() {
	result, err := holidays.GetSNthWorkingDay("2022-09-30 00:00:00", 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
