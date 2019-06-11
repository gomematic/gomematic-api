package mysql

import (
	"context"
	"fmt"

	"github.com/gomematic/gomematic-api/pkg/model"
)

// Teams implements teams.Store interface.
type Teams struct {
}

// List implements List from teams.Store interface.
func (t *Teams) List(ctx context.Context) ([]*model.Team, error) {
	return nil, fmt.Errorf("not implemented")
}

// Show implements Show from teams.Store interface.
func (t *Teams) Show(ctx context.Context, name string) (*model.Team, error) {
	return nil, fmt.Errorf("not implemented")
}

// Create implements Create from teams.Store interface.
func (t *Teams) Create(ctx context.Context, team *model.Team) (*model.Team, error) {
	return nil, fmt.Errorf("not implemented")
}

// Update implements Update from teams.Store interface.
func (t *Teams) Update(ctx context.Context, team *model.Team) (*model.Team, error) {
	return nil, fmt.Errorf("not implemented")
}

// Delete implements Delete from teams.Store interface.
func (t *Teams) Delete(ctx context.Context, name string) error {
	return fmt.Errorf("not implemented")
}

// ListUsers implements ListUsers from teams.Store interface.
func (t *Teams) ListUsers(ctx context.Context, id string) ([]*model.TeamUser, error) {
	return nil, fmt.Errorf("not implemented")
}

// AppendUser implements AppendUser from teams.Store interface.
func (t *Teams) AppendUser(ctx context.Context, teamID, userID, perm string) error {
	return fmt.Errorf("not implemented")
}

// PermitUser implements PermitUser from teams.Store interface.
func (t *Teams) PermitUser(ctx context.Context, teamID, userID, perm string) error {
	return fmt.Errorf("not implemented")
}

// DropUser implements DropUser from teams.Store interface.
func (t *Teams) DropUser(ctx context.Context, teamID, userID string) error {
	return fmt.Errorf("not implemented")
}
