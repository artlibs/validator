package validator

type Fn func(value string, allowEmpty bool) bool

// TagMaps tags map, when you want to add a new validation to the validator
// you can range this map to get the tag and valid function
var TagMaps = map[string]Fn{
	"email":      IsValidEmail,
	"http_url":   IsValidHttpURL,
	"id_card":    IsValidIdCard,
	"identifier": IsValidIdentifier,
	"money":      IsValidMoney,
	"phone_number": func(value string, allowEmpty bool) bool {
		return IsValidPhoneNumber(value, allowEmpty, AllTypePhoneNumber)
	},
}
