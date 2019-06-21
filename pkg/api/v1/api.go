package v1

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/strfmt"
	"github.com/gomematic/gomematic-api/pkg/api/v1/models"
	"github.com/gomematic/gomematic-api/pkg/api/v1/restapi"
	"github.com/gomematic/gomematic-api/pkg/api/v1/restapi/operations"
	"github.com/gomematic/gomematic-api/pkg/config"
	"github.com/gomematic/gomematic-api/pkg/middleware/requestid"
	"github.com/gomematic/gomematic-api/pkg/model"
	"github.com/gomematic/gomematic-api/pkg/service"
	"github.com/gomematic/gomematic-api/pkg/token"
	"github.com/gomematic/gomematic-api/pkg/upload"
	"github.com/rs/zerolog/log"
)

//go:generate gorunpkg github.com/go-swagger/go-swagger/cmd/swagger generate server --target . --name Gomematic --spec ../../../openapi/v1.yml --template-dir ../../../hack/openapi/templates -P models.User --exclude-main --regenerate-configureapi

// API provides the http.Handler for the OpenAPI implementation.
type API struct {
	Handler http.Handler
}

// New creates a new API that adds the custom Handler implementations.
func New(cfg *config.Config, uploads upload.Upload, registry *service.Registry) *API {
	spec, err := loads.Analyzed(restapi.SwaggerJSON, "")

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to analyze openapi")

		return nil
	}

	api := operations.NewGomematicAPI(spec)

	api.AuthVerifyAuthHandler = VerifyAuthHandler(cfg, uploads, registry)
	api.AuthRefreshAuthHandler = RefreshAuthHandler(cfg, uploads, registry)
	api.AuthLoginUserHandler = LoginUserHandler(cfg, uploads, registry)

	api.ProfileTokenProfileHandler = TokenProfileHandler(cfg, uploads, registry)
	api.ProfileShowProfileHandler = ShowProfileHandler(cfg, uploads, registry)
	api.ProfileUpdateProfileHandler = UpdateProfileHandler(cfg, uploads, registry)

	api.TeamListTeamsHandler = ListTeamsHandler(cfg, uploads, registry)
	api.TeamShowTeamHandler = ShowTeamHandler(cfg, uploads, registry)
	api.TeamCreateTeamHandler = CreateTeamHandler(cfg, uploads, registry)
	api.TeamUpdateTeamHandler = UpdateTeamHandler(cfg, uploads, registry)
	api.TeamDeleteTeamHandler = DeleteTeamHandler(cfg, uploads, registry)
	api.TeamListTeamUsersHandler = ListTeamUsersHandler(cfg, uploads, registry)
	api.TeamAppendTeamToUserHandler = AppendTeamToUserHandler(cfg, uploads, registry)
	api.TeamPermitTeamUserHandler = PermitTeamUserHandler(cfg, uploads, registry)
	api.TeamDeleteTeamFromUserHandler = DeleteTeamFromUserHandler(cfg, uploads, registry)

	api.UserListUsersHandler = ListUsersHandler(cfg, uploads, registry)
	api.UserShowUserHandler = ShowUserHandler(cfg, uploads, registry)
	api.UserCreateUserHandler = CreateUserHandler(cfg, uploads, registry)
	api.UserUpdateUserHandler = UpdateUserHandler(cfg, uploads, registry)
	api.UserDeleteUserHandler = DeleteUserHandler(cfg, uploads, registry)
	api.UserListUserTeamsHandler = ListUserTeamsHandler(cfg, uploads, registry)
	api.UserAppendUserToTeamHandler = AppendUserToTeamHandler(cfg, uploads, registry)
	api.UserPermitUserTeamHandler = PermitUserTeamHandler(cfg, uploads, registry)
	api.UserDeleteUserFromTeamHandler = DeleteUserFromTeamHandler(cfg, uploads, registry)

	api.HeaderAuth = func(ctx context.Context, val string) (context.Context, *models.User, error) {
		t, err := token.Parse(val, cfg.Session.Secret)

		if err != nil {
			log.Warn().
				Str("request", requestid.Get(ctx)).
				Err(err).
				Str("token", val).
				Msg("failed to parse token")

			return ctx, nil, errors.New(401, "incorrect auth")
		}

		user, err := registry.Users.Show(
			ctx,
			t.Text,
		)

		if err != nil {
			log.Warn().
				Str("request", requestid.Get(ctx)).
				Err(err).
				Str("token", val).
				Msg("failed to fetch user")

			return ctx, nil, errors.New(401, "incorrect auth")
		}

		return ctx, convertUser(user), nil
	}

	api.BasicAuth = func(ctx context.Context, username, password string) (context.Context, *models.User, error) {
		user, err := registry.Users.ByBasicAuth(
			ctx,
			username,
			password,
		)

		if err != nil {
			log.Warn().
				Str("request", requestid.Get(ctx)).
				Err(err).
				Str("username", username).
				Msg("failed to auth user")

			return ctx, nil, errors.New(401, "incorrect auth")
		}

		return ctx, convertUser(user), nil
	}

	return &API{
		Handler: api.Serve(nil),
	}
}

