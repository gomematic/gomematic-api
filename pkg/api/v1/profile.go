package v1

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gomematic/gomematic-api/pkg/api/v1/models"
	"github.com/gomematic/gomematic-api/pkg/api/v1/restapi/operations/profile"
	"github.com/gomematic/gomematic-api/pkg/config"
	"github.com/gomematic/gomematic-api/pkg/service"
	"github.com/gomematic/gomematic-api/pkg/service/users"
	"github.com/gomematic/gomematic-api/pkg/token"
	"github.com/gomematic/gomematic-api/pkg/upload"
)

// TokenProfileHandler implements the handler for the ProfileTokenProfile operation.
func TokenProfileHandler(cfg *config.Config, uploads upload.Upload, registry *service.Registry) profile.TokenProfileHandlerFunc {
	return func(params profile.TokenProfileParams, principal *models.User) middleware.Responder {
		result, err := token.New(*principal.Username).Unlimited(cfg.Session.Secret)

		if err != nil {
			return profile.NewTokenProfileDefault(http.StatusInternalServerError)
		}

		return profile.NewTokenProfileOK().WithPayload(convertAuthToken(result))
	}
}

// ShowProfileHandler implements the handler for the ProfileShowProfile operation.
func ShowProfileHandler(cfg *config.Config, uploads upload.Upload, registry *service.Registry) profile.ShowProfileHandlerFunc {
	return func(params profile.ShowProfileParams, principal *models.User) middleware.Responder {
		record, err := registry.Users.Show(params.HTTPRequest.Context(), principal.ID.String())

		if err != nil {
			return profile.NewShowProfileDefault(http.StatusInternalServerError)
		}

		return profile.NewShowProfileOK().WithPayload(convertProfile(record))
	}
}

// UpdateProfileHandler implements the handler for the ProfileUpdateProfile operation.
func UpdateProfileHandler(cfg *config.Config, uploads upload.Upload, registry *service.Registry) profile.UpdateProfileHandlerFunc {
	return func(params profile.UpdateProfileParams, principal *models.User) middleware.Responder {
		record, err := registry.Users.Show(params.HTTPRequest.Context(), principal.ID.String())

		if err != nil {
			return profile.NewUpdateProfileDefault(http.StatusInternalServerError)
		}

		if params.Profile.Slug != nil {
			record.Slug = *params.Profile.Slug
		}

		if params.Profile.Username != nil {
			record.Username = *params.Profile.Username
		}

		if params.Profile.Password != nil {
			record.Password = (*params.Profile.Password).String()
		}

		if params.Profile.Email != nil {
			record.Email = *params.Profile.Email
		}

		updated, err := registry.Users.Update(params.HTTPRequest.Context(), record)

		if err != nil {
			if v, ok := err.(users.ValidationErrors); ok {
				message := "failed to validate profile"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return profile.NewUpdateProfileUnprocessableEntity().WithPayload(payload)
			}

			return profile.NewUpdateProfileDefault(http.StatusInternalServerError)
		}

		return profile.NewUpdateProfileOK().WithPayload(convertProfile(updated))
	}
}
