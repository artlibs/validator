## Validator for Go
[![Run Tests](https://github.com/x-validator/validator/actions/workflows/test.yml/badge.svg)](https://github.com/x-validator/validator/actions/workflows/test.yml)  [![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](http://opensource.org/licenses/MIT)

一个适用于国内生态的go验证器代码库，支持如下一些验证场景：

-   `phone/手机号`：大陆手机号码验证
-   `id card/身份证`：大陆身份证号码验证
-   `email/电子邮件`：电子邮件验证
-   `money/金额`：金额字符串验证

## 开始使用

### 前置条件

-   **[Go](https://go.dev/)**：`1.18`及以后版本，[ go releases](https://go.dev/doc/devel/release) 

### 获取Validator

使用[Go module](https://github.com/golang/go/wiki/Modules) 并往你的代码中导入validator包

```go
import "github.com/x-validator/validator"
```

然后执行 `go [build|run|test]` 就会自动下载该依赖；或者手动下载依赖：

```shell
$ go get -u github.com/x-validator/validator
```

### API列表

```go
// 默认支持Basic、Virtual、NetOnly三种类型
bool valid = validator.IsValidPhoneNumber("11011001111", false)

// 指定支持类型
bool valid = validator.IsValidPhoneNumber(
    "11011001111", false, BasicPhoneNumber, VirtualPhoneNumber
)

bool valid = validator.IsValidMoney("23.54", false)

bool valid = validator.IsValidIdentifier("5asdf", false)

bool valid = validator.IsValidIdCard("11012345", false)

bool valid = validator.IsValidHttpURL("http://example.org", false)

bool valid = validator.IsValidEmail("abcd@example.org", false)
```

#### PhoneType 手机号类型

可以选择验证哪些手机号，或者全部手机号，参考了[ChinaMobilePhoneNumberRegex](https://github.com/VincentSit/ChinaMobilePhoneNumberRegex)项目：

-   **BasicPhoneNumber**：仅基础运营商手机卡，11位手机卡-基础运营商,支持语音通话/短信/数据流量
-   **VirtualPhoneNumber**：仅虚拟运营商号码，11位手机卡-虚拟运营商,支持语音通话/短信/数据流量
-   **NetOnlyPhoneNumber**：仅上网数据卡，11位上网卡,支持语音通话(部分)/短信/数据流量
-   **IotOnlyPhoneNumber**：仅物联网数据卡，13位物联网数据卡,支持数据流量
-   **AllTypePhoneNumber**：验证所有有效号码（手机卡 + 数据卡 + 上网卡）

## 贡献

欢迎贡献你的代码，一起完善validator库！
