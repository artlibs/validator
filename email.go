// @Since 2024-03-21.
// @Author Fury, All rights Reserved.

package validator

import "regexp"

var emailRegexp = regexp.MustCompile("\\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Z|a-z]{2,}\\b")

// IsValidEmail 电子邮件
// allowEmpty: return true if email is nil or ""
// Validate is the email a valid email address
func IsValidEmail(email string, allowEmpty bool) bool {
	if value, ok := testNotBlank(email); ok {
		return emailRegexp.MatchString(value)
	}

	return allowEmpty && isEmpty(email)
}
