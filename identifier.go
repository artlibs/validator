// @Since 2024-03-21.
// @Author Fury, All rights Reserved.

package validator

import "regexp"

var identifierRegexp = regexp.MustCompile("^[_a-zA-Z][_a-zA-Z0-9]*$")

// IsValidIdentifier 字母数字下划线标识符
// allowEmpty: return true if id is nil or ""
// Validate is the id a valid alpha number (with _) identifier
func IsValidIdentifier(id string, allowEmpty bool) bool {
	if value, ok := testNotBlank(id); ok {
		return identifierRegexp.MatchString(value)
	}

	return allowEmpty && isEmpty(id)
}
