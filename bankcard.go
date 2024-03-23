// @Since 2024-03-21.
// @Author Fury, All rights Reserved.

package validator

import "regexp"

var numberOnlyRegexp = regexp.MustCompile("\\d+")

// IsValidBankCard 银行卡号
// Validate is the bankCardNo a valid bank card number
func IsValidBankCard(bankCardNo string, allowBlank bool) bool {
	if value, ok := testBlank(bankCardNo); ok {
		size := len(value)
		if size < 15 || size > 19 {
			return false
		}

		last := rune(value[size-1])
		check := bkCvc(value[:size-1])

		return check != 'N' && check == last
	}

	return allowBlank
}

func bkCvc(value string) rune {
	if !numberOnlyRegexp.MatchString(value) {
		return 'N'
	}

	sum, j := 0, 0
	for i := len(value) - 1; i >= 0; i-- {
		j++
		k := int(value[i]) - int('0')

		if j%2 == 0 {
			k *= 2
			k = k/10 + k%10
		}
		sum += k
	}

	if sum%10 == 0 {
		return '0'
	}

	return rune(10-sum%10) + '0'
}
