package postgres

import (
	"net/url"

	"github.com/gomematic/gomematic-api/pkg/service/teams"
	"github.com/gomematic/gomematic-api/pkg/service/users"
	"github.com/gomematic/gomematic-api/pkg/store"
)

type postgres struct {
	dsn   *url.URL
	teams teams.Store
	users users.Store
}

func (db *postgres) Teams() teams.Store {
	return db.teams
}

func (db *postgres) Users() users.Store {
	return db.users
}

func (db *postgres) Admin(username, password, email string) error {
	return nil
}

// Info returns some basic db informations.
func (db *postgres) Info() map[string]interface{} {
	result := make(map[string]interface{})
	result["driver"] = "postgres"
	result["host"] = db.dsn.Host
	result["username"] = db.dsn.User
	result["name"] = db.dsn.Path

	return result
}

// Close simply closes the PostgreSQL connection.
func (db *postgres) Open() error {
	return nil
}

// Close simply closes the PostgreSQL connection.
func (db *postgres) Close() error {
	return nil
}

// Close simply closes the PostgreSQL connection.
func (db *postgres) Ping() error {
	return nil
}

// New initializes a new PostgreSQL connection.
func New(dsn *url.URL) (store.Store, error) {
	return &postgres{
		dsn:   dsn,
		teams: &Teams{},
		users: &Users{},
	}, nil
}

// Must simply calls New and panics on an error.
func Must(dsn *url.URL) store.Store {
	db, err := New(dsn)

	if err != nil {
		panic(err)
	}

	return db
}
