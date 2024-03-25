// @Since 2024-03-21.
// @Author Fury, All rights Reserved.

package validator

import "regexp"

// https://github.com/VincentSit/ChinaMobilePhoneNumberRegex/blob/master/README-CN.md
var basicRegexp = regexp.MustCompile("^(?:\\+?86)?1(?:3\\d{3}|5[^4\\D]\\d{2}|8\\d{3}|7(?:[235-8]\\d{2}|4(?:0\\d|1[0-2]|9\\d))|9[0-35-9]\\d{2}|66\\d{2})\\d{6}$")
var virtualRegexp = regexp.MustCompile("^(?:\\+?86)?1(?:7[01]|6[257])\\d{8}$")
var netOnlyRegexp = regexp.MustCompile("^(?:\\+?86)?14[579]\\d{8}$")
var iotOnlyRegexp = regexp.MustCompile("^(?:\\+?86)?14(?:[14]0|41|[68]\\d)\\d{9}$")
var allTypeRegexp = regexp.MustCompile("^(?:\\+?86)?1(?:3\\d{3}|5[^4\\D]\\d{2}|8\\d{3}|7(?:[0-35-9]\\d{2}|4(?:0\\d|1[0-2]|9\\d))|9[0-35-9]\\d{2}|6[2567]\\d{2}|4(?:(?:10|4[01])\\d{3}|[68]\\d{4}|[579]\\d{2}))\\d{6}$")

type PhoneNumberType int

const (
	// BasicPhoneNumber 11位手机卡-基础运营商,支持语音通话/短信/数据流量
	BasicPhoneNumber PhoneNumberType = iota
	// VirtualPhoneNumber 11位手机卡-虚拟运营商,支持语音通话/短信/数据流量
	VirtualPhoneNumber
	// NetOnlyPhoneNumber 11位上网卡,支持语音通话(部分)/短信/数据流量
	NetOnlyPhoneNumber
	// IotOnlyPhoneNumber 13位物联网数据卡,支持数据流量
	IotOnlyPhoneNumber
	// AllTypePhoneNumber 包含以上4种
	AllTypePhoneNumber
)

func getPhoneRegexp(phoneType PhoneNumberType) *regexp.Regexp {
	switch phoneType {
	case BasicPhoneNumber:
		return basicRegexp
	case VirtualPhoneNumber:
		return virtualRegexp
	case NetOnlyPhoneNumber:
		return netOnlyRegexp
	case IotOnlyPhoneNumber:
		return iotOnlyRegexp
	default: // all type, 11或13位,支持所有号码
		return allTypeRegexp
	}
}

// IsValidPhoneNumber 手机号码
// allowEmpty: return true if phone is nil or ""
// Validate is the phone a valid phone number
func IsValidPhoneNumber(phone string, allowEmpty bool, types ...PhoneNumberType) bool {
	if value, ok := testNotBlank(phone); ok {
		if len(types) == 0 {
			types = []PhoneNumberType{BasicPhoneNumber, VirtualPhoneNumber, NetOnlyPhoneNumber}
		}

		for _, phoneType := range types {
			if getPhoneRegexp(phoneType).MatchString(value) {
				return true
			}
			if phoneType == AllTypePhoneNumber {
				break
			}
		}

		return false
	}

	return allowEmpty && isEmpty(phone)
}
