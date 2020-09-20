// Code generated by "enumer -type=Type -transform=lower -output=type_string.go"; DO NOT EDIT.

package plugin

import (
	"fmt"
	"strings"
)

const _TypeName = "localremoteembedded"

var _TypeIndex = [...]uint8{0, 5, 11, 19}

const _TypeLowerName = "localremoteembedded"

func (i Type) String() string {
	if i < 0 || i >= Type(len(_TypeIndex)-1) {
		return fmt.Sprintf("Type(%d)", i)
	}
	return _TypeName[_TypeIndex[i]:_TypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _TypeNoOp() {
	var x [1]struct{}
	_ = x[Local-(0)]
	_ = x[Remote-(1)]
	_ = x[Embedded-(2)]
}

var _TypeValues = []Type{Local, Remote, Embedded}

var _TypeNameToValueMap = map[string]Type{
	_TypeName[0:5]:        Local,
	_TypeLowerName[0:5]:   Local,
	_TypeName[5:11]:       Remote,
	_TypeLowerName[5:11]:  Remote,
	_TypeName[11:19]:      Embedded,
	_TypeLowerName[11:19]: Embedded,
}

var _TypeNames = []string{
	_TypeName[0:5],
	_TypeName[5:11],
	_TypeName[11:19],
}

// TypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func TypeString(s string) (Type, error) {
	if val, ok := _TypeNameToValueMap[s]; ok {
		return val, nil
	}
	s = strings.ToLower(s)
	if val, ok := _TypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Type values", s)
}

// TypeValues returns all values of the enum
func TypeValues() []Type {
	return _TypeValues
}

// TypeStrings returns a slice of all String values of the enum
func TypeStrings() []string {
	strs := make([]string, len(_TypeNames))
	copy(strs, _TypeNames)
	return strs
}

// IsAType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Type) IsAType() bool {
	for _, v := range _TypeValues {
		if i == v {
			return true
		}
	}
	return false
}
