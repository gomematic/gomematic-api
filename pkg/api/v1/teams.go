package v1

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gomematic/gomematic-api/pkg/api/v1/models"
	"github.com/gomematic/gomematic-api/pkg/api/v1/restapi/operations/team"
	"github.com/gomematic/gomematic-api/pkg/config"
	"github.com/gomematic/gomematic-api/pkg/model"
	"github.com/gomematic/gomematic-api/pkg/service"
	"github.com/gomematic/gomematic-api/pkg/service/teams"
	"github.com/gomematic/gomematic-api/pkg/service/users"
	"github.com/gomematic/gomematic-api/pkg/upload"
)

// ListTeamsHandler implements the handler for the ListTeams operation.
func ListTeamsHandler(cfg *config.Config, uploads upload.Upload, registry *service.Registry) team.ListTeamsHandlerFunc {
	return func(params team.ListTeamsParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return team.NewListTeamsForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		records, err := registry.Teams.List(params.HTTPRequest.Context())

		if err != nil {
			return team.NewListTeamsDefault(http.StatusInternalServerError)
		}

		payload := make([]*models.Team, len(records))
		for id, record := range records {
			payload[id] = convertTeam(record)
		}

		return team.NewListTeamsOK().WithPayload(payload)
	}
}

// ShowTeamHandler implements the handler for the ShowTeam operation.
func ShowTeamHandler(cfg *config.Config, uploads upload.Upload, registry *service.Registry) team.ShowTeamHandlerFunc {
	return func(params team.ShowTeamParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return team.NewShowTeamForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		record, err := registry.Teams.Show(params.HTTPRequest.Context(), params.TeamID)

		if err != nil {
			if err == teams.ErrNotFound {
				message := "team not found"

				return team.NewShowTeamNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewShowTeamDefault(http.StatusInternalServerError)
		}

		return team.NewShowTeamOK().WithPayload(convertTeam(record))
	}
}

// CreateTeamHandler implements the handler for the CreateTeam operation.
func CreateTeamHandler(cfg *config.Config, uploads upload.Upload, registry *service.Registry) team.CreateTeamHandlerFunc {
	return func(params team.CreateTeamParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return team.NewCreateTeamForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		record := &model.Team{}
		record.Slug = *params.Team.Slug
		record.Name = *params.Team.Name

		created, err := registry.Teams.Create(params.HTTPRequest.Context(), record)

		if err != nil {
			if v, ok := err.(teams.ValidationErrors); ok {
				message := "failed to validate team"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return team.NewCreateTeamUnprocessableEntity().WithPayload(payload)
			}

			return team.NewCreateTeamDefault(http.StatusInternalServerError)
		}

		return team.NewCreateTeamOK().WithPayload(convertTeam(created))
	}
}

// UpdateTeamHandler implements the handler for the UpdateTeam operation.
func UpdateTeamHandler(cfg *config.Config, uploads upload.Upload, registry *service.Registry) team.UpdateTeamHandlerFunc {
	return func(params team.UpdateTeamParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return team.NewUpdateTeamForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		record, err := registry.Teams.Show(params.HTTPRequest.Context(), params.TeamID)

		if err != nil {
			if err == teams.ErrNotFound {
				message := "team not found"

				return team.NewUpdateTeamNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewUpdateTeamDefault(http.StatusInternalServerError)
		}

		if params.Team.Slug != nil {
			record.Slug = *params.Team.Slug
		}

		if params.Team.Name != nil {
			record.Name = *params.Team.Name
		}

		updated, err := registry.Teams.Update(params.HTTPRequest.Context(), record)

		if err != nil {
			if v, ok := err.(users.ValidationErrors); ok {
				message := "failed to validate team"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return team.NewUpdateTeamUnprocessableEntity().WithPayload(payload)
			}

			return team.NewUpdateTeamDefault(http.StatusInternalServerError)
		}

		return team.NewUpdateTeamOK().WithPayload(convertTeam(updated))
	}
}

// DeleteTeamHandler implements the handler for the DeleteTeam operation.
func DeleteTeamHandler(cfg *config.Config, uploads upload.Upload, registry *service.Registry) team.DeleteTeamHandlerFunc {
	return func(params team.DeleteTeamParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return team.NewDeleteTeamForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		record, err := registry.Teams.Show(params.HTTPRequest.Context(), params.TeamID)

		if err != nil {
			if err == teams.ErrNotFound {
				message := "team not found"

				return team.NewDeleteTeamNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewDeleteTeamDefault(http.StatusInternalServerError)
		}

		if err := registry.Teams.Delete(params.HTTPRequest.Context(), record.ID); err != nil {
			message := "failed to delete team"

			return team.NewDeleteTeamBadRequest().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		message := "successfully deleted team"
		return team.NewDeleteTeamOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// ListTeamUsersHandler implements the handler for the ListTeamUsers operation.
func ListTeamUsersHandler(cfg *config.Config, uploads upload.Upload, registry *service.Registry) team.ListTeamUsersHandlerFunc {
	return func(params team.ListTeamUsersParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return team.NewListTeamUsersForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		records, err := registry.Teams.ListUsers(params.HTTPRequest.Context(), params.TeamID)

		if err != nil {
			return team.NewListTeamUsersDefault(http.StatusInternalServerError)
		}

		payload := make([]*models.TeamUser, len(records))
		for id, record := range records {
			payload[id] = convertTeamUser(record)
		}

		return team.NewListTeamUsersOK().WithPayload(payload)
	}
}

// AppendTeamToUserHandler implements the handler for the AppendTeamToUser operation.
func AppendTeamToUserHandler(cfg *config.Config, uploads upload.Upload, registry *service.Registry) team.AppendTeamToUserHandlerFunc {
	return func(params team.AppendTeamToUserParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return team.NewAppendTeamToUserForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		t, err := registry.Teams.Show(params.HTTPRequest.Context(), params.TeamID)

		if err != nil {
			if err == teams.ErrNotFound {
				message := "team not found"

				return team.NewAppendTeamToUserNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewAppendTeamToUserDefault(http.StatusInternalServerError)
		}

		u, err := registry.Users.Show(params.HTTPRequest.Context(), *params.TeamUser.User)

		if err != nil {
			if err == users.ErrNotFound {
				message := "user not found"

				return team.NewAppendTeamToUserNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewAppendTeamToUserDefault(http.StatusInternalServerError)
		}

		if err := registry.Teams.AppendUser(params.HTTPRequest.Context(), t.ID, u.ID, *params.TeamUser.Perm); err != nil {
			if err == teams.ErrAlreadyAssigned {
				message := "user is already assigned"

				return team.NewAppendTeamToUserPreconditionFailed().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			if v, ok := err.(users.ValidationErrors); ok {
				message := "failed to validate team user"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return team.NewAppendTeamToUserUnprocessableEntity().WithPayload(payload)
			}

			return team.NewAppendTeamToUserDefault(http.StatusInternalServerError)
		}

		message := "successfully assigned team to user"
		return team.NewAppendTeamToUserOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// PermitTeamUserHandler implements the handler for the PermitTeamUser operation.
func PermitTeamUserHandler(cfg *config.Config, uploads upload.Upload, registry *service.Registry) team.PermitTeamUserHandlerFunc {
	return func(params team.PermitTeamUserParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return team.NewPermitTeamUserForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		t, err := registry.Teams.Show(params.HTTPRequest.Context(), params.TeamID)

		if err != nil {
			if err == teams.ErrNotFound {
				message := "team not found"

				return team.NewPermitTeamUserNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewPermitTeamUserDefault(http.StatusInternalServerError)
		}

		u, err := registry.Users.Show(params.HTTPRequest.Context(), *params.TeamUser.User)

		if err != nil {
			if err == users.ErrNotFound {
				message := "user not found"

				return team.NewPermitTeamUserNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewPermitTeamUserDefault(http.StatusInternalServerError)
		}

		if err := registry.Teams.PermitUser(params.HTTPRequest.Context(), t.ID, u.ID, *params.TeamUser.Perm); err != nil {
			if err == teams.ErrNotAssigned {
				message := "user is not assigned"

				return team.NewPermitTeamUserPreconditionFailed().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			if v, ok := err.(users.ValidationErrors); ok {
				message := "failed to validate team user"

				payload := &models.ValidationError{
					Message: &message,
				}

				for _, verr := range v.Errors {
					payload.Errors = append(payload.Errors, &models.ValidationErrorErrorsItems0{
						Field:   verr.Field,
						Message: verr.Error.Error(),
					})
				}

				return team.NewPermitTeamUserUnprocessableEntity().WithPayload(payload)
			}

			return team.NewPermitTeamUserDefault(http.StatusInternalServerError)
		}

		message := "successfully updated user perms"
		return team.NewPermitTeamUserOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}

// DeleteTeamFromUserHandler implements the handler for the DeleteTeamFromUser operation.
func DeleteTeamFromUserHandler(cfg *config.Config, uploads upload.Upload, registry *service.Registry) team.DeleteTeamFromUserHandlerFunc {
	return func(params team.DeleteTeamFromUserParams, principal *models.User) middleware.Responder {
		if !*principal.Admin {
			message := "only admins can access this resource"

			return team.NewDeleteTeamFromUserForbidden().WithPayload(&models.GeneralError{
				Message: &message,
			})
		}

		t, err := registry.Teams.Show(params.HTTPRequest.Context(), params.TeamID)

		if err != nil {
			if err == teams.ErrNotFound {
				message := "team not found"

				return team.NewDeleteTeamFromUserNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewDeleteTeamFromUserDefault(http.StatusInternalServerError)
		}

		u, err := registry.Users.Show(params.HTTPRequest.Context(), *params.TeamUser.User)

		if err != nil {
			if err == users.ErrNotFound {
				message := "user not found"

				return team.NewDeleteTeamFromUserNotFound().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewDeleteTeamFromUserDefault(http.StatusInternalServerError)
		}

		if err := registry.Teams.DropUser(params.HTTPRequest.Context(), t.ID, u.ID); err != nil {
			if err == teams.ErrNotAssigned {
				message := "user is not assigned"

				return team.NewDeleteTeamFromUserPreconditionFailed().WithPayload(&models.GeneralError{
					Message: &message,
				})
			}

			return team.NewDeleteTeamFromUserDefault(http.StatusInternalServerError)
		}

		message := "successfully removed from user"
		return team.NewDeleteTeamFromUserOK().WithPayload(&models.GeneralError{
			Message: &message,
		})
	}
}
