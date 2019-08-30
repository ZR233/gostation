/*
@Time : 2019-07-08 11:12
@Author : zr
*/
package utils

import (
	"time"
)

func TimeFormatter() string {
	return "2006-01-02 15:04:05"
}
func TimeFormatterDate() string {
	return "2006-01-02"
}

func TimeStdQueryBegin(stdTimeStr string) (newTime time.Time, err error) {
	if stdTimeStr == "" {
		return time.Parse(TimeFormatter(), "2000-01-01 00:00:00")
	}

	newTime, err = time.Parse(TimeFormatter(), stdTimeStr)
	if err != nil {
		return newTime, err
	}
	newTimeStr := newTime.Format("2006-01-02") + " 00:00:00"

	return time.Parse(TimeFormatter(), newTimeStr)
}
func TimeStdQueryEnd(stdTimeStr string) (newTime time.Time, err error) {
	if stdTimeStr == "" {
		newTime = time.Now()
	} else {
		newTime, err = time.Parse(TimeFormatter(), stdTimeStr)
		if err != nil {
			return newTime, err
		}
	}
	newTimeStr := newTime.Format("2006-01-02") + " 23:59:59"
	return time.Parse(TimeFormatter(), newTimeStr)
}

func TimeDayBegin(day time.Time) time.Time {
	dayBeginStr := day.Format("2006-01-02") + " 00:00:00"
	dayBegin, _ := time.Parse(TimeFormatter(), dayBeginStr)
	return dayBegin
}
func TimeTodayBegin() time.Time {
	timeNow := time.Now()
	dayBeginStr := timeNow.Format("2006-01-02") + " 00:00:00"
	dayBegin, _ := time.Parse(TimeFormatter(), dayBeginStr)
	return dayBegin
}

func TimeDayEnd(day time.Time) time.Time {
	dayStr := day.Format("2006-01-02") + " 23:59:59"
	day, _ = time.Parse(TimeFormatter(), dayStr)
	return day
}
