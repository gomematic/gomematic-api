package boltdb

import (
	"context"
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

	bolt "go.etcd.io/bbolt"
)

type boltdb struct {
	path    string
	perms   os.FileMode
	timeout time.Duration

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

	return nil
}

// Info returns some basic db informations.
func (db *boltdb) Info() map[string]interface{} {
	result := make(map[string]interface{})
	result["driver"] = "boltdb"
	result["path"] = db.path
	result["perms"] = db.perms.String()
	result["timeout"] = db.timeout.String()

	return result
}

// Prepare is preparing some database behavior.
func (db *boltdb) Prepare() error {
	return nil
}

// Open simply opens the BoltDB connection.
func (db *boltdb) Open() error {
	handle, err := storm.Open(
		db.path,
		storm.BoltOptions(
			db.perms,
			&bolt.Options{
				Timeout: db.timeout,
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

// Ping checks the connection to BoltDB.
func (db *boltdb) Ping() error {
	return nil
}

// Migrate executes required db migrations.
func (db *boltdb) Migrate() error {
	return nil
}

// New initializes a new BoltDB connection.
func New(dsn *url.URL) (store.Store, error) {
	client := &boltdb{
		path: path.Join(
			dsn.Host,
			dsn.EscapedPath(),
		),
	}

	if val := dsn.Query().Get("perms"); val != "" {
		res, err := strconv.ParseUint(val, 8, 32)

		if err != nil {
			client.perms = os.FileMode(0600)
		} else {
			client.perms = os.FileMode(res)
		}
	} else {
		client.perms = os.FileMode(0600)
	}

	if val := dsn.Query().Get("timeout"); val != "" {
		res, err := time.ParseDuration(val)

		if err != nil {
			client.timeout = 1 * time.Second
		} else {
			client.timeout = res
		}
	} else {
		client.timeout = 1 * time.Second
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
