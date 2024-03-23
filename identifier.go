// @Since 2024-03-21.
// @Author Fury, All rights Reserved.

package validator

import "regexp"

var identifierRegexp = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9]*$")

// IsValidIdentifier 字母数字标识符
// Validate is the value a valid alpha number
func IsValidIdentifier(id string, allowBlank bool) bool {
	if value, ok := testBlank(id); ok {
		return identifierRegexp.MatchString(value)
	}

	return allowBlank
}
