// Code generated by "stringer -type=ModuleType -linecomment"; DO NOT EDIT.

package vobj

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ModuleTypeUnknown-0]
	_ = x[ModuleTypeMenu-1]
}

const _ModuleType_name = "未知菜单模块"

var _ModuleType_index = [...]uint8{0, 6, 18}

func (i ModuleType) String() string {
	if i < 0 || i >= ModuleType(len(_ModuleType_index)-1) {
		return "ModuleType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ModuleType_name[_ModuleType_index[i]:_ModuleType_index[i+1]]
}

// IsUnknown 是否是：未知
func (i ModuleType) IsUnknown() bool {
	return i == ModuleTypeUnknown
}

// IsMenu 是否是：菜单模块
func (i ModuleType) IsMenu() bool {
	return i == ModuleTypeMenu
}

// GetValue 获取原始类型值
func (i ModuleType) GetValue() int {
	return int(i)
}
