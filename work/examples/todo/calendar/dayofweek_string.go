// Code generated by "stringer -type dayOfWeek"; DO NOT EDIT.

package calendar

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Sunday-0]
	_ = x[Monday-1]
	_ = x[Tuesday-2]
	_ = x[Wednesday-3]
	_ = x[Thursday-4]
	_ = x[Friday-5]
	_ = x[Saturday-6]
}

const _dayOfWeek_name = "SundayMondayTuesdayWednesdayThursdayFridaySaturday"

var _dayOfWeek_index = [...]uint8{0, 6, 12, 19, 28, 36, 42, 50}

func (i dayOfWeek) String() string {
	if i < 0 || i >= dayOfWeek(len(_dayOfWeek_index)-1) {
		return "dayOfWeek(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _dayOfWeek_name[_dayOfWeek_index[i]:_dayOfWeek_index[i+1]]
}