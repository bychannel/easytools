// Copyright 2020 bychannel(bychannel@qq.com). All rights reserved.
// Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.

package easytime

import "time"

func IsSameDay(time1, time2 time.Time, offset int32) bool {
	time1 = time1.Add(time.Duration(-offset) * time.Hour)
	time2 = time2.Add(time.Duration(-offset) * time.Hour)
	return time1.Year() == time2.Year() && time1.YearDay() == time2.YearDay()
}

func IsSameWeek(time1, time2 time.Time) bool {
	week1, w1 := time1.ISOWeek()
	week2, w2 := time2.ISOWeek()
	return w1 == w2 && week1 == week2
}

// 格式化成指定layout的字符串
func FormatToString(t time.Time, layout string) string {
	return t.Format(layout)
}

// Unix时间戳转换秒数
func GetTimeSec(t time.Time) int64 {
	return t.Unix()
}

// Unix时间戳转换毫秒数
func GetTimeMills(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// Unix时间戳转换成时间
func GetTimeBySec(t int64) time.Time {
	return time.Unix(t, 0)
}

// 字符串转换时间
func GetTimeByString(timeStr string, layout string) (time.Time, error) {
	if timeStr == "" {
		return time.Time{}, nil
	}
	return time.ParseInLocation(layout, timeStr, time.Local)
}

// 判断当前时间是否为整点
func CheckHours() bool {
	_, min, sec := time.Now().Clock()
	if min == sec && min == 0 {
		return true
	}
	return false
}

// 判断两个日期时间是否是同年同月份.
func IsSameMonth(t1, t2 time.Time) bool {
	// fixme
	return false
}

// 给定的时间上增加年数.
func AddYears(t time.Time, add int) time.Time {
	// fixme
	//return t.AddDate(add, 0, 0)

	//time_part, err := time.ParseDuration("2h")
	//if err != nil {
	//	return t
	//}
	//return t.Add(time_part)
	return time.Now()
}

// 给定的时间上增加月数.
func AddMonths() {

}

// 给定的时间上增加日数.
func AddDays() {

}

// 给定的时间上增加小时数.
func AddHours() {

}

// 给定的时间上增加分钟数.
func AddMinutes() {

}

// 给定的时间上增加秒数.
func AddSeconds() {

}

// 计算两个时间相差的天数，
// t1小于t2可能会出现负值.
func DiffDays(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return int(t1.Sub(t2).Hours() / 24)
}

// 计算两个时间相差的秒数.
// t1小于t2可能会出现负值.
func DiffSeconds(t1, t2 time.Time) int64 {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), t1.Hour(), t1.Minute(), t1.Second(), 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), t2.Hour(), t2.Minute(), t2.Second(), 0, time.Local)

	return int64(t1.Sub(t2).Seconds())
}
