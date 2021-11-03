# chinese-holidays-go

![badge](https://github.com/bastengao/chinese-holidays-go/workflows/Go/badge.svg)

##### 提供具有中国特色的休假安排或者工作日查询。

## Install

    go get github.com/ysluckly/chinese-holiday

## Usage

```go
package main
import (
    "github.com/ysluckly/chinese-holiday/holidays"
)
func main(){
    d := time.Date(2019, 10, 1, 0, 0, 0, 0, china)
    dStr := "2019-10-01 00:00:00"
    holidays.IsHoliday(d)    // true
    holidays.IsWorkingday(d) // false
    holidays.GetTNthWorkingDay(d,3) // 距d的第三个工作日（time param）
    holidays.GetSNthWorkingDay(dStr,3) // 距d的第三个工作日（string param）   
    return
}
```

## Features
> [假期政策查询](http://www.gov.cn/zhengce/zuixin.htm)
- [x] bundled data
  - support [2022](http://www.gov.cn/zhengce/content/2021-10/25/content_5644835.htm)
  - support [2021](http://www.gov.cn/zhengce/content/2020-11/25/content_5564127.htm)
  - support [2020](http://www.gov.cn/zhengce/content/2019-11/21/content_5454164.htm)
  - support [2019](http://www.gov.cn/zhengce/content/2018-12/06/content_5346276.htm) and 5.1 [changes](http://www.gov.cn/zhengce/content/2019-03/22/content_5375877.htm)
  - support [2018](http://www.gov.cn/zhengce/content/2017-11/30/content_5243579.htm)
  - support [2017](http://www.gov.cn/zhengce/content/2016-12/01/content_5141603.htm)
  - support 2016
- [ ] online data

## Other
```c
json -> 二进制 
命令 eg： statik -src=./data(重写: -f)
```
详见：[statik命令使用](http://blog.fatedier.com/2016/08/01/compile-assets-into-binary-file-with-statik-in-golang/)
