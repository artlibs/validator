// @Since 2024-03-21.
// @Author Fury, All rights Reserved.

package validator

import "regexp"

var emailRegexp = regexp.MustCompile("\\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Z|a-z]{2,}\\b")

// IsValidEmail 电子邮件
// Validate is the email a valid email address
func IsValidEmail(email string, allowBlank bool) bool {
	if value, ok := testBlank(email); ok {
		return emailRegexp.MatchString(value)
	}

	return allowBlank
}
