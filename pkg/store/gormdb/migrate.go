package gormdb

import (
	"time"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var (
	migrations = []*gormigrate.Migration{
		{
			ID: "201609011300",
			Migrate: func(tx *gorm.DB) error {
				type User struct {
					ID        string `gorm:"primary_key"`
					Slug      string `sql:"unique_index"`
					Username  string `sql:"unique_index"`
					Email     string `sql:"unique_index"`
					Password  string
					Active    bool `sql:"default:false"`
					Admin     bool `sql:"default:false"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.CreateTable(&User{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("users").Error
			},
		},
		{
			ID: "201609011301",
			Migrate: func(tx *gorm.DB) error {
				type Team struct {
					ID        string `gorm:"primary_key"`
					Slug      string `sql:"unique_index"`
					Name      string `sql:"unique_index"`
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.CreateTable(&Team{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("teams").Error
			},
		},
		{
			ID: "201609011302",
			Migrate: func(tx *gorm.DB) error {
				type TeamUser struct {
					TeamID    string `sql:"index"`
					UserID    string `sql:"index"`
					Perm      string
					CreatedAt time.Time
					UpdatedAt time.Time
				}

				return tx.CreateTable(&TeamUser{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("team_users").Error
			},
		},
		{
			ID: "201609011303",
			Migrate: func(tx *gorm.DB) error {
				return tx.Table(
					"team_users",
				).AddForeignKey(
					"team_id",
					"teams(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Table(
					"team_users",
				).RemoveForeignKey(
					"team_id",
					"teams(id)",
				).Error
			},
		},
		{
			ID: "201609011304",
			Migrate: func(tx *gorm.DB) error {
				return tx.Table(
					"team_users",
				).AddForeignKey(
					"user_id",
					"users(id)",
					"RESTRICT",
					"RESTRICT",
				).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Table(
					"team_users",
				).RemoveForeignKey(
					"user_id",
					"users(id)",
				).Error
			},
		},
	}
)
