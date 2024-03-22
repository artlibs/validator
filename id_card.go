// @Since 2024-03-21.
// @Author Fury, All rights Reserved.

package validator

// IsValidIdCard 身份证
// Validate is the value a valid id card number
func IsValidIdCard(value string, allowBlank bool) bool {
	if isBlank(value) {
		return allowBlank
	}

	return true
}
