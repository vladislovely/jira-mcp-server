package validator

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/mark3labs/mcp-go/mcp"
)

func ValidateInput(validate *validator.Validate, input interface{}) *mcp.CallToolResult {
	if err := validate.Struct(input); err != nil {
		return mcp.NewToolResultError(humanizeValidationError(err).Error())
	}
	return nil
}

func humanizeValidationError(err error) error {
	if err == nil {
		return nil
	}

	var validationErrs validator.ValidationErrors
	if errors.As(err, &validationErrs) {
		for _, e := range validationErrs {
			field := e.Field()
			tag := e.Tag()

			if field == "ProjectType" && tag == "required_if" {
				return errors.New(
					"поле 'project_type' обязательно, если 'project_type_key' = 'software'",
				)
			}
		}
	}

	return err
}
