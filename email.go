// @Since 2024-03-21.
// @Author Fury, All rights Reserved.

package validator

import "regexp"

var emailRegexp = regexp.MustCompile("\\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Z|a-z]{2,}\\b")

// IsValidEmail 电子邮件
// Validate is the value a valid email
func IsValidEmail(value string, allowBlank bool) bool {
	if isBlank(value) {
		return allowBlank
	}

	return emailRegexp.MatchString(value)
}
