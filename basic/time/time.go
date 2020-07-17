package main

import (
	"fmt"
	"time"
)

func main() {
	// 自定义时区
	// China doesn't have daylight saving. It uses a fixed 8 hour offset from UTC.
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijingLocation := time.FixedZone("Beijing Time", secondsEastOfUTC)

	t, _ := time.ParseInLocation("20060102", "20200420", beijingLocation)
	fmt.Println(t)

	// 小时 分钟 秒 使用 Add
	for i := 1; i < 60; i++ {
		t1 := t.AddDate(0, 0, i) // 天以上的
		fmt.Println(t1)
		fmt.Println(t1.Format("20060102"))
	}

	if "20200422" > "20200420" {
		fmt.Println("---------")
	}

	fmt.Println(time.Now().Format(time.RFC3339))

	tt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	fmt.Println(tt.Format("20060102"))
}
