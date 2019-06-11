package users

import (
	"fmt"

	"github.com/gomematic/gomematic-api/pkg/model"
)

type (
	// ValidationErrors are returned with a slice of all invalid fields.
	ValidationErrors struct {
		Errors []ValidationError
	}

	// ValidationError knows for a given field the error.
	ValidationError struct {
		Field string
		Error error
	}
)

func (e ValidationErrors) Error() string {
	return fmt.Sprintf("there are %d validation errors", len(e.Errors))
}

// ValidateCreate takes a user and validates its fields.
func ValidateCreate(record *model.User) error {
	errs := ValidationErrors{}

	// if err := validateFoo(record.Foo); err != nil {
	// 	errs.Errors = append(errs.Errors, ValidationError{
	// 		Field: "foo",
	// 		Error: err,
	// 	})
	// }

	if len(errs.Errors) > 0 {
		return errs
	}

	return nil
}

// ValidateUpdate takes a user and validates its fields.
func ValidateUpdate(record *model.User) error {
	errs := ValidationErrors{}

	// if err := validateFoo(record.Foo); err != nil {
	// 	errs.Errors = append(errs.Errors, ValidationError{
	// 		Field: "foo",
	// 		Error: err,
	// 	})
	// }

	if len(errs.Errors) > 0 {
		return errs
	}

	return nil
}

func validatePerm(value string) error {
	perms := []string{
		"user",
		"admin",
		"owner",
	}

	for _, perm := range perms {
		if perm == value {
			return nil
		}
	}

	return ValidationErrors{
		Errors: []ValidationError{
			ValidationError{
				Field: "perm",
				Error: fmt.Errorf("invalid permission value"),
			},
		},
	}
}
