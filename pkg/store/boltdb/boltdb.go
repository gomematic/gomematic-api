package boltdb

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	"github.com/gomematic/gomematic-api/pkg/model"
	"github.com/gomematic/gomematic-api/pkg/service/teams"
	"github.com/gomematic/gomematic-api/pkg/service/users"
	"github.com/gomematic/gomematic-api/pkg/store"
	"github.com/rs/zerolog/log"

	bolt "go.etcd.io/bbolt"
)

type boltdb struct {
	dsn    *url.URL
	handle *storm.DB
	teams  teams.Store
	users  users.Store
}

func (db *boltdb) Teams() teams.Store {
	return db.teams
}

func (db *boltdb) Users() users.Store {
	return db.users
}

func (db *boltdb) Admin(username, password, email string) error {
	admin := &model.User{}

	if err := db.handle.Select(
		q.Eq("Username", username),
	).First(admin); err != nil && err != storm.ErrNotFound {
		return err
	}

	admin.Username = username
	admin.Password = password
	admin.Email = email
	admin.Active = true
	admin.Admin = true

	if admin.ID == "" {
		if _, err := db.users.Create(
			context.Background(),
			admin,
		); err != nil {
			return err
		}
	} else {
		if _, err := db.users.Update(
			context.Background(),
			admin,
		); err != nil {
			return err
		}
	}

	for i := 1; i <= 9; i++ {
		if _, err := db.users.Create(context.TODO(), &model.User{
			Username: fmt.Sprintf("user%d", i),
			Password: fmt.Sprintf("user%d", i),
			Email:    fmt.Sprintf("user%d@webhippie.de", i),
			Active:   true,
			Admin:    false,
		}); err != nil {
			log.Warn().
				Err(err).
				Msg("user")
		}

		if _, err := db.teams.Create(context.TODO(), &model.Team{
			Name: fmt.Sprintf("Team %d", i),
		}); err != nil {
			log.Warn().
				Err(err).
				Msg("team")
		}
	}

	return nil
}

// Info returns some basic db informations.
func (db *boltdb) Info() map[string]interface{} {
	result := make(map[string]interface{})
	result["driver"] = "boltdb"
	result["path"] = db.path()
	result["timeout"] = db.timeout().String()

	return result
}

// Close simply closes the BoltDB connection.
func (db *boltdb) Open() error {
	handle, err := storm.Open(
		db.path(),
		storm.BoltOptions(
			db.perms(),
			&bolt.Options{
				Timeout: db.timeout(),
			},
		),
	)

	if err != nil {
		return err
	}

	db.handle = handle
	return nil
}

// Close simply closes the BoltDB connection.
func (db *boltdb) Close() error {
	return db.handle.Close()
}

// Close simply closes the BoltDB connection.
func (db *boltdb) Ping() error {
	return nil
}

// perms retrieves the dir perms from dsn or fallback.
func (db *boltdb) perms() os.FileMode {
	if val := db.dsn.Query().Get("perms"); val != "" {
		res, err := strconv.ParseUint(val, 8, 32)

		if err != nil {
			return os.FileMode(0600)
		}

		return os.FileMode(res)
	}

	return os.FileMode(0600)
}

// timeout retrieves the timeout from dsn or fallback.
func (db *boltdb) timeout() time.Duration {
	if val := db.dsn.Query().Get("timeout"); val != "" {
		res, err := time.ParseDuration(val)

		if err != nil {
			return 1 * time.Second
		}

		return res
	}

	return 1 * time.Second
}

// path cleans the dsn and returns a valid path.
func (db *boltdb) path() string {
	return path.Join(
		db.dsn.Host,
		db.dsn.EscapedPath(),
	)
}

// New initializes a new BoltDB connection.
func New(dsn *url.URL) (store.Store, error) {
	client := &boltdb{
		dsn: dsn,
	}

	if err := client.Open(); err != nil {
		return nil, err
	}

	client.teams = &Teams{
		client: client,
	}

	client.users = &Users{
		client: client,
	}

	return client, nil
}

// Must simply calls New and panics on an error.
func Must(dsn *url.URL) store.Store {
	db, err := New(dsn)

	if err != nil {
		panic(err)
	}

	return db
}
