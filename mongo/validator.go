package mongo

import (
	"regexp"
	"api/debugger"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate


var (
	firstNameRegex    = regexp.MustCompile(`^[a-zA-Z]+$`)
	lastNameRegex    = regexp.MustCompile(`^[a-zA-Z]+$`)
	phoneRegex = regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	cityRegex    = regexp.MustCompile(`^[a-zA-Z a-zA-Z]+$`)
)

func firstNameValidator(fl validator.FieldLevel) bool {
	return firstNameRegex.MatchString(fl.Field().String())
}

func lastNameValidator(fl validator.FieldLevel) bool {
	return lastNameRegex.MatchString(fl.Field().String())
}

func phoneValidator(fl validator.FieldLevel) bool {
    return phoneRegex.MatchString(fl.Field().String())
}

func cityValidator(fl validator.FieldLevel) bool {
	return cityRegex.MatchString(fl.Field().String())
}

func (cv *Cv) Validate() error {
    validate := validator.New()

	if err := validate.RegisterValidation("first_name", firstNameValidator); err != nil {
		debugger.CheckError("RegisterValidation", err)
	}

	if err := validate.RegisterValidation("last_name", lastNameValidator); err != nil {
		debugger.CheckError("RegisterValidation", err)
	}

	if err := validate.RegisterValidation("phone_number", phoneValidator); err != nil {
		debugger.CheckError("RegisterValidation", err)
	}

	if err := validate.RegisterValidation("city", cityValidator); err != nil {
		debugger.CheckError("RegisterValidation", err)
	}

	err := validate.Struct(cv)
	debugger.CheckError("Struct", err)

	return err
}


