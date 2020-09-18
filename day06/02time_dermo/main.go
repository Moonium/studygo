package main

import (
	"fmt"
	"time"
)

func f1() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Date())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	// 时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	// tine.Unix()
	ret := time.Unix(1599790183, 0)
	fmt.Println(ret)
	fmt.Println(ret.Year())
	fmt.Println(ret.Date())
	// 时间间隔
	fmt.Println(time.Second)
	fmt.Println(now.Add(24 * time.Hour))
	nextYear, err := time.Parse("2006-01-02 15:04:05", "2021-07-18 21:31:12")
	if err != nil {
		fmt.Printf("Parse time failed, err: %v", err)
		return
	}
	now = now.UTC()
	nextYear = nextYear.UTC()
	d := nextYear.Sub(now)
	fmt.Println(d)
	fmt.Println("------------")

	// // 定时器
	// timer := time.Tick(time.Second)
	// for t := range timer {
	// 	fmt.Println(t)
	// }

	// 格式化，将时间转换成字符串类型
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006/01/02 03:04:05 PM"))
	fmt.Println(now.Format("2006/01/02 15:04:05.000"))

	// 解析对应的格式将一个字符串类型的时间转换成时间戳
	timeObj, err := time.Parse("2006-01-02", "1995-07-18")
	if err != nil {
		fmt.Printf("Parse time failed, err: %v", err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())

	// Sleep
	n := 100
	fmt.Println("开始睡眠~")
	time.Sleep(time.Duration(n))
	time.Sleep(5 * time.Second)
	fmt.Println("5秒钟过去了！")
}

// 时区
func f2() {
	now := time.Now() // 获取的是本地的时间
	fmt.Println(now)
	// 明天的这个时间
	time.Parse("2006-01-02 15:04:05", "2020-09-11 11:06:10")
	// 按照东八区的时区和格式解析字符串格式的时间
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("load location failed, err: %v", err)
		return
	}
	// 按照指定时区解析时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-09-12 11:06:10", loc)
	if err != nil {
		fmt.Printf("parse time failed, err: %v", err)
		return
	}
	fmt.Println(timeObj)
	// 时间对象相减
	td := timeObj.Sub(now)
	fmt.Println(td)
}

func main() {
	// f1()
	f2()
}
