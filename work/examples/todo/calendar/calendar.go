package calendar

import (
	"fmt"
	"strconv"
)

// this is basically the same as Haskell's "derive Show"
//go:generate stringer -type dayOfWeek

type dayOfWeek int

const (
	Sunday dayOfWeek = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

//go:generate stringer -type month

type month int

const (
	January month = iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func next(x dayOfWeek) dayOfWeek {
	switch x {
	case Saturday: // reset to 0
		return Sunday
	default:
		return dayOfWeek(int(x) + 1)
	}
}

func pad(n int) string {
	switch {
	case n < 10:
		return " " + strconv.Itoa(n)
	default:
		return strconv.Itoa(n)
	}
}

const weekHeader = "Su Mo Tu We Th Fr Sa"

func calendarMonth(m month, startDay dayOfWeek, maxDay int) string {
	var days func(dayOfWeek, int) string
	days = func(d dayOfWeek, n int) string {
		switch {
		case d == Sunday && n > maxDay:
			return "\n"
		case n > maxDay:
			return "\n\n"
		case d == Saturday:
			return pad(n) + "\n" + days(Sunday, n+1)
		default:
			return pad(n) + " " + days(next(d), n+1)
		}
	}
	// spaces is used to add the initial padding in every month,
	// depending on the starting day
	var spaces func(dayOfWeek) string
	spaces = func(currDay dayOfWeek) string {
		switch currDay {
		case startDay:
			return days(startDay, 1)
		default:
			return "   " + spaces(next(currDay))
		}
	}

	return fmt.Sprintf("%v 2020\n%v\n%v", m, weekHeader, spaces(Sunday))
}

func Calendar() {
	year := calendarMonth(January, Wednesday, 31) +
		calendarMonth(February, Saturday, 29) +
		calendarMonth(March, Sunday, 31) +
		calendarMonth(April, Wednesday, 30) +
		calendarMonth(May, Friday, 31) +
		calendarMonth(June, Monday, 30) +
		calendarMonth(July, Wednesday, 31) +
		calendarMonth(August, Saturday, 31) +
		calendarMonth(September, Tuesday, 30) +
		calendarMonth(October, Thursday, 31) +
		calendarMonth(November, Sunday, 30) +
		calendarMonth(December, Tuesday, 31)

	fmt.Println(year)

}
