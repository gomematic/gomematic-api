package boltdb

import (
	"context"
	"fmt"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asaskevich/govalidator"
	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	"github.com/gomematic/gomematic-api/pkg/model"
	"github.com/gomematic/gomematic-api/pkg/service/teams"
	"github.com/gomematic/gomematic-api/pkg/uuid"
	"github.com/gomematic/gomematic-api/pkg/validation"
)

// Teams implements teams.Store interface.
type Teams struct {
	client *boltdb
}

// List implements List from teams.Store interface.
func (t *Teams) List(ctx context.Context) ([]*model.Team, error) {
	records := make([]*model.Team, 0)

	err := t.client.handle.AllByIndex(
		"Name",
		&records,
	)

	if err == storm.ErrNotFound {
		return records, nil
	}

	return records, nil
}

// Show implements Show from teams.Store interface.
func (t *Teams) Show(ctx context.Context, name string) (*model.Team, error) {
	record := &model.Team{}

	if err := t.client.handle.Select(
		q.Or(
			q.Eq("ID", name),
			q.Eq("Slug", name),
		),
	).First(record); err != nil {
		if err == storm.ErrNotFound {
			return record, teams.ErrNotFound
		}

		return nil, err
	}

	return record, nil
}

// Create implements Create from teams.Store interface.
func (t *Teams) Create(ctx context.Context, team *model.Team) (*model.Team, error) {
	tx, err := t.client.handle.Begin(true)

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	if team.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				team.Slug = slugify.Slugify(team.Name)
			} else {
				team.Slug = slugify.Slugify(
					fmt.Sprintf("%s-%d", team.Name, i),
				)
			}

			if err := tx.Select(
				q.Eq("Slug", team.Slug),
			).First(new(model.Team)); err != nil {
				if err == storm.ErrNotFound {
					break
				}

				return nil, err
			}
		}
	}

	team.ID = uuid.New().String()
	team.UpdatedAt = time.Now().UTC()
	team.CreatedAt = time.Now().UTC()

	if err := t.validateCreate(team); err != nil {
		return nil, err
	}

	if err := tx.Save(team); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return team, nil
}

// Update implements Update from teams.Store interface.
func (t *Teams) Update(ctx context.Context, team *model.Team) (*model.Team, error) {
	tx, err := t.client.handle.Begin(true)

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	if team.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				team.Slug = slugify.Slugify(team.Name)
			} else {
				team.Slug = slugify.Slugify(
					fmt.Sprintf("%s-%d", team.Name, i),
				)
			}

			if err := tx.Select(
				q.And(
					q.Eq("Slug", team.Slug),
					q.Not(
						q.Eq("ID", team.ID),
					),
				),
			).First(new(model.Team)); err != nil {
				if err == storm.ErrNotFound {
					break
				}

				return nil, err
			}
		}
	}

	team.UpdatedAt = time.Now().UTC()

	if err := t.validateUpdate(team); err != nil {
		return nil, err
	}

	if err := tx.Save(team); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return team, nil
}

// Delete implements Delete from teams.Store interface.
func (t *Teams) Delete(ctx context.Context, name string) error {
	tx, err := t.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := tx.Select(
		q.Or(
			q.Eq("ID", name),
			q.Eq("Slug", name),
		),
	).Delete(new(model.Team)); err != nil {
		return err
	}

	return tx.Commit()
}

// ListUsers implements ListUsers from teams.Store interface.
func (t *Teams) ListUsers(ctx context.Context, id string) ([]*model.TeamUser, error) {
	records := make([]*model.TeamUser, 0)

	if err := t.client.handle.Select(
		q.Eq("TeamID", id),
	).Find(&records); err != nil {
		if err == storm.ErrNotFound {
			return records, nil
		}

		return nil, err
	}

	for _, record := range records {
		team, err := t.Show(ctx, record.TeamID)

		if err != nil {
			return nil, err
		}

		user, err := t.client.Users().Show(ctx, record.UserID)

		if err != nil {
			return nil, err
		}

		record.Team = team
		record.User = user
	}

	return records, nil
}

// AppendUser implements AppendUser from teams.Store interface.
func (t *Teams) AppendUser(ctx context.Context, teamID, userID, perm string) error {
	tx, err := t.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := t.client.handle.Select(
		q.And(
			q.Eq("TeamID", teamID),
			q.Eq("UserID", userID),
		),
	).First(new(model.TeamUser)); err == nil {
		return teams.ErrAlreadyAssigned
	}

	record := &model.TeamUser{
		TeamID:    teamID,
		UserID:    userID,
		Perm:      perm,
		UpdatedAt: time.Now().UTC(),
		CreatedAt: time.Now().UTC(),
	}

	if err := t.validatePerm(record); err != nil {
		return err
	}

	if err := tx.Save(record); err != nil {
		return err
	}

	return tx.Commit()
}

