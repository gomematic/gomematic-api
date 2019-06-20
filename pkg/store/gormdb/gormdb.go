package gormdb

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"strings"

	"github.com/gomematic/gomematic-api/pkg/model"
	"github.com/gomematic/gomematic-api/pkg/service/teams"
	"github.com/gomematic/gomematic-api/pkg/service/users"
	"github.com/gomematic/gomematic-api/pkg/store"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"

	// Register MySQL driver for GORM
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// Register Postgres driver for GORM
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type gormdb struct {
	driver   string
	username string
	password string
	host     string
	database string
	meta     url.Values

	handle *gorm.DB
	teams  teams.Store
	users  users.Store
}

func (db *gormdb) Teams() teams.Store {
	return db.teams
}

func (db *gormdb) Users() users.Store {
	return db.users
}

func (db *gormdb) Admin(username, password, email string) error {
	admin := &model.User{}

	if err := db.handle.Where(
		&model.User{
			Username: username,
		},
	).First(
		admin,
	).Error; err != nil && err != gorm.ErrRecordNotFound {
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
func (db *gormdb) Info() map[string]interface{} {
	result := make(map[string]interface{})
	result["driver"] = db.driver
	result["host"] = db.host
	result["name"] = db.database
	result["username"] = db.username

	for key, value := range db.meta {
		result[key] = strings.Join(value, "&")
	}

	return result
}

// Close simply closes the MySQL connection.
func (db *gormdb) Open() error {
	connect := ""

	switch db.driver {
	case "mysql":
		if db.password != "" {
			connect = fmt.Sprintf(
				"%s:%s@(%s)/%s?%s",
				db.username,
				db.password,
				db.host,
				db.database,
				db.meta.Encode(),
			)
		} else {
			connect = fmt.Sprintf(
				"%s@(%s)/%s?%s",
				db.username,
				db.host,
				db.database,
				db.meta.Encode(),
			)
		}
	case "postgres":
		host, port, err := net.SplitHostPort(db.host)

		if err != nil {
			return err
		}

		if db.password != "" {
			connect = fmt.Sprintf(
				"host=%s port=%s dbname=%s user=%s password=%s",
				host,
				port,
				db.database,
				db.username,
				db.password,
			)
		} else {
			connect = fmt.Sprintf(
				"host=%s port=%s dbname=%s user=%s",
				host,
				port,
				db.database,
				db.username,
			)
		}

		for key, val := range db.meta {
			connect = fmt.Sprintf("%s %s=%s", connect, key, strings.Join(val, ""))
		}
	}

	handle, err := gorm.Open(
		db.driver,
		connect,
	)

	if err != nil {
		return err
	}

	db.handle = handle

	switch db.driver {
	case "mysql":
		db.handle.DB().SetMaxIdleConns(0)
	}

	return nil
}

// Close simply closes the MySQL connection.
func (db *gormdb) Close() error {
	return db.handle.Close()
}

// Close simply closes the MySQL connection.
func (db *gormdb) Ping() error {
	return db.handle.DB().Ping()
}

// Migrate executes required db migrations.
func (db *gormdb) Migrate() error {
	migrate := gormigrate.New(
		db.handle,
		gormigrate.DefaultOptions,
		migrations,
	)

	return migrate.Migrate()
}

// New initializes a new MySQL connection.
func New(dsn *url.URL) (store.Store, error) {
	client := &gormdb{
		driver:   dsn.Scheme,
		host:     dsn.Host,
		database: strings.TrimPrefix(dsn.Path, "/"),

		username: dsn.User.Username(),
		password: "",
	}

	if password, ok := dsn.User.Password(); ok {
		client.password = password
	}

	switch client.driver {
	case "mysql":
		client.meta = dsn.Query()

		if val := client.meta.Get("charset"); val == "" {
			client.meta.Add("charset", "utf8")
		}

		if val := client.meta.Get("parseTime"); val == "" {
			client.meta.Set("parseTime", "True")
		}

		if val := client.meta.Get("loc"); val == "" {
			client.meta.Set("loc", "Local")
		}
	case "postgres":
		client.meta = dsn.Query()

		if val := client.meta.Get("sslmode"); val == "" {
			client.meta.Set("sslmode", "disable")
		}
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
