package users

import (
	"fmt"
	"strings"

	"github.com/asaskevich/govalidator"
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

	if err := validateEmail(record.Email); err != nil {
		errs.Errors = append(errs.Errors, ValidationError{
			Field: "email",
			Error: err,
		})
	}

	if err := validateUsername(record.Username); err != nil {
		errs.Errors = append(errs.Errors, ValidationError{
			Field: "username",
			Error: err,
		})
	}

	if err := validatePassword(record.Password); err != nil {
		errs.Errors = append(errs.Errors, ValidationError{
			Field: "password",
			Error: err,
		})
	}

	if len(errs.Errors) > 0 {
		return errs
	}

	return nil
}

// ValidateUpdate takes a user and validates its fields.
func ValidateUpdate(record *model.User) error {
	errs := ValidationErrors{}

	if err := validateID(record.ID); err != nil {
		errs.Errors = append(errs.Errors, ValidationError{
			Field: "id",
			Error: err,
		})
	}

	if err := validateEmail(record.Email); err != nil {
		errs.Errors = append(errs.Errors, ValidationError{
			Field: "email",
			Error: err,
		})
	}

	if err := validateUsername(record.Username); err != nil {
		errs.Errors = append(errs.Errors, ValidationError{
			Field: "username",
			Error: err,
		})
	}

	if err := validatePassword(record.Password); err != nil {
		errs.Errors = append(errs.Errors, ValidationError{
			Field: "password",
			Error: err,
		})
	}

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

func validateID(val string) error {
	errs := []string{}

	if ok := govalidator.IsUUIDv4(val); !ok {
		errs = append(errs, "is not a valid uuid v4")
	}

	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}

	return nil
}

func validateEmail(val string) error {
	errs := []string{}

	if ok := govalidator.IsEmail(val); !ok {
		errs = append(errs, "is not valid")
	}

	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}

	return nil
}

func validateUsername(val string) error {
	errs := []string{}

	if ok := govalidator.IsAlphanumeric(val); !ok {
		errs = append(errs, "is not alphanumeric")
	}

	if ok := govalidator.IsByteLength(val, 3, 255); !ok {
		errs = append(errs, "is not between 3 and 255 characters long")
	}

	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}

	return nil
}

func validatePassword(val string) error {
	errs := []string{}

	if ok := govalidator.IsByteLength(val, 3, 255); !ok {
		errs = append(errs, "is not between 3 and 255 characters long")
	}

	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}

	return nil
}
