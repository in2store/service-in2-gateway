package client_in2_book

import (
	"bytes"
	"encoding"
	"errors"

	github_com_johnnyeven_libtools_courier_enumeration "github.com/johnnyeven/libtools/courier/enumeration"
)

// swagger:enum
type In2BookBookLanguage uint

const (
	IN2_BOOK_BOOK_LANGUAGE_UNKNOWN  In2BookBookLanguage = iota
	IN2_BOOK_BOOK_LANGUAGE__ZH_CN                       // 简体中文
	IN2_BOOK_BOOK_LANGUAGE__ZH_TW                       // 繁体中文
	IN2_BOOK_BOOK_LANGUAGE__ENGLISH                     // 英文
)

// swagger:enum
type In2BookBookStatus uint

const (
	IN2_BOOK_BOOK_STATUS_UNKNOWN   In2BookBookStatus = iota
	IN2_BOOK_BOOK_STATUS__PENGDING                   // 等待导入
	IN2_BOOK_BOOK_STATUS__PROCESS                    // 导入中
	IN2_BOOK_BOOK_STATUS__READY                      // 就绪
	IN2_BOOK_BOOK_STATUS__NORMAL                     // 正常展示
)

// swagger:enum
type In2BookCodeLanguage uint

const (
	IN2_BOOK_CODE_LANGUAGE_UNKNOWN     In2BookCodeLanguage = iota
	IN2_BOOK_CODE_LANGUAGE__C_CPP                          // C/C++
	IN2_BOOK_CODE_LANGUAGE__JAVA                           // Java
	IN2_BOOK_CODE_LANGUAGE__JAVASCRIPT                     // Javascript
	IN2_BOOK_CODE_LANGUAGE__PYTHON                         // Python
	IN2_BOOK_CODE_LANGUAGE__CSHARP                         // C#
	IN2_BOOK_CODE_LANGUAGE__GOLANG                         // Golang
	IN2_BOOK_CODE_LANGUAGE__PHP                            // PHP
)

var InvalidIn2BookBookLanguage = errors.New("invalid In2BookBookLanguage")

func init() {
	github_com_johnnyeven_libtools_courier_enumeration.RegisterEnums("In2BookBookLanguage", map[string]string{
		"ZH_CN":   "简体中文",
		"ZH_TW":   "繁体中文",
		"ENGLISH": "英文",
	})
}

func ParseIn2BookBookLanguageFromString(s string) (In2BookBookLanguage, error) {
	switch s {
	case "":
		return IN2_BOOK_BOOK_LANGUAGE_UNKNOWN, nil
	case "ZH_CN":
		return IN2_BOOK_BOOK_LANGUAGE__ZH_CN, nil
	case "ZH_TW":
		return IN2_BOOK_BOOK_LANGUAGE__ZH_TW, nil
	case "ENGLISH":
		return IN2_BOOK_BOOK_LANGUAGE__ENGLISH, nil
	}
	return IN2_BOOK_BOOK_LANGUAGE_UNKNOWN, InvalidIn2BookBookLanguage
}

func ParseIn2BookBookLanguageFromLabelString(s string) (In2BookBookLanguage, error) {
	switch s {
	case "":
		return IN2_BOOK_BOOK_LANGUAGE_UNKNOWN, nil
	case "简体中文":
		return IN2_BOOK_BOOK_LANGUAGE__ZH_CN, nil
	case "繁体中文":
		return IN2_BOOK_BOOK_LANGUAGE__ZH_TW, nil
	case "英文":
		return IN2_BOOK_BOOK_LANGUAGE__ENGLISH, nil
	}
	return IN2_BOOK_BOOK_LANGUAGE_UNKNOWN, InvalidIn2BookBookLanguage
}

func (In2BookBookLanguage) EnumType() string {
	return "In2BookBookLanguage"
}

func (In2BookBookLanguage) Enums() map[int][]string {
	return map[int][]string{
		int(IN2_BOOK_BOOK_LANGUAGE__ZH_CN):   {"ZH_CN", "简体中文"},
		int(IN2_BOOK_BOOK_LANGUAGE__ZH_TW):   {"ZH_TW", "繁体中文"},
		int(IN2_BOOK_BOOK_LANGUAGE__ENGLISH): {"ENGLISH", "英文"},
	}
}
func (v In2BookBookLanguage) String() string {
	switch v {
	case IN2_BOOK_BOOK_LANGUAGE_UNKNOWN:
		return ""
	case IN2_BOOK_BOOK_LANGUAGE__ZH_CN:
		return "ZH_CN"
	case IN2_BOOK_BOOK_LANGUAGE__ZH_TW:
		return "ZH_TW"
	case IN2_BOOK_BOOK_LANGUAGE__ENGLISH:
		return "ENGLISH"
	}
	return "UNKNOWN"
}

