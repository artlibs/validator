// @Since 2024-03-22.
// @Author Leo, All rights Reserved.

package validator

import "regexp"

var moneyRegexp = regexp.MustCompile("^([0-9]+|[0-9]{1,3}(,[0-9]{3})*)(.[0-9]{1,2})?$")

// IsValidMoney 金钱
// Validate is the money a valid money
func IsValidMoney(money string, allowBlank bool) bool {
	if value, ok := testBlank(money); ok {
		return moneyRegexp.MatchString(value)
	}

	return allowBlank
}
