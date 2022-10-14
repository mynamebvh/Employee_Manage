package department

import (
	errorModels "employee_manage/constant"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func validateStructDepartment(department NewDepartment) (errValidate []errorModels.ErrorValidate) {
	validate = validator.New()
	err := validate.Struct(department)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errValidate = append(errValidate, errorModels.ErrorValidate{
				Field:   err.StructField(),
				Message: fmt.Sprintf("%s a is %s", strings.ToLower(err.StructField()), err.Tag()),
			})
		}
	}

	return
}

func updateValidation(request map[string]interface{}) (errValidate []errorModels.ErrorValidate) {

	validationMap := map[string]string{
		"name":            "required,gt=3,lt=100",
		"department_code": "required,gt=1,lt=100",
		"address":         "required,gt=3,lt=100",
		"status":          "required,boolean",
	}

	validate := validator.New()
	err := validate.RegisterValidation("update_department", func(fl validator.FieldLevel) bool {
		m, ok := fl.Field().Interface().(map[string]interface{})
		if !ok {
			return false
		}

		for k, v := range validationMap {
			errVar := validate.Var(m[k], v)
			if errVar != nil {
				validatorErr := errVar.(validator.ValidationErrors)
				errValidate = append(errValidate, errorModels.ErrorValidate{
					Field:   k,
					Message: fmt.Sprintf("%s a is %s", k, validatorErr[0].Tag()),
				})
			}
		}

		return true
	})

	if err != nil {
		err = errorModels.NewAppError(err, errorModels.UnknownError)
		return
	}

	err = validate.Var(request, "update_department")

	if err != nil {
		err = errorModels.NewAppError(err, errorModels.UnknownError)
		return
	}

	if errValidate != nil {
		return
	}

	return nil
}
