// @Since 2024-03-21.
// @Author Fury, All rights Reserved.

package validator

const identifierRegexp string = "^[a-zA-Z][a-zA-Z0-9]*$"

// IsValidIdentifier 字母数字标识符
// Validate is the value a valid alpha number
func IsValidIdentifier(value string, allowBlank bool) bool {
	if isBlank(value) {
		return allowBlank
	}

	return regexp.MustCompile(identifierRegexp).MatchString(value)
}