// PermitUser implements PermitUser from teams.Store interface.
func (t *Teams) PermitUser(ctx context.Context, teamID, userID, perm string) error {
	tx, err := t.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()
	record := &model.TeamUser{}

	if err := t.client.handle.Select(
		q.And(
			q.Eq("TeamID", teamID),
			q.Eq("UserID", userID),
		),
	).First(record); err == storm.ErrNotFound {
		return teams.ErrNotAssigned
	}

	record.Perm = perm
	record.UpdatedAt = time.Now().UTC()

	if err := t.validatePerm(record); err != nil {
		return err
	}

	if err := tx.Save(record); err != nil {
		return err
	}

	return tx.Commit()
}

// DropUser implements DropUser from teams.Store interface.
func (t *Teams) DropUser(ctx context.Context, teamID, userID string) error {
	tx, err := t.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()
	record := &model.TeamUser{}

	if err := t.client.handle.Select(
		q.And(
			q.Eq("TeamID", teamID),
			q.Eq("UserID", userID),
		),
	).First(record); err == storm.ErrNotFound {
		return teams.ErrNotAssigned
	}

	if err := tx.DeleteStruct(record); err != nil {
		return err
	}

	return tx.Commit()
}

func (t *Teams) validateCreate(record *model.Team) error {
	errs := validation.Errors{}

	if ok := govalidator.IsByteLength(record.Slug, 3, 255); !ok {
		errs.Errors = append(errs.Errors, validation.Error{
			Field: "slug",
			Error: fmt.Errorf("is not between 3 and 255 characters long"),
		})
	}

	if t.uniqueValueIsPresent("Slug", record.Slug, record.ID) {
		errs.Errors = append(errs.Errors, validation.Error{
			Field: "slug",
			Error: fmt.Errorf("is already taken"),
		})
	}

	if ok := govalidator.IsByteLength(record.Name, 3, 255); !ok {
		errs.Errors = append(errs.Errors, validation.Error{
			Field: "name",
			Error: fmt.Errorf("is not between 3 and 255 characters long"),
		})
	}

	if t.uniqueValueIsPresent("Name", record.Name, record.ID) {
		errs.Errors = append(errs.Errors, validation.Error{
			Field: "name",
			Error: fmt.Errorf("is already taken"),
		})
	}

	if len(errs.Errors) > 0 {
		return errs
	}

	return nil
}

func (t *Teams) validateUpdate(record *model.Team) error {
	errs := validation.Errors{}

	if ok := govalidator.IsUUIDv4(record.ID); !ok {
		errs.Errors = append(errs.Errors, validation.Error{
			Field: "id",
			Error: fmt.Errorf("is not a valid uuid v4"),
		})
	}

	if ok := govalidator.IsByteLength(record.Slug, 3, 255); !ok {
		errs.Errors = append(errs.Errors, validation.Error{
			Field: "slug",
			Error: fmt.Errorf("is not between 3 and 255 characters long"),
		})
	}

	if t.uniqueValueIsPresent("Slug", record.Slug, record.ID) {
		errs.Errors = append(errs.Errors, validation.Error{
			Field: "slug",
			Error: fmt.Errorf("is already taken"),
		})
	}

	if ok := govalidator.IsByteLength(record.Name, 3, 255); !ok {
		errs.Errors = append(errs.Errors, validation.Error{
			Field: "name",
			Error: fmt.Errorf("is not between 3 and 255 characters long"),
		})
	}

	if t.uniqueValueIsPresent("Name", record.Name, record.ID) {
		errs.Errors = append(errs.Errors, validation.Error{
			Field: "name",
			Error: fmt.Errorf("is already taken"),
		})
	}

	if len(errs.Errors) > 0 {
		return errs
	}

	return nil
}

func (t *Teams) validatePerm(record *model.TeamUser) error {
	if ok := govalidator.IsIn(record.Perm, "user", "admin", "owner"); !ok {
		return validation.Errors{
			Errors: []validation.Error{
				validation.Error{
					Field: "perm",
					Error: fmt.Errorf("invalid permission value"),
				},
			},
		}
	}

	return nil
}

func (t *Teams) uniqueValueIsPresent(key, val, id string) bool {
	if err := t.client.handle.Select(
		q.And(
			q.Eq(key, val),
			q.Not(
				q.Eq("ID", id),
			),
		),
	).First(new(model.Team)); err == storm.ErrNotFound {
		return false
	}

	return true
}
