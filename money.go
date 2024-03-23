package validator

import "regexp"

var moneyRegexp = regexp.MustCompile("^([0-9]+|[0-9]{1,3}(,[0-9]{3})*)(.[0-9]{1,2})?$")

func IsValidMoney(money string, allowBlank bool) bool {
	if value, ok := testBlank(money); ok {
		return moneyRegexp.MatchString(value)
	}

	return allowBlank
}
