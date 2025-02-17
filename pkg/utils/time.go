package utils

import (
	"device-manager/utils/constants"
	"math"
	"time"
)

func ConvertStringToDate(dateStr string) (time.Time, error) {
	result, err := time.Parse(string(constants.DateFormat), dateStr)
	return result, err
}

func ConvertStringToDateTime(datetimeStr string) (time.Time, error) {
	result, err := time.Parse(string(constants.DatetimeFormat), datetimeStr)
	return result, err
}

func ConvertIOSStringToDateTime(datetimeStr string) (time.Time, error) {
	result, err := time.Parse(string(constants.DateTimeIOSFormat), datetimeStr)
	return result, err
}

func ConvertStringToTime(timeStr string) (time.Time, error) {
	result, err := time.Parse(string(constants.TimeFormat), timeStr)
	return result, err
}

func CalculateAge(birthday time.Time) (int64, error) {
	today := time.Now()
	age := today.Year() - birthday.Year()
	if birthday.AddDate(age, 0, 0).After(today) {
		age--
	}
	return int64(age), nil
}
func CalculateAgeFromString(birthday string) (int64, error) {
	bodTime, err := ConvertStringToDate(birthday)
	if err != nil {
		return 0, err
	}
	return CalculateAge(bodTime)
}
func IsSameYear(source1, source2 string, layout string) (bool, error) {
	dt1, err := time.Parse(layout, source1)
	if err != nil {
		return false, err
	}
	dt2, err := time.Parse(layout, source2)
	if err != nil {
		return false, err
	}
	return dt1.Year() == dt2.Year(), nil
}
func IsSameMonth(source1, source2 string, layout string) (bool, error) {
	dt1, err := time.Parse(layout, source1)
	if err != nil {
		return false, err
	}
	dt2, err := time.Parse(layout, source2)
	if err != nil {
		return false, err
	}
	return dt1.Year() != dt2.Year() || dt1.Month() != dt2.Month(), nil
}
func IsSameWeek(source1, source2 string, layout string) (bool, error) {
	dt1, err := time.Parse(layout, source1)
	if err != nil {
		return false, err
	}
	dt2, err := time.Parse(layout, source2)
	if err != nil {
		return false, err
	}
	dt1Year, dt1Week := dt1.ISOWeek()
	dt2Year, dt2Week := dt2.ISOWeek()
	return dt1Year != dt2Year || dt1Week != dt2Week, nil
}
func IsSameDay(source1, source2 string, layout string) (bool, error) {
	dt1, err := time.Parse(layout, source1)
	if err != nil {
		return false, err
	}
	dt2, err := time.Parse(layout, source2)
	if err != nil {
		return false, err
	}
	return dt1.Year() != dt2.Year() || dt1.Month() != dt2.Month() || dt1.Day() != dt2.Day(), nil
}

func IsSameHour(source1, source2 string, layout string) (bool, error) {
	dt1, err := time.Parse(layout, source1)
	if err != nil {
		return false, err
	}
	dt2, err := time.Parse(layout, source2)
	if err != nil {
		return false, err
	}
	return dt1.Year() != dt2.Year() || dt1.Month() != dt2.Month() || dt1.Day() != dt2.Day() || dt1.Hour() != dt2.Hour(), nil
}

func DiffOfTime(source1, source2 string, layout string) (int64, error) {
	dt1, err := time.Parse(layout, source1)
	if err != nil {
		return 0, err
	}
	dt2, err := time.Parse(layout, source2)
	if err != nil {
		return 0, err
	}
	return int64(math.Abs(float64(dt1.Unix() - dt2.Unix()))), nil
}
