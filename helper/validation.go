package helper

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/muhammadsarimin/simple-api-xmu/types"
)

var validation *validator.Validate

func init() {
	validation = validator.New()
	validation.RegisterTagNameFunc(jsonName)
}

func Validate(T interface{}) error {

	e := validation.Struct(T)
	if e != nil {

		err := e.(validator.ValidationErrors)[0]

		if err.Tag() == "required" {
			return &types.CustomError{
				Code:    "001",
				Message: err.Field() + " is required",
			}
		}

		return &types.CustomError{
			Code:    "002",
			Message: err.Field() + "must be " + err.Tag(),
		}
	}

	return nil
}

func jsonName(fld reflect.StructField) string {

	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

	if name == "-" {
		return ""
	}

	return name
}
