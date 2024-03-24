// @Since 2024-03-24.
// @Author Fury, All rights Reserved.

package validator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidMoney(t *testing.T) {
	assert.True(t, IsValidMoney("", true))
	assert.False(t, IsValidMoney("", false))

	validMoney := []string{
		"0",
		"0.00",
		"123",
		"123,456",
		"0.123",
		"12.34",
		"123,23.12",
		"e6",
		"12e",
		"1.23e6",
		"1.2345E-4",
	}

	for _, money := range validMoney {
		fmt.Println(money)
		assert.True(t, IsValidMoney(money, true))
		assert.True(t, IsValidMoney(money, false))
	}

	invalidMoney := []string{
		"alpha",
		"12a.34",
		"12a34",
		"12,34a",
		"123-2344.134E",
		"123, 456",
	}

	for _, money := range invalidMoney {
		assert.False(t, IsValidMoney(money, true))
		assert.False(t, IsValidMoney(money, false))
	}
}
