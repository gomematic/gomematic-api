package model

import (
	"time"
)

// User within Gomematic.
type User struct {
	ID        string `storm:"id"`
	Slug      string `storm:"index,unique"`
	Email     string `storm:"unique"`
	Username  string `storm:"unique"`
	Password  string
	Active    bool
	Admin     bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
