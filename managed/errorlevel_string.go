// Code generated by "stringer -type=ErrorLevel"; DO NOT EDIT.

package manager

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ErrorLevelUnspecified-0]
	_ = x[ErrorLevelOk-1]
	_ = x[ErrorLevelLo-2]
	_ = x[ErrorLevelHi-3]
}

const _ErrorLevel_name = "ErrorLevelUnspecifiedErrorLevelOkErrorLevelLoErrorLevelHi"

var _ErrorLevel_index = [...]uint8{0, 21, 33, 45, 57}

func (i ErrorLevel) String() string {
	if i >= ErrorLevel(len(_ErrorLevel_index)-1) {
		return "ErrorLevel(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ErrorLevel_name[_ErrorLevel_index[i]:_ErrorLevel_index[i+1]]
}
