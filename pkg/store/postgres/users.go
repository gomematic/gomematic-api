package postgres

import (
	"context"
	"fmt"

	"github.com/gomematic/gomematic-api/pkg/model"
)

// Users implements users.Store interface.
type Users struct {
}

// ByBasicAuth implements ByBasicAuth from users.Store interface.
func (u *Users) ByBasicAuth(ctx context.Context, username, password string) (*model.User, error) {
	return nil, fmt.Errorf("not implemented")
}

// List implements List from users.Store interface.
func (u *Users) List(ctx context.Context) ([]*model.User, error) {
	return nil, fmt.Errorf("not implemented")
}

// Show implements Show from users.Store interface.
func (u *Users) Show(ctx context.Context, name string) (*model.User, error) {
	return nil, fmt.Errorf("not implemented")
}

// Create implements Create from users.Store interface.
func (u *Users) Create(ctx context.Context, user *model.User) (*model.User, error) {
	return nil, fmt.Errorf("not implemented")
}

// Update implements Update from users.Store interface.
func (u *Users) Update(ctx context.Context, user *model.User) (*model.User, error) {
	return nil, fmt.Errorf("not implemented")
}

// Delete implements Delete from users.Store interface.
func (u *Users) Delete(ctx context.Context, name string) error {
	return fmt.Errorf("not implemented")
}

// ListTeams implements ListTeams from users.Store interface.
func (u *Users) ListTeams(ctx context.Context, id string) ([]*model.TeamUser, error) {
	return nil, fmt.Errorf("not implemented")
}

// AppendTeam implements AppendTeam from teams.Store interface.
func (u *Users) AppendTeam(ctx context.Context, userID, teamID, perm string) error {
	return fmt.Errorf("not implemented")
}

// PermitTeam implements PermitTeam from teams.Store interface.
func (u *Users) PermitTeam(ctx context.Context, userID, teamID, perm string) error {
	return fmt.Errorf("not implemented")
}

// DropTeam implements DropTeam from teams.Store interface.
func (u *Users) DropTeam(ctx context.Context, userID, teamID string) error {
	return fmt.Errorf("not implemented")
}
