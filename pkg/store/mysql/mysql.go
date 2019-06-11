package mysql

import (
	"net/url"

	"github.com/gomematic/gomematic-api/pkg/service/teams"
	"github.com/gomematic/gomematic-api/pkg/service/users"
	"github.com/gomematic/gomematic-api/pkg/store"
)

type mysql struct {
	dsn   *url.URL
	teams teams.Store
	users users.Store
}

func (db *mysql) Teams() teams.Store {
	return db.teams
}

func (db *mysql) Users() users.Store {
	return db.users
}

func (db *mysql) Admin(username, password, email string) error {
	return nil
}

// Info returns some basic db informations.
func (db *mysql) Info() map[string]interface{} {
	result := make(map[string]interface{})
	result["driver"] = "mysql"
	result["host"] = db.dsn.Host
	result["username"] = db.dsn.User
	result["name"] = db.dsn.Path

	return result
}

// Close simply closes the MySQL connection.
func (db *mysql) Open() error {
	return nil
}

// Close simply closes the MySQL connection.
func (db *mysql) Close() error {
	return nil
}

// Close simply closes the MySQL connection.
func (db *mysql) Ping() error {
	return nil
}

// New initializes a new MySQL connection.
func New(dsn *url.URL) (store.Store, error) {
	return &mysql{
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
