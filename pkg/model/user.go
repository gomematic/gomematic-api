package model

import (
	"time"
)

// User within Gomematic.
type User struct {
	ID        string `storm:"id" gorm:"primary_key"`
	Slug      string `storm:"unique" sql:"unique_index"`
	Email     string `storm:"unique" sql:"unique_index"`
	Username  string `storm:"unique" sql:"unique_index"`
	Password  string
	Active    bool `sql:"default:false"`
	Admin     bool `sql:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Teams     []*TeamUser
}
