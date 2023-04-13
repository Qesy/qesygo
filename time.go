package qesygo

import (
	"strconv"
	"time"
)

func Week() int { //获取周几 （1-7）
	t := time.Now()
	return int(t.Weekday())
}

func Time(str string) int64 {
	now := time.Now()
	t := now.UnixNano()
	switch str {
	case "Microsecond":
		t = now.UnixMicro()
	case "Millisecond":
		t = now.UnixMilli()
	case "Second":
		t = now.Unix()
	}
	return t
}

func TimeStr(str string) string {
	t := Time(str)
	return strconv.FormatInt(t, 10)
}

//-- format : "2006-01-02 03:04:05 PM" --
/*
@timestamp 传0 ，即现在时间

月份 1,01,Jan,January
日　 2,02,_2
时　 3,03,15,PM,pm,AM,am
分　 4,04
秒　 5,05
年　 06,2006
周几 Mon,Monday
时区时差表示 -07,-0700,Z0700,Z07:00,-07:00,MST
时区字母缩写 MST
*/
func Date(timestamp int64, format string) string {
	if timestamp == 0 {
		timestamp = Time("Second")
	}
	tm := time.Unix(timestamp, 0)
	return tm.Format(format)
}

func DateTimeGet() int64 { //获取当天0点时间戳
	t := time.Now()
	timeStr := t.Format("2006-01-02")
	t, _ = time.ParseInLocation("2006-01-02", timeStr, time.Now().Location())
	return t.Unix()
}

// -- "01/02/2006", "02/08/2015" --
func StrToTimeByDate(format string, input string) int64 {
	tm2, _ := time.ParseInLocation(format, input, time.Now().Location())
	return tm2.Unix()
}

// format:2006-01-02 15:04:05, input:2023-12-31 12:59:59
func StrToTime(format string, input string) int64 {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(format, input, loc)
	return theTime.Unix()
}

// 获取两个时间相差的天数，0表同一天，正数表t1>t2，负数表t1<t2
func GetDiffDays(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return int(t1.Sub(t2).Hours() / 24)
}

// 获取t1和t2的相差天数，单位：秒，0表同一天，正数表t1>t2，负数表t1<t2
func GetDiffDaysBySecond(t1, t2 int64) int {
	time1 := time.Unix(t1, 0)
	time2 := time.Unix(t2, 0)
	return GetDiffDays(time1, time2)
}
