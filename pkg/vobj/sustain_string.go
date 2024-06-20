// Code generated by "stringer -type=Sustain -linecomment"; DO NOT EDIT.

package vobj

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SustainUnknown-0]
	_ = x[SustainFor-1]
	_ = x[SustainMax-2]
	_ = x[SustainMin-3]
}

const _Sustain_name = "未知m时间内出现n次m时间内最多出现n次m时间内最少出现n次"

var _Sustain_index = [...]uint8{0, 6, 26, 52, 78}

func (i Sustain) String() string {
	if i < 0 || i >= Sustain(len(_Sustain_index)-1) {
		return "Sustain(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Sustain_name[_Sustain_index[i]:_Sustain_index[i+1]]
}

// IsUnknown 是否是：未知
func (i Sustain) IsUnknown() bool {
	return i == SustainUnknown
}

// IsFor 是否是：m时间内出现n次
func (i Sustain) IsFor() bool {
	return i == SustainFor
}

// IsMax 是否是：m时间内最多出现n次
func (i Sustain) IsMax() bool {
	return i == SustainMax
}

// IsMin 是否是：m时间内最少出现n次
func (i Sustain) IsMin() bool {
	return i == SustainMin
}

// GetValue 获取原始类型值
func (i Sustain) GetValue() int {
	return int(i)
}