func convertAuthToken(record *token.Result) *models.AuthToken {
	if record.ExpiresAt.IsZero() {
		return &models.AuthToken{
			Token: record.Token,
		}
	}

	expiresAt := strfmt.DateTime(record.ExpiresAt)

	return &models.AuthToken{
		Token:     record.Token,
		ExpiresAt: &expiresAt,
	}
}

func convertAuthVerify(record *models.User) *models.AuthVerify {
	createdAt := strfmt.DateTime(record.CreatedAt)

	return &models.AuthVerify{
		Username:  *record.Username,
		CreatedAt: &createdAt,
	}
}

// convertProfile is a simple helper to convert between different model formats.
func convertProfile(record *model.User) *models.Profile {
	res := &models.Profile{
		ID:        strfmt.UUID(record.ID),
		Slug:      &record.Slug,
		Username:  &record.Username,
		Password:  nil,
		Email:     &record.Email,
		Active:    &record.Active,
		Admin:     &record.Admin,
		CreatedAt: strfmt.DateTime(record.CreatedAt),
		UpdatedAt: strfmt.DateTime(record.UpdatedAt),
	}

	for _, row := range record.Teams {
		res.Teams = append(res.Teams, convertTeamUser(row))
	}

	return res
}

// convertTeam is a simple helper to convert between different model formats.
func convertTeam(record *model.Team) *models.Team {
	res := &models.Team{
		ID:        strfmt.UUID(record.ID),
		Slug:      &record.Slug,
		Name:      &record.Name,
		CreatedAt: strfmt.DateTime(record.CreatedAt),
		UpdatedAt: strfmt.DateTime(record.UpdatedAt),
	}

	for _, row := range record.Users {
		res.Users = append(res.Users, convertTeamUser(row))
	}

	return res
}

// convertUser is a simple helper to convert between different model formats.
func convertUser(record *model.User) *models.User {
	res := &models.User{
		ID:        strfmt.UUID(record.ID),
		Slug:      &record.Slug,
		Username:  &record.Username,
		Password:  nil,
		Email:     &record.Email,
		Active:    &record.Active,
		Admin:     &record.Admin,
		CreatedAt: strfmt.DateTime(record.CreatedAt),
		UpdatedAt: strfmt.DateTime(record.UpdatedAt),
	}

	for _, row := range record.Teams {
		res.Teams = append(res.Teams, convertTeamUser(row))
	}

	return res
}

// convertTeamUser is a simple helper to convert between different model formats.
func convertTeamUser(record *model.TeamUser) *models.TeamUser {
	userID := strfmt.UUID(record.UserID)
	teamID := strfmt.UUID(record.TeamID)

	return &models.TeamUser{
		TeamID:    &teamID,
		Team:      convertTeam(record.Team),
		UserID:    &userID,
		User:      convertUser(record.User),
		Perm:      &record.Perm,
		CreatedAt: strfmt.DateTime(record.CreatedAt),
		UpdatedAt: strfmt.DateTime(record.UpdatedAt),
	}
}
