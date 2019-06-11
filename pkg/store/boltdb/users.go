package boltdb

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"

	"github.com/Machiel/slugify"
	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	"github.com/gomematic/gomematic-api/pkg/model"
	"github.com/gomematic/gomematic-api/pkg/service/users"
	"github.com/gomematic/gomematic-api/pkg/uuid"
)

var (
	// ErrPasswordEncrypt inditcates that bcrypt failed to create password.
	ErrPasswordEncrypt = errors.New("failed to encrypt password")
)

// Users implements users.Store interface.
type Users struct {
	client *boltdb
}

// ByBasicAuth implements ByBasicAuth from users.Store interface.
func (u *Users) ByBasicAuth(ctx context.Context, username, password string) (*model.User, error) {
	record := &model.User{}

	if err := u.client.handle.Select(
		q.Or(
			q.Eq("Username", username),
			q.Eq("Email", username),
		),
	).First(record); err != nil {
		if err == storm.ErrNotFound {
			return nil, users.ErrNotFound
		}

		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(record.Password),
		[]byte(password),
	); err != nil {
		return nil, users.ErrWrongAuth
	}

	return record, nil
}

// List implements List from users.Store interface.
func (u *Users) List(ctx context.Context) ([]*model.User, error) {
	records := make([]*model.User, 0)

	if err := u.client.handle.AllByIndex("Username", &records); err != nil {
		return nil, err
	}

	return records, nil
}

// Show implements Show from users.Store interface.
func (u *Users) Show(ctx context.Context, name string) (*model.User, error) {
	record := &model.User{}

	if err := u.client.handle.Select(
		q.Or(
			q.Eq("ID", name),
			q.Eq("Slug", name),
		),
	).First(record); err != nil {
		if err == storm.ErrNotFound {
			return record, users.ErrNotFound
		}

		return nil, err
	}

	return record, nil
}

// Create implements Create from users.Store interface.
func (u *Users) Create(ctx context.Context, user *model.User) (*model.User, error) {
	tx, err := u.client.handle.Begin(true)

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	if user.Password != "" && !strings.HasPrefix(user.Password, "$2a") {
		encrypt, err := bcrypt.GenerateFromPassword(
			[]byte(user.Password),
			bcrypt.DefaultCost,
		)

		if err != nil {
			return nil, ErrPasswordEncrypt
		}

		user.Password = string(encrypt)
	}

	if user.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				user.Slug = slugify.Slugify(user.Username)
			} else {
				user.Slug = slugify.Slugify(
					fmt.Sprintf("%s-%d", user.Username, i),
				)
			}

			if err := tx.Select(
				q.Eq("Slug", user.Slug),
			).First(new(model.User)); err != nil {
				if err == storm.ErrNotFound {
					break
				}

				return nil, err
			}
		}
	}

	user.ID = uuid.New().String()
	user.UpdatedAt = time.Now().UTC()
	user.CreatedAt = time.Now().UTC()

	if err := tx.Save(user); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return user, nil
}

// Update implements Update from users.Store interface.
func (u *Users) Update(ctx context.Context, user *model.User) (*model.User, error) {
	tx, err := u.client.handle.Begin(true)

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	if user.Password != "" && !strings.HasPrefix(user.Password, "$2a") {
		encrypt, err := bcrypt.GenerateFromPassword(
			[]byte(user.Password),
			bcrypt.DefaultCost,
		)

		if err != nil {
			return nil, ErrPasswordEncrypt
		}

		user.Password = string(encrypt)
	}

	if user.Slug == "" {
		for i := 0; true; i++ {
			if i == 0 {
				user.Slug = slugify.Slugify(user.Username)
			} else {
				user.Slug = slugify.Slugify(
					fmt.Sprintf("%s-%d", user.Username, i),
				)
			}

			if err := tx.Select(
				q.And(
					q.Eq("Slug", user.Slug),
					q.Not(
						q.Eq("ID", user.ID),
					),
				),
			).First(new(model.User)); err != nil {
				if err == storm.ErrNotFound {
					break
				}

				return nil, err
			}
		}
	}

	user.UpdatedAt = time.Now().UTC()

	if err := tx.Save(user); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return user, nil
}

// Delete implements Delete from users.Store interface.
func (u *Users) Delete(ctx context.Context, name string) error {
	tx, err := u.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := tx.Select(
		q.Or(
			q.Eq("ID", name),
			q.Eq("Slug", name),
		),
	).Delete(new(model.User)); err != nil {
		return err
	}

	return tx.Commit()
}

// ListTeams implements ListTeams from users.Store interface.
func (u *Users) ListTeams(ctx context.Context, id string) ([]*model.TeamUser, error) {
	records := make([]*model.TeamUser, 0)

	if err := u.client.handle.Select(
		q.Eq("UserID", id),
	).Find(&records); err != nil {
		if err == storm.ErrNotFound {
			return records, nil
		}

		return nil, err
	}

	for _, record := range records {
		user, err := u.Show(ctx, record.UserID)

		if err != nil {
			return nil, err
		}

		team, err := u.client.Teams().Show(ctx, record.TeamID)

		if err != nil {
			return nil, err
		}

		record.User = user
		record.Team = team
	}

	return records, nil
}

// AppendTeam implements AppendTeam from teams.Store interface.
func (u *Users) AppendTeam(ctx context.Context, userID, teamID, perm string) error {
	tx, err := u.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := u.client.handle.Select(
		q.And(
			q.Eq("UserID", userID),
			q.Eq("TeamID", teamID),
		),
	).First(new(model.TeamUser)); err == nil {
		return users.ErrAlreadyAssigned
	}

	record := &model.TeamUser{
		UserID:    userID,
		TeamID:    teamID,
		Perm:      perm,
		UpdatedAt: time.Now().UTC(),
		CreatedAt: time.Now().UTC(),
	}

	if err := tx.Save(record); err != nil {
		return err
	}

	return tx.Commit()
}

// PermitTeam implements PermitTeam from teams.Store interface.
func (u *Users) PermitTeam(ctx context.Context, userID, teamID, perm string) error {
	tx, err := u.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()
	record := &model.TeamUser{}

	if err := u.client.handle.Select(
		q.And(
			q.Eq("UserID", userID),
			q.Eq("TeamID", teamID),
		),
	).First(record); err == storm.ErrNotFound {
		return users.ErrNotAssigned
	}

	record.Perm = perm
	record.UpdatedAt = time.Now().UTC()

	if err := tx.Save(record); err != nil {
		return err
	}

	return tx.Commit()
}

// DropTeam implements DropTeam from teams.Store interface.
func (u *Users) DropTeam(ctx context.Context, userID, teamID string) error {
	tx, err := u.client.handle.Begin(true)

	if err != nil {
		return err
	}

	defer tx.Rollback()
	record := &model.TeamUser{}

	if err := u.client.handle.Select(
		q.And(
			q.Eq("UserID", userID),
			q.Eq("TeamID", teamID),
		),
	).First(record); err == storm.ErrNotFound {
		return users.ErrNotAssigned
	}

	if err := tx.DeleteStruct(record); err != nil {
		return err
	}

	return tx.Commit()
}
