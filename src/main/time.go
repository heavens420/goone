package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	date := time.Date(2020, 03, 04, 23, 3, 45, 22, time.Local)
	fmt.Println(date)

	//format 日期固定 否则不准
	f1 := now.Format("2006-1-2 15:04:05")
	fmt.Println(f1)
	f2 := date.Format("2006-1-2")
	fmt.Println(f2)

	s := "2020-3-4 23:32:55"
	p1, err := time.Parse("2006-1-2 15:04:05", s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p1)
	fmt.Printf("%T\n", p1)

	//获取年月日 时分秒
	year, month, day := now.Date()
	fmt.Println(year, month, day)
	clock, min, sec := now.Clock()
	fmt.Println(clock, min, sec)

	fmt.Println(now.Year(), now.Month(), now.Day())

	//一年中的第几天
	fmt.Println(now.YearDay())

	//今天周几
	fmt.Println(now.Weekday())

	fmt.Println(now.ISOWeek())
}