func (v In2BookBookLanguage) Label() string {
	switch v {
	case IN2_BOOK_BOOK_LANGUAGE_UNKNOWN:
		return ""
	case IN2_BOOK_BOOK_LANGUAGE__ZH_CN:
		return "简体中文"
	case IN2_BOOK_BOOK_LANGUAGE__ZH_TW:
		return "繁体中文"
	case IN2_BOOK_BOOK_LANGUAGE__ENGLISH:
		return "英文"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*In2BookBookLanguage)(nil)

func (v In2BookBookLanguage) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidIn2BookBookLanguage
	}
	return []byte(str), nil
}

func (v *In2BookBookLanguage) UnmarshalText(data []byte) (err error) {
	*v, err = ParseIn2BookBookLanguageFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidIn2BookBookStatus = errors.New("invalid In2BookBookStatus")

func init() {
	github_com_johnnyeven_libtools_courier_enumeration.RegisterEnums("In2BookBookStatus", map[string]string{
		"PENGDING": "等待导入",
		"PROCESS":  "导入中",
		"READY":    "就绪",
		"NORMAL":   "正常展示",
	})
}

func ParseIn2BookBookStatusFromString(s string) (In2BookBookStatus, error) {
	switch s {
	case "":
		return IN2_BOOK_BOOK_STATUS_UNKNOWN, nil
	case "PENGDING":
		return IN2_BOOK_BOOK_STATUS__PENGDING, nil
	case "PROCESS":
		return IN2_BOOK_BOOK_STATUS__PROCESS, nil
	case "READY":
		return IN2_BOOK_BOOK_STATUS__READY, nil
	case "NORMAL":
		return IN2_BOOK_BOOK_STATUS__NORMAL, nil
	}
	return IN2_BOOK_BOOK_STATUS_UNKNOWN, InvalidIn2BookBookStatus
}

func ParseIn2BookBookStatusFromLabelString(s string) (In2BookBookStatus, error) {
	switch s {
	case "":
		return IN2_BOOK_BOOK_STATUS_UNKNOWN, nil
	case "等待导入":
		return IN2_BOOK_BOOK_STATUS__PENGDING, nil
	case "导入中":
		return IN2_BOOK_BOOK_STATUS__PROCESS, nil
	case "就绪":
		return IN2_BOOK_BOOK_STATUS__READY, nil
	case "正常展示":
		return IN2_BOOK_BOOK_STATUS__NORMAL, nil
	}
	return IN2_BOOK_BOOK_STATUS_UNKNOWN, InvalidIn2BookBookStatus
}

func (In2BookBookStatus) EnumType() string {
	return "In2BookBookStatus"
}

func (In2BookBookStatus) Enums() map[int][]string {
	return map[int][]string{
		int(IN2_BOOK_BOOK_STATUS__PENGDING): {"PENGDING", "等待导入"},
		int(IN2_BOOK_BOOK_STATUS__PROCESS):  {"PROCESS", "导入中"},
		int(IN2_BOOK_BOOK_STATUS__READY):    {"READY", "就绪"},
		int(IN2_BOOK_BOOK_STATUS__NORMAL):   {"NORMAL", "正常展示"},
	}
}
func (v In2BookBookStatus) String() string {
	switch v {
	case IN2_BOOK_BOOK_STATUS_UNKNOWN:
		return ""
	case IN2_BOOK_BOOK_STATUS__PENGDING:
		return "PENGDING"
	case IN2_BOOK_BOOK_STATUS__PROCESS:
		return "PROCESS"
	case IN2_BOOK_BOOK_STATUS__READY:
		return "READY"
	case IN2_BOOK_BOOK_STATUS__NORMAL:
		return "NORMAL"
	}
	return "UNKNOWN"
}

func (v In2BookBookStatus) Label() string {
	switch v {
	case IN2_BOOK_BOOK_STATUS_UNKNOWN:
		return ""
	case IN2_BOOK_BOOK_STATUS__PENGDING:
		return "等待导入"
	case IN2_BOOK_BOOK_STATUS__PROCESS:
		return "导入中"
	case IN2_BOOK_BOOK_STATUS__READY:
		return "就绪"
	case IN2_BOOK_BOOK_STATUS__NORMAL:
		return "正常展示"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*In2BookBookStatus)(nil)

func (v In2BookBookStatus) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidIn2BookBookStatus
	}
	return []byte(str), nil
}

func (v *In2BookBookStatus) UnmarshalText(data []byte) (err error) {
	*v, err = ParseIn2BookBookStatusFromString(string(bytes.ToUpper(data)))
	return
}

var InvalidIn2BookCodeLanguage = errors.New("invalid In2BookCodeLanguage")

func init() {
	github_com_johnnyeven_libtools_courier_enumeration.RegisterEnums("In2BookCodeLanguage", map[string]string{
		"C_CPP":      "C/C++",
		"JAVA":       "Java",
		"JAVASCRIPT": "Javascript",
		"PYTHON":     "Python",
		"CSHARP":     "C#",
		"GOLANG":     "Golang",
		"PHP":        "PHP",
	})
}

