package request

import (
	"reflect"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-elastic-search/src/api/config"
	"gopkg.in/go-playground/validator.v8"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("isValidDateFormat", isValidDateFormat)
	}
}

func isValidDateFormat(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	fieldStr := field.String()
	_, parseErr := time.Parse(config.TIME_LAYOUT, fieldStr)
	return parseErr == nil
}
