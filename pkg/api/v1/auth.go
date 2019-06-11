package v1

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gomematic/gomematic-api/pkg/api/v1/models"
	"github.com/gomematic/gomematic-api/pkg/api/v1/restapi/operations/auth"
	"github.com/gomematic/gomematic-api/pkg/config"
	"github.com/gomematic/gomematic-api/pkg/service"
	"github.com/gomematic/gomematic-api/pkg/service/users"
	"github.com/gomematic/gomematic-api/pkg/token"
	"github.com/gomematic/gomematic-api/pkg/upload"
)

// LoginUserHandler implements the handler for the AuthLoginUser operation.
func LoginUserHandler(cfg *config.Config, uploads upload.Upload, registry *service.Registry) auth.LoginUserHandlerFunc {
	return func(params auth.LoginUserParams) middleware.Responder {
		user, err := registry.Users.ByBasicAuth(
			params.HTTPRequest.Context(),
			*params.AuthLogin.Username,
			params.AuthLogin.Password.String(),
		)

		if err != nil {
			if err == users.ErrNotFound {
				message := "wrong username or password"

				return auth.NewLoginUserUnauthorized().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			if err == users.ErrWrongAuth {
				message := "wrong username or password"

				return auth.NewLoginUserUnauthorized().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return auth.NewLoginUserDefault(http.StatusInternalServerError)
		}

		result, err := token.New(user.Username).Expiring(cfg.Session.Secret, cfg.Session.Expire)

		if err != nil {
			return auth.NewLoginUserDefault(http.StatusInternalServerError)
		}

		return auth.NewLoginUserOK().WithPayload(convertAuthToken(result))
	}
}

// RefreshAuthHandler implements the handler for the AuthRefreshAuth operation.
func RefreshAuthHandler(cfg *config.Config, uploads upload.Upload, registry *service.Registry) auth.RefreshAuthHandlerFunc {
	return func(params auth.RefreshAuthParams, principal *models.User) middleware.Responder {
		result, err := token.New(*principal.Username).Expiring(cfg.Session.Secret, cfg.Session.Expire)

		if err != nil {
			return auth.NewRefreshAuthDefault(http.StatusInternalServerError)
		}

		return auth.NewRefreshAuthOK().WithPayload(convertAuthToken(result))
	}
}

// VerifyAuthHandler implements the handler for the AuthVerifyAuth operation.
func VerifyAuthHandler(cfg *config.Config, uploads upload.Upload, registry *service.Registry) auth.VerifyAuthHandlerFunc {
	return func(params auth.VerifyAuthParams, principal *models.User) middleware.Responder {
		return auth.NewVerifyAuthOK().WithPayload(convertAuthVerify(principal))
	}
}
