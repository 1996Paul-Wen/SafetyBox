package structvalidator

import (
	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
)

var Validator *validator.Validate

var Myjson jsoniter.API

func init() {
	Myjson = jsoniter.Config{
		EscapeHTML:             true,
		CaseSensitive:          true, // 配置大小写敏感
		ValidateJsonRawMessage: true,
	}.Froze()
	Validator = validator.New()
}