func ParseIn2BookCodeLanguageFromString(s string) (In2BookCodeLanguage, error) {
	switch s {
	case "":
		return IN2_BOOK_CODE_LANGUAGE_UNKNOWN, nil
	case "C_CPP":
		return IN2_BOOK_CODE_LANGUAGE__C_CPP, nil
	case "JAVA":
		return IN2_BOOK_CODE_LANGUAGE__JAVA, nil
	case "JAVASCRIPT":
		return IN2_BOOK_CODE_LANGUAGE__JAVASCRIPT, nil
	case "PYTHON":
		return IN2_BOOK_CODE_LANGUAGE__PYTHON, nil
	case "CSHARP":
		return IN2_BOOK_CODE_LANGUAGE__CSHARP, nil
	case "GOLANG":
		return IN2_BOOK_CODE_LANGUAGE__GOLANG, nil
	case "PHP":
		return IN2_BOOK_CODE_LANGUAGE__PHP, nil
	}
	return IN2_BOOK_CODE_LANGUAGE_UNKNOWN, InvalidIn2BookCodeLanguage
}

func ParseIn2BookCodeLanguageFromLabelString(s string) (In2BookCodeLanguage, error) {
	switch s {
	case "":
		return IN2_BOOK_CODE_LANGUAGE_UNKNOWN, nil
	case "C/C++":
		return IN2_BOOK_CODE_LANGUAGE__C_CPP, nil
	case "Java":
		return IN2_BOOK_CODE_LANGUAGE__JAVA, nil
	case "Javascript":
		return IN2_BOOK_CODE_LANGUAGE__JAVASCRIPT, nil
	case "Python":
		return IN2_BOOK_CODE_LANGUAGE__PYTHON, nil
	case "C#":
		return IN2_BOOK_CODE_LANGUAGE__CSHARP, nil
	case "Golang":
		return IN2_BOOK_CODE_LANGUAGE__GOLANG, nil
	case "PHP":
		return IN2_BOOK_CODE_LANGUAGE__PHP, nil
	}
	return IN2_BOOK_CODE_LANGUAGE_UNKNOWN, InvalidIn2BookCodeLanguage
}

func (In2BookCodeLanguage) EnumType() string {
	return "In2BookCodeLanguage"
}

func (In2BookCodeLanguage) Enums() map[int][]string {
	return map[int][]string{
		int(IN2_BOOK_CODE_LANGUAGE__C_CPP):      {"C_CPP", "C/C++"},
		int(IN2_BOOK_CODE_LANGUAGE__JAVA):       {"JAVA", "Java"},
		int(IN2_BOOK_CODE_LANGUAGE__JAVASCRIPT): {"JAVASCRIPT", "Javascript"},
		int(IN2_BOOK_CODE_LANGUAGE__PYTHON):     {"PYTHON", "Python"},
		int(IN2_BOOK_CODE_LANGUAGE__CSHARP):     {"CSHARP", "C#"},
		int(IN2_BOOK_CODE_LANGUAGE__GOLANG):     {"GOLANG", "Golang"},
		int(IN2_BOOK_CODE_LANGUAGE__PHP):        {"PHP", "PHP"},
	}
}
func (v In2BookCodeLanguage) String() string {
	switch v {
	case IN2_BOOK_CODE_LANGUAGE_UNKNOWN:
		return ""
	case IN2_BOOK_CODE_LANGUAGE__C_CPP:
		return "C_CPP"
	case IN2_BOOK_CODE_LANGUAGE__JAVA:
		return "JAVA"
	case IN2_BOOK_CODE_LANGUAGE__JAVASCRIPT:
		return "JAVASCRIPT"
	case IN2_BOOK_CODE_LANGUAGE__PYTHON:
		return "PYTHON"
	case IN2_BOOK_CODE_LANGUAGE__CSHARP:
		return "CSHARP"
	case IN2_BOOK_CODE_LANGUAGE__GOLANG:
		return "GOLANG"
	case IN2_BOOK_CODE_LANGUAGE__PHP:
		return "PHP"
	}
	return "UNKNOWN"
}

func (v In2BookCodeLanguage) Label() string {
	switch v {
	case IN2_BOOK_CODE_LANGUAGE_UNKNOWN:
		return ""
	case IN2_BOOK_CODE_LANGUAGE__C_CPP:
		return "C/C++"
	case IN2_BOOK_CODE_LANGUAGE__JAVA:
		return "Java"
	case IN2_BOOK_CODE_LANGUAGE__JAVASCRIPT:
		return "Javascript"
	case IN2_BOOK_CODE_LANGUAGE__PYTHON:
		return "Python"
	case IN2_BOOK_CODE_LANGUAGE__CSHARP:
		return "C#"
	case IN2_BOOK_CODE_LANGUAGE__GOLANG:
		return "Golang"
	case IN2_BOOK_CODE_LANGUAGE__PHP:
		return "PHP"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*In2BookCodeLanguage)(nil)

func (v In2BookCodeLanguage) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidIn2BookCodeLanguage
	}
	return []byte(str), nil
}

func (v *In2BookCodeLanguage) UnmarshalText(data []byte) (err error) {
	*v, err = ParseIn2BookCodeLanguageFromString(string(bytes.ToUpper(data)))
	return
}
