// Code generated by "stringer -type parity"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[even-0]
	_ = x[odd-1]
}

const _parity_name = "evenodd"

var _parity_index = [...]uint8{0, 4, 7}

func (i parity) String() string {
	if i >= parity(len(_parity_index)-1) {
		return "parity(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _parity_name[_parity_index[i]:_parity_index[i+1]]
}
