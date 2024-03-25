## validator
[![Run Tests](https://github.com/x-validator/validator/actions/workflows/test.yml/badge.svg)](https://github.com/x-validator/validator/actions/workflows/test.yml)

A go library for validate phone, bank card, id card, email, money e.g.

## Usage

### Install

```shell
$ go get -u "github.com/x-validator/validator"
```

### Using

```go
import "github.com/x-validator/validator"

bool valid = validator.IsValidPhone("11011001111", false, []PhoneType{ AllType })

bool valid = validator.IsValidMoney("23.54", false)

bool valid = validator.IsValidIdentifier("5asdf", false)

bool valid = validator.IsValidIdCard("11012345", false)

bool valid = validator.IsValidHttpURL("http://example.org", false)

bool valid = validator.IsValidEmail("abcd@example.org", false)

bool valid = validator.IsValidBankCard("1101010110", false)
```